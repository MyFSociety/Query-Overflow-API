package models

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}

type LoginUser struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type SignedDetails struct {
	Email string `json:"email,omitempty" validate:"required"`
	Name  string `json:"name,omitempty" validate:"required"`
	Id    string `json:"id,omitempty"`
	jwt.StandardClaims
}
