// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package controllers

import (
	"github.com/SahaPratik6267/instagramscrapper/ScrapperProject/pkg/models" // Google OpenID client
	"github.com/google/uuid"
)

//swagger:parameters login loginparam
type CredentialWrapper struct {
	//in:body
	Credentials Credentials `json:"credentials"`
}

//swagger:response IdResponse
type tokenwrapper struct {
	//in: body
	//required: true
	ID uuid.UUID `json:"id"`
}

//swagger:parameters registeruser registerparam
type RegisterWrapper struct {
	//in: body
	//required: true
	RegisterUserDetails models.User
}

//swagger:parameters twitterProfile
type ProfileNamewrapper struct {
	//in:body
	//required: true
	InputProfile ProfileName `json:"profile"`
}
