package models

import "time"

type Books struct {
	ID            string    `json:"id,omitempty"  bson:"_id,omitempty"`
	Title         string    `json:"title"`
	Story         string    `json:"story"`
	Likes         []string  `json:"likes"`
	UserID        string    `json:"userid"`
	PublishedDate time.Time `json:"published_date"`
	LikeCount     int       `json:"likecount"`
}
