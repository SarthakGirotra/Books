package models

import "time"

// swagger:model Books
type Books struct {
	// id of book
	// example: 620aa5e18e9c117fdbb9f4d2
	ID string `json:"id,omitempty"  bson:"_id,omitempty"`
	// title of book
	// example: abc
	Title string `json:"title"`
	// story of book
	// example: lorem ipsum
	Story string `json:"story"`
	// likes array
	// example: ["620aa5e18e9c117fdbb9f4d2","620aa5e18e9c117fdbb9f4d3"]
	Likes []string `json:"likes"`
	// id of user
	// example: 620aa5e18e9c117fdbb9f4d2
	UserID string `json:"userid"`
	// date of upload
	// example: 2022-02-02T15:04:00Z
	PublishedDate time.Time `json:"published_date"`
	// example: 1
	LikeCount int `json:"likecount"`
}

type Response struct {
	// example: error message
	Message string `json:"msg"`
}
type SuccessfulUpload struct {
	// example: successfully uploaded
	Message string `json:"msg"`
}
