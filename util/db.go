package util

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"goshop-api/app/model"
	"log"
)

var DB *gorm.DB

func InitializeSqlDB(DbHost, DbPort, DbUser, DbName, DbPassword string) *gorm.DB{
	DbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",DbUser,DbPassword,DbHost,DbPort,DbName)
	Database,err:=gorm.Open(mysql.Open(DbUri),&gorm.Config{})
	if err!=nil{
		panic(err.Error())
	}
	DB = Database

	////Auto migrate
	DB.AutoMigrate(&model.Store{},&model.Item{})
	//
	//join table
	//result := DB.SetupJoinTable(&model.Store{}, "Items", &model.ItemStore{})
	//if result !=nil{
	//	log.Fatal(result)
	//}
	////
	//migrate
	//Db.AutoMigrate(&model.Store{})
	//store := &model.Store{
	//	Name: "Toko otong",
	//	Owner: "Andromeda",
	//}
	//Db.Create(store)
	//
	//Db.AutoMigrate(&model.Store{})
	//store2 := &model.Store{
	//	Name: "Toko Tokoan",
	//	Owner: "Andromeda",
	//}
	//Db.Create(store2)
	return DB
}

func InitializeMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
