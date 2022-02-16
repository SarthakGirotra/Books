package models

type User struct {
	// example: 6034f125f5dc910af4d64484
	ID string `json:"id,omitempty"  bson:"_id,omitempty" example:"462784683264287647223"`
	// example: sarthakgirotra@gmail.com
	Email string `json:"email" validate:"required,email" example:"sarthak@gmail.com"`
	// hashed password
	// example: $2a$14$pYn/wEAQrwS3MAKJlu6.xOILAre9jWJNWjfFe4mr/PBjYi2jlc7Ty
	Password string `json:"password" validate:"pass"`
}

type UserParams struct {
	// example: sarthakgirotra@gmail.com
	Email string `json:"email" validate:"required,email" example:"sarthak@gmail.com"`
	// example: Test@A000000
	Password string `json:"password" validate:"pass" example:"Test@1234000"`
}

type Response struct {
	// example: error message
	Message string `json:"msg"`
}
