package handlers

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/adoontheway/HttpServer/db"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)


func Login(w http.ResponseWriter,r *http.Request, ps httprouter.Params)  {
	account := ps.ByName("account")
	password := ps.ByName("password")
	dbConnector := db.GetDBConnector()
	client := dbConnector.GetClient()
	collection := client.Database("gamemain").Collection("game_user")
	ctx,_ := context.WithTimeout(context.Background(), 5*time.Second)
	user := &db.User{}
	err := collection.FindOne(ctx,bson.M{"account":account,"password":password}).Decode(user)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w,"Can't find user :%v",err)
		return
	}
	//pass := ps.ByName("password")
	fmt.Fprint(w,"Welcome to login, %s!", account)
}

func Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	account := ps.ByName("account")
	fmt.Fprint(w,"Welcome to register, %s!", account)
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	fmt.Fprint(w, "Welcome to my zone.")
}