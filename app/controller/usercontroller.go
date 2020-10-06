package controller

import (
	"context"
	"time"

	//"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goshop-api/app/model"
	"goshop-api/util"
	"log"
	"net/http"
	"os"
	"strings"
)

var userCollection = util.InitializeMongoDB().Database("go-shop").Collection("users")

// GetUsers return all users
func GetUsers(c *gin.Context) {
	var users []model.UserJSON
	cursor, err := userCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
	}

	if cursor == nil {
		log.Fatal("Cursor nil")
	}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var user model.UserJSON
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err.Error())
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		c.JSON(http.StatusOK, util.FailResponse(http.StatusOK, "No data Available"))
		return
	}
	c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success", users))
}

//GetUser return single user only
func GetUser(c *gin.Context) {
	var user model.UserJSON
	userQuery := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(userQuery)
	result := userCollection.FindOne(context.TODO(), bson.M{"_id": objectID})
	_ = result.Decode(&user)
	c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success", user))
}

//DeleteUser use to delete a user
func DeleteUser(c *gin.Context) {
	userQuery := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(userQuery)
	result, err := userCollection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		log.Fatal(err.Error())
	}
	c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success", result))
}

//UpdateUser use to update user database
func UpdateUser(c *gin.Context) {
	var user model.User
	//get current id
	userQuery := c.Param("id")
	ObjectID, _ := primitive.ObjectIDFromHex(userQuery)
	filter := bson.D{{"_id", ObjectID}}

	//return document after update data
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	//find current user information and store it to user memory address
	objectID, _ := primitive.ObjectIDFromHex(userQuery)
	userResult := userCollection.FindOne(context.TODO(), bson.M{"_id": objectID})
	_ = userResult.Decode(&user)

	//bind input form postman into model user
	_ = c.ShouldBindJSON(&user)

	hashedPassword, _ := user.HashPassword(user.Password)
	//set user field based on json. Can update certain field only or all field
	update := bson.M{"$set": bson.M{
		"name":         user.Name,
		"email":        user.Email,
		"username":     user.Username,
		"password":     hashedPassword,
		"phone_number": user.PhoneNumber,
		"address":      user.Address,
	}}
	//update user data
	updateResult := userCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)
	var result bson.M
	_ = updateResult.Decode(&result)
	c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success", result))
}

//Extract token information
func ExtractJWT(tokenStr string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Fatal("Invalid JWT Token")
		return nil, false
	}
}

//GetProfile to get user detail
func GetProfile(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	extractedData, failed := ExtractJWT(tokenString)
	if failed == false {
		c.JSON(http.StatusUnauthorized, util.FailResponse(http.StatusUnauthorized, "Invalid token"))
		return
	}
	var user model.UserJSON
	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	err := userCollection.FindOne(ctx, bson.M{"email": extractedData["email"]}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.FailResponse(http.StatusBadRequest, "User not found"))
		return
	}
	c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success", user))

}
