package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)


type IDBConnector interface {
	Connect(addr string) bool
	Close()
}

type dbconnector struct {
	addr   string
	client mongo.Client
}
var connector *dbconnector

func InitMongoConnector(addr string)  {
	connector = &dbconnector{
		addr: addr,
	}
}

func (c dbconnector) Connect(addr string) bool {
	clientOptions := options.Client().ApplyURI(c.addr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		//utils.Zapper.Error(err.Error())
		log.Fatal(err)
		return false
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		//utils.Zapper.Error(err.Error())
		log.Fatal(err)
		return false
	}
	//utils.Zapper.Info("Connect to MongoDB successed!")
	log.Println("Connect to MongoDB successed!")
	return true
}

func (c dbconnector) Close() {
	err := c.client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	//utils.Zapper.Info("Disconnect to MongoDB successed!")
	log.Println("Disconnect MongoDB successed!")
}

func GetDBConnector()*IDBConnector  {
	if connector == nil {
		InitMongoConnector("")
	}
	return connector
}