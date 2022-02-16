package controllers

import (
	"b/container"
	"b/kafkaDocker"
	"b/models"
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BookController struct {
	container container.Container
}

type User struct {
	// example: 620916dd86cae49baedc7426
	Id string `json:"id"`
	// example: 620aa5e18e9c117fdbb9f4d2
	Story string `json:"story"`
}

func NewBookController(container container.Container) *BookController {
	return &BookController{container: container}
}

func (controller *BookController) SaveFromCSV(c echo.Context) (err error) {
	fmt.Println("Uploading File")
	_ = c.Request().ParseMultipartForm(10 << 20)
	fileHeader, err := c.FormFile("MyFile")
	if err != nil {
		msg := models.Response{Message: err.Error()}
		return c.JSON(500, msg)
	}
	file, err := fileHeader.Open()
	if err != nil {
		msg := models.Response{Message: err.Error()}
		return c.JSON(500, msg)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		msg := models.Response{Message: err.Error()}
		return c.JSON(500, msg)
	}
	var books = make([]models.Books, 0)
	for i, rec := range records {
		if i == 0 {
			continue
		}
		layout := "02-01-2006 15:04"
		t, err := time.Parse(layout, rec[3])
		if err != nil {
			msg := models.Response{Message: err.Error()}
			return c.JSON(500, msg)
		}
		var likes []string
		if len(rec[4]) > 0 {
			likes = strings.Split(rec[4], ",")
		} else {
			likes = make([]string, 0)
		}
		b := &models.Books{UserID: rec[0], Title: rec[1], Story: rec[2], PublishedDate: t, Likes: likes, LikeCount: len(likes)}
		books = append(books, *b)
	}
	//---------------------------------------------------//
	if controller.container.GetEnv() == "develop" {
		collection := controller.container.GetDB().Db.Collection("books")
		var bi []interface{}
		for _, b := range books {
			bi = append(bi, b)
		}
		_, err = collection.InsertMany(c.Request().Context(), bi)
		if err != nil {
			msg := models.Response{Message: err.Error()}
			return c.JSON(500, msg)
		}
	} else {
		go kafkaDocker.Produce(context.Background(), books)
	}

	//-----------------------------------------//
	return c.JSON(200, models.SuccessfulUpload{Message: "Successfully uploaded file"})
}

func (controller *BookController) TopBooks(c echo.Context) (err error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "likecount", Value: -1}})
	cursor, err := controller.container.GetDB().Db.Collection("books").Find(c.Request().Context(), bson.D{}, findOptions)
	if err != nil {
		msg := models.Response{Message: err.Error()}
		return c.JSON(500, msg)
	}
	var books = make([]models.Books, 0)
	if err := cursor.All(c.Request().Context(), &books); err != nil {
		msg := models.Response{Message: err.Error()}
		return c.JSON(500, msg)

	}
	return c.JSON(http.StatusOK, books)

}
func (controller *BookController) Like(c echo.Context) (err error) {
	userReq := new(User)
	if err = c.Bind(userReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var get string
	if controller.container.GetEnv() == "develop" {
		get = "http://localhost:1323/user/" + userReq.Id
	} else {
		get = "http://user-microservice:1323/user/" + userReq.Id
	}
	res, err := http.Get(get)
	if err != nil {
		msg := models.Response{Message: "user server down"}
		return c.JSON(500, msg)
	}
	if res.StatusCode == 404 {
		msg := models.Response{Message: "incorrect user id"}
		return c.JSON(404, msg)
	}

	collection := controller.container.GetDB().Db.Collection("books")
	objID, err := primitive.ObjectIDFromHex(userReq.Story)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	likedBookMongo := collection.FindOne(c.Request().Context(), filter)
	likedBook := &models.Books{}
	err = likedBookMongo.Decode(likedBook)
	if err != nil {
		msg := models.Response{Message: err.Error()}
		return c.JSON(400, msg)
	}

	userInLike := contains(likedBook.Likes, userReq.Id)
	if !userInLike {
		likedBook.Likes = append(likedBook.Likes, userReq.Id)
		likedBook.LikeCount += 1
	} else {
		likedBook.Likes = remove(likedBook.Likes, userReq.Id)
		likedBook.LikeCount -= 1
	}
	update := bson.D{{"$set", bson.D{{"likes", likedBook.Likes}, {"likecount", likedBook.LikeCount}}}}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	likedBookMongo = collection.FindOneAndUpdate(c.Request().Context(), filter, update, &opt)
	err = likedBookMongo.Decode(likedBook)
	if err != nil {
		msg := models.Response{Message: err.Error()}
		return c.JSON(400, msg)
	}
	return c.JSON(200, likedBook)
}
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
