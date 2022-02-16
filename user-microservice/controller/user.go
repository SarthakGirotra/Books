package controller

import (
	"net/http"
	"t/container"

	"t/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	container container.Container
}

func NewUserController(container container.Container) *UserController {
	return &UserController{container: container}
}

// GetAllUsers - dev function to fetch all users
func (controller *UserController) GetAllUsers(c echo.Context) (err error) {

	query := bson.D{{}}
	cursor, err := controller.container.GetDB().Db.Collection("users").Find(c.Request().Context(), query)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	var users = make([]models.User, 0)
	if err := cursor.All(c.Request().Context(), &users); err != nil {
		return c.JSON(500, err.Error())

	}
	return c.JSON(http.StatusOK, users)
}

// Login - Returns logged in user details
func (controller *UserController) Login(c echo.Context) (err error) {
	collection := controller.container.GetDB().Db.Collection("Users")
	user := new(models.User)
	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(user); err != nil {
		return err
	}
	filter := bson.D{{Key: "email", Value: user.Email}}
	checkExistM := collection.FindOne(c.Request().Context(), filter)
	checkExist := &models.User{}
	_ = checkExistM.Decode(checkExist)

	if checkExist.ID == "" {
		return c.JSON(http.StatusNotFound, models.Response{Message: "User not found"})
	} else {
		if checkPasswordHash(user.Password, checkExist.Password) {
			return c.JSON(http.StatusOK, checkExist)
		} else {
			return c.JSON(http.StatusUnauthorized, models.Response{Message: "Incorrect Password"})
		}
	}
}

// Signup - Returns signed up user details
func (controller *UserController) Signup(c echo.Context) (err error) {
	collection := controller.container.GetDB().Db.Collection("Users")

	user := new(models.User)
	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(user); err != nil {
		return err
	}
	filter := bson.D{{Key: "email", Value: user.Email}}
	checkExistM := collection.FindOne(c.Request().Context(), filter)
	checkExist := &models.User{}
	_ = checkExistM.Decode(checkExist)

	if checkExist.ID != "" {
		msg := &models.Response{Message: "User Already Exists"}
		return c.JSON(http.StatusUnauthorized, msg)
	}
	user.Password, _ = hashPassword(user.Password)
	res, err := collection.InsertOne(c.Request().Context(), user)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	filterS := bson.D{{Key: "_id", Value: res.InsertedID}}
	newUserMongo := collection.FindOne(c.Request().Context(), filterS)
	newUser := &models.User{}
	if err = newUserMongo.Decode(newUser); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	return c.JSON(201, newUser)
}

// ValidateUser - validates user based on url param
func (controller *UserController) ValidateUser(c echo.Context) (err error) {
	id := c.Param("id")
	objid, _ := primitive.ObjectIDFromHex(id)
	collection := controller.container.GetDB().Db.Collection("Users")
	filterS := bson.D{{Key: "_id", Value: objid}}
	mongoUser := collection.FindOne(c.Request().Context(), filterS)
	u := &models.User{}
	_ = mongoUser.Decode(u)

	if u.ID == "" {
		return c.JSON(http.StatusNotFound, "User Doesn't Exist")
	} else {
		return c.JSON(http.StatusOK, "User Exists")
	}
}
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
