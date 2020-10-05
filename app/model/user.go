package model

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//User struct to insert to database
type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty" `
	Username    string             `json:"username" bson:"username,omitempty" binding:"required"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty" binding:"required"`
	Email       string             `json:"email" bson:"email,omitempty" binding:"required"`
	Address     string             `json:"address" bson:"address,omitempty" binding:"required"`
	Name        string             `json:"name" bson:"name,omitempty" binding:"required"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number,omitempty" binding:"required"`
}

//UserJSON struct used to show data with certain field
type UserJSON struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty" `
	Username    string             `json:"username" bson:"username,omitempty" binding:"required"`
	Email       string             `json:"email" bson:"email,omitempty" binding:"required"`
	Address     string             `json:"address" bson:"address,omitempty" binding:"required"`
	Name        string             `json:"name" bson:"name,omitempty" binding:"required"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number,omitempty" binding:"required"`
}


type Claim struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}

//HashPassword return the password hashed
func (user User) HashPassword(password string) (string, error) {
	hash, errs := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), errs
}
