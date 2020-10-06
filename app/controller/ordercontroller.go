package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"goshop-api/app/model"
	"goshop-api/util"
	log "log"
	"net/http"
	"strings"
	"time"
)

var orderCollection = util.InitializeMongoDB().Database("go-shop").Collection("orders")

func GetDetailUser(c *gin.Context) (detail string) {
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
	return user.ID.Hex()
}

func AddOrder(c *gin.Context) {
	//struct for receive user input
	type InputOrder struct {
		ItemId   int
		StoreId  int
		Quantity int
	}
	var inputOrder []InputOrder
	var order model.Order
	c.ShouldBindJSON(&inputOrder)

	for _, value := range inputOrder {
		//used for get detail item and append to item.OrderItem struct
		rows, _ := util.DB.Raw("SELECT stores.id, stores.name, items.id, items.name, items.price, store_items.quantity FROM stores join store_items on stores.id = store_items.store_id join items on items.id=store_items.item_id where stores.id=? and items.id = ?", value.StoreId, value.ItemId).Rows()
		for rows.Next() {
			var IdStore, IdItem, itemPrice, itemQuantity int
			var storeName, itemName string
			rows.Scan(&IdStore, &storeName, &IdItem, &itemName, &itemPrice, &itemQuantity)
			order.OrderItem = append(order.OrderItem, model.OrderItem{
				ItemName:  itemName,
				ItemId:    IdItem,
				StoreName: storeName,
				StoreId:   IdStore,
				Quantity:  value.Quantity,
			})
		}

		//change quantity after purchasing
		if value.Quantity > 0 {
			util.DB.Model(&model.StoreItem{}).Where("store_id = ? And item_id = ?", value.StoreId, value.ItemId).UpdateColumn("quantity", gorm.Expr("quantity - ?", value.Quantity))
		}
	}
	//get detail user and insert into database
	detailUser := GetDetailUser(c)
	model.NewOrder()
	UserId, _ := primitive.ObjectIDFromHex(detailUser)
	order.OrderDate = time.Now()
	order.UserID = UserId
	result, _ := orderCollection.InsertOne(context.TODO(), &order)
	c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success order", result))
}

func GetOrders(c *gin.Context) {
	var Orders []model.Order
	cursor, err := orderCollection.Find(context.TODO(), bson.M{})
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
		var order model.Order
		err := cursor.Decode(&order)
		if err != nil {
			log.Fatal(err.Error())
		}

		Orders = append(Orders, order)
	}
	if len(Orders) == 0 {
		c.JSON(http.StatusOK, util.FailResponse(http.StatusOK, "No data Available"))
		return
	}
	c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success", Orders))
}
