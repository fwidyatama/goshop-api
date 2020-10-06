package controller

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"goshop-api/app/model"
	"goshop-api/util"
	"log"
	"net/http"
	"os"
	"time"
)

func Register(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	//check duplicate email
	var userExist model.User
	_ = userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&userExist)
	if user.Email == userExist.Email {
		c.JSON(http.StatusBadRequest, util.FailResponse(http.StatusBadRequest, "Email already registered"))
		return
	}

	//hash password
	user.Password, _ = user.HashPassword(user.Password)
	//insert new data
	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	result, _ := userCollection.InsertOne(ctx, user)
	c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success", result))
}

//Login user
func Login(c *gin.Context) {
	var userCredential model.User
	var userDb model.User

	_ = c.ShouldBindJSON(&userCredential)
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	err := userCollection.FindOne(ctx, bson.M{"email": userCredential.Email}).Decode(&userDb)
	if err != nil {
		c.JSON(http.StatusUnauthorized, util.FailResponse(http.StatusUnauthorized,
			"Wrong email, please try again"))
		return
	}

	userPassword := []byte(userCredential.Password)
	userDbPassword := []byte(userDb.Password)
	comparePassword := bcrypt.CompareHashAndPassword(userDbPassword, userPassword)

	if comparePassword != nil {
		c.JSON(http.StatusUnauthorized, util.FailResponse(http.StatusUnauthorized,
			"Wrong password, please try again"))
		return
	}
	//jwt section
	var jwtKey = []byte(os.Getenv("JWT_SECRET"))
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := model.Claim{
		Email:    userCredential.Email,
		Username: userDb.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Success",
		"Token":   tokenString,
		"Expires": expirationTime.Format("2006-01-02 3:4:5 pm"),
	})
}
