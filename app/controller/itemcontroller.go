package controller

import (
	"github.com/gin-gonic/gin"
	"goshop-api/app/model"
	"goshop-api/util"
	"log"
	"net/http"
)

//AddItem used for add new item data
func AddItem(c *gin.Context) {
	var item model.Item
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	util.DB.Create(&item)
	c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success add item", item))
}

//GetItems used for get all items data
func GetItems(c *gin.Context) {
	var items []model.ItemJson
	err := util.DB.Table("items").Find(&items).Error
	if err != nil {
		log.Fatal(err.Error())
	} else {
		c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success", items))
	}
}

//GetItem used for get single item data
func GetItem(c *gin.Context) {
	var item model.ItemJson
	id := c.Param("id")
	if err := util.DB.Table("items").Where("id = ?", id).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, util.FailResponse(http.StatusNotFound, "Item not found"))
	} else {
		c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success", item))
	}
}

//DeleteItem used for delete item
func DeleteItem(c *gin.Context) {
	var item model.Item
	id := c.Param("id")
	if err := util.DB.Where("id = ?", id).Unscoped().Delete(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, util.FailResponse(http.StatusNotFound, "Item not found"))
	} else {
		c.JSON(http.StatusOK, util.MessageResponse(http.StatusOK, "Success delete item"))
	}
}

//UpdateItem used for update item data
func UpdateItem(c *gin.Context) {
	var item model.Item
	id := c.Param("id")
	if err := util.DB.Where("id = ?", id).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, util.FailResponse(http.StatusNotFound, "Item not found"))
	} else {
		c.ShouldBindJSON(&item)
		util.DB.Save(&item)
		c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success update item", item))

	}
}
