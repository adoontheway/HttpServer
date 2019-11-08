package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type IDBConnector interface {
	Connect() bool
	Close()
}

type dbconnector struct {
	addr string
	client mongo.Client
}

func NewDBConnector(addr string) IDBConnector {
	return &dbconnector{
		addr:addr,
	}
}

func (c dbconnector)Connect() bool {
	clientOptions := options.Client().ApplyURI(c.addr)
	client,err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	log.Println("Connect to MongoDB successed!")
	return true
}

func (c dbconnector)Close() {
	err := c.client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnect MongoDB successed!")
}
