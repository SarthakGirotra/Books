// Package classification Books Microservice
//
// Documentation of books microservice
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: localhost:1322
//
//     Consumes:
//     - application/json
//	   - multipart/form-data
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package docs

import (
	"b/controllers"
	"b/models"
	"mime/multipart"
)

// swagger:route GET /topBooks books booksEndPointId
// API endpoint for getting books sorted desc by likes.
// responses:
//	 200: Books

// swagger:route POST /Like books likeBooksID
// API endpoint for liking/unliking a book.
// responses:
//	 200: Like
// 	 400: err
// 	 500: err
//   404: err
// Top Books sorted by like desc.
// swagger:response Books

// swagger:route POST /saveCSV books saveCSVID
// API endpoint for saving csv to db.
// Consumes:
//     - multipart/formdata
// responses:
//    200: upload
//    500: err

type TopBooksResponseWrapper struct {
	// in:body
	Body []models.Books
}

// Unliked/Liked Book
// swagger:response Like
type LikeBookResponseWrapper struct {
	// in:body
	Body models.Books
}

// swagger:parameters likeBooksID
type LikeBooksParamsWrapper struct {
	// valid story id and user id
	// in:body
	Body controllers.User
}

// Error Response
// swagger:response err
type ErrorResponseWrapper struct {
	// error response
	// in:body
	Body models.Response
}

// swagger:parameters saveCSVID
type CSVParamsWrapper struct {
	// csv file to be uploaded
	// in: formData
	// name: MyFile
	// required: true
	// type: file
	// example: id, title, story, date, likes(comma seperated userid array) (order of columns)
	// swagger:file
	MyFile *multipart.FileHeader
}

// upload response
// swagger:response upload
type SuccessFulUploadWrapper struct {
	// successful upload
	// in:body
	Body models.SuccessfulUpload
}
