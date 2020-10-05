package controller

import (
	"github.com/gin-gonic/gin"
	"goshop-api/app/model"
	"goshop-api/util"
	"log"
	"net/http"
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

//GetStore used for  get single store
func GetStore(c *gin.Context) {
	var store model.Store
	id := c.Param("id")
	if err := util.DB.Preload("Items").Table("stores").Where("id = ?", id).First(&store).Error; err != nil {
		c.JSON(http.StatusNotFound, util.FailResponse(http.StatusNotFound, "Store not found"))
	} else {
		c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success", store))
	}
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

//GetStoreItems used for get all store items
func GetStoreItems(c *gin.Context) {
	var stores []model.Store
	//if err := util.DB.Preload("Items").Find(&stores).Error; err != nil {
	if err := util.DB.Table("item_store").Preload("Items").Find(&stores).Error; err != nil{
		log.Fatal(err.Error())
	} else {
		c.JSON(http.StatusOK, util.SuccessResponse(http.StatusOK, "Success", stores))
	}

}
