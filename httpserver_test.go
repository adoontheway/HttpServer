package HttpServer

import (
	"fmt"
	"gitlab.com/adoontheway/HttpServer/db"
	"gitlab.com/adoontheway/HttpServer/handlers"
	"gitlab.com/adoontheway/HttpServer/redis"
	"log"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestNewHttpServer(t *testing.T) {
	server := NewHttpServer(":8888")
	server.AddHandler("/register",GET, handlers.Register)
	server.AddHandler("/login", GET,handlers.Login)
	server.AddHandler("/", GET,handlers.Index)
	go server.Start()
	//go func() {
	//	log.Println(http.ListenAndServe(":6060",nil))
	//}()
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Println()
		log.Println("Recieve Signal :", sig)
		done <- true
	}()
	<-done
	log.Println("Application terminated")
}

func TestInitFromConfig(t *testing.T) {

	config, err := ReadConfig("./config.json")
	if err != nil {
		//utils.Zapper.Error(err.Error())
		log.Fatal(err)
		t.Fail()
	}

	db.InitMongoConnector(config.DB)
	redis.InitRedisPool(config.Redis, config.RedisPass)

	server := NewHttpServer(fmt.Sprintf(":%d", config.Port))
	server.AddHandler("/register/:account/:password",GET, handlers.Register)
	server.AddHandler("/login/:account/:password", GET,handlers.Login)
	server.AddHandler("/", GET,handlers.Index)
	go server.Start()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Println()
		log.Println("Recieve Signal :", sig)
		done <- true
	}()
	<-done
	//utils.Zapper.Info("Application terminated")
	log.Println("Application terminated")
}
