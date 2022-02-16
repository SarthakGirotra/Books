// Package classification User Microservice
//
// Documentation of user microservice
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: localhost:1323
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package docs

import (
	"t/middlewareLocal"
	"t/models"
)

// swagger:route POST /login user idOfLoginEndpoint
// API endpoint for logging in.
// responses:
//   200: User
//	 404: notFound
//   401: incorrect
// 	 400: invalid

// swagger:route POST /signup user idOfSignupEndpoint
// API endpoint for signing up.
// responses:
// 200: User
// 401: exists
// 400: invalid

// Login Response.
// swagger:response User
type LoginResponseWrapper struct {
	// in:body
	Body models.User
}

// User not found.
// swagger:response notFound
type UserNotFound struct {
	// in:body
	Body models.Response
}

// Incorrect password.
// swagger:response incorrect
type IncorrectPassword struct {
	// in:body
	Body models.Response
}

// Invalid values.
// swagger:response invalid
type InvalidValues struct {
	// in:body
	Body []middlewareLocal.ApiError
}

// swagger:parameters idOfLoginEndpoint idOfSignupEndpoint
type LoginParamsWrapper struct {
	// valid email and password.
	// in:body
	Body models.UserParams
}

// User Already Exists
// swagger:response exists
type AlreadyExists struct {
	// in:body
	Body models.Response
}
