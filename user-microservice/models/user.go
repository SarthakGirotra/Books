package models

type User struct {
	// example: 6034f125f5dc910af4d64484
	ID string `json:"id,omitempty"  bson:"_id,omitempty" example:"462784683264287647223"`
	// example: sarthakgirotra@gmail.com
	Email string `json:"email" validate:"required,email" example:"sarthak@gmail.com"`
	// example: Test@A000000
	Password string `json:"password" validate:"pass" example:"Test@1234000"`
}

type UserParams struct {
	// example: sarthakgirotra@gmail.com
	Email string `json:"email" validate:"required,email" example:"sarthak@gmail.com"`
	// example: Test@A000000
	Password string `json:"password" validate:"pass" example:"Test@1234000"`
}

type Response struct {
	// example: User not found || Incorrect Password || user already exists
	Message string `json:"msg"`
}
