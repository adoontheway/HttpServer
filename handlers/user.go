package handlers

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/adoontheway/HttpServer/db"
	"gitlab.com/adoontheway/HttpServer/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"strings"
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
	token := genToken(user.UserId)
	redis.Send("HSET",token,"account",account)
	redis.Send("SET",user.UserId,token)
	fmt.Fprintf(w,"Welcome, %s! Here is your token:%s", account,token)
}

func Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	account := ps.ByName("account")
	password := ps.ByName("password")
	fmt.Printf("Welcome to register, %s!\n", account)
	dbConnector := db.GetDBConnector()
	client := dbConnector.GetClient()
	collection := client.Database("gamemain").Collection("game_user")
	ctx,_ := context.WithTimeout(context.Background(), 5*time.Second)
	user := &db.User{}
	num,err := collection.CountDocuments(ctx,bson.M{"account":account})
	if err == nil {
		if num == 0 {
			user.Account = account
			user.Password = password
			user.RegisterTime = primitive.NewDateTimeFromTime(time.Now())
			//user.LastLoginTime = primitive.NewDateTimeFromTime(time.Now())
			user.NickName = account
			user.RegisterIp = strings.Split(r.RemoteAddr,":")[0]
			_,err := collection.InsertOne(ctx,user)
			if err != nil {
				fmt.Fprintf(w,"Create user failed:%v",err)
			}else {
				fmt.Fprint(w, "Create user success.")
			}
		}else {
			fmt.Fprintf(w, "User Already exsits.")
		}
	}else {
		fmt.Fprintf(w,"DB Error:%v",err)
	}
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	fmt.Fprint(w, "Welcome to my zone.")
}

func genToken(userid int64) string {
	curTime := time.Now().Unix()
	h := md5.New()
	info := fmt.Sprintf("%d:%d",curTime,userid)
	io.WriteString(h,info)
	token := fmt.Sprintf("%x",h.Sum(nil))
	return token
}