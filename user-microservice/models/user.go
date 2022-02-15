package models

type User struct {
	ID       string `json:"id,omitempty"  bson:"_id,omitempty" example:"462784683264287647223"`
	Email    string `json:"email" validate:"required,email" example:"sarthak@gmail.com"`
	Password string `json:"password" validate:"pass" example:"Test@1234000"`
}
