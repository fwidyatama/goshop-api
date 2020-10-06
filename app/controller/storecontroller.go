package controller

import (
	"github.com/gin-gonic/gin"
	"goshop-api/app/model"
	"goshop-api/util"
	"log"
	"net/http"
	"time"
)

//GetStores used for  get all store data
func GetStores(c *gin.Context) {
	var stores []model.StoreJson
	err := util.DB.Table("stores").Find(&stores).Error
	if err != nil {
		log.Fatal(err.Error())
	} else {
		c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success", stores))
	}
}

//GetStore used for get store detail and their items
func GetStore(c *gin.Context) {
	var result model.Result
	id := c.Param("id")
	rows, err := util.DB.Raw("SELECT stores.id, stores.name, items.id, items.name, items.price, store_items.quantity, store_items.created_at FROM stores join store_items on stores.id = store_items.store_id join items on items.id=store_items.item_id where stores.id=?", id).Rows()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var createdAt time.Time
		var IdStore, IdItem, itemPrice, itemQuantity int
		var storeName, itemName string
		rows.Scan(&IdStore, &storeName, &IdItem, &itemName, &itemPrice, &itemQuantity, &createdAt)
		result.StoreID = IdStore
		result.StoreName = storeName
		result.ItemDetail = append(result.ItemDetail, model.ItemDetail{
			ItemID:    IdItem,
			ItemName:  itemName,
			Price:     itemPrice,
			Quantity:  itemQuantity,
			CreatedAt: createdAt,
		})
	}
	c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success", result))
}

//AddStore used for create new store
func AddStore(c *gin.Context) {
	var store model.Store
	err := c.ShouldBindJSON(&store)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	util.DB.Create(&store)
	c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success add store", store))
}

//DeleteStore used for delete a store form database
func DeleteStore(c *gin.Context) {
	var store model.Store
	id := c.Param("id")
	if err := util.DB.Where("id = ? ", id).First(&store).Unscoped().Delete(&store).Error; err != nil {
		c.JSON(http.StatusNotFound, util.FailResponse(http.StatusNotFound, "Store not found"))
		return
	} else {
		c.JSON(http.StatusOK, util.MessageResponse(http.StatusOK, "Success delete store"))
	}
}

//UpdateStore used for update store data
func UpdateStore(c *gin.Context) {
	var store model.Store
	id := c.Param("id")
	if err := util.DB.Where("id = ? ", id).First(&store).Error; err != nil {
		c.JSON(http.StatusNotFound, util.FailResponse(http.StatusNotFound, "Store not found"))
		return
	} else {
		c.ShouldBindJSON(&store)
		util.DB.Save(&store)
		c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success update store", store))
	}
}

//AddStoreItem used for add items owned by a store
func AddStoreItem(c *gin.Context) {
	var items []model.StoreItem
	c.ShouldBindJSON(&items)
	if len(items) == 0 {
		c.JSON(http.StatusBadRequest, util.MessageResponse(http.StatusBadRequest, "Please input item"))
		return
	}
	util.DB.Model(model.StoreItem{}).Create(&items)
	c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success add store", items))
}
