package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(viper.GetString("URI"))

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// 连接到自己的MongoDB
	collection := client.Database(viper.GetString("DBName")).Collection(viper.GetString("DBCollection"))

	// 插入操作
	res1, err := collection.InsertOne(context.TODO(), bson.M{"key123": "123"})
	if err != nil {
		logrus.Info("There was a problem with the insert operation")
	}

	// 插入操作
	res2, err := collection.InsertOne(context.TODO(), bson.M{"key123": "234"})
	if err != nil {
		logrus.Info("There was a problem with the insert operation")
	}

	// 删除操作
	_, err = collection.DeleteOne(context.TODO(), bson.D{{"_id", res1.InsertedID}})
	if err != nil {
		logrus.Info("There was a problem with the delete operation")
		log.Fatal(err)
	}

	// 查找操作
	var result bson.M
	err = collection.FindOne(context.TODO(), bson.D{{"_id", res2.InsertedID}}).Decode(&result)
	if err != nil {
		logrus.Info("There was a problem with the find operation")
	}
	fmt.Println(result)

	// 修改操作
	res, err := collection.UpdateByID(context.Background(), res2.InsertedID, bson.D{{"$set", bson.D{{"key123", 12}}}})
	if err != nil {
		// log.Fatal(err)
		logrus.Info("There was a problem with the update operation")
	}
	fmt.Println(res)

}
