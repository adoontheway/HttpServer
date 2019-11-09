package HttpServer

import (
	"gitlab.com/adoontheway/HttpServer/handlers"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestNewHttpServer(t *testing.T) {
	server := NewHttpServer(":8888")
	server.AddHandler("/register", handlers.Register)
	server.AddHandler("/login", handlers.Login)
	server.AddHandler("/", handlers.Index)
	server.Start()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <- sigs
		log.Println()
		log.Println("Recieve Signal :",sig)
		done <- true
	}()
	<-done
	log.Println("Application terminated")
}

func TestInitFromConfig(t *testing.T) {
	server,err := InitFromConfig("./config.json")
	if err != nil {
		log.Fatal(err)
		t.Fail()
	}
	server.AddHandler("/register", handlers.Register)
	server.AddHandler("/login", handlers.Login)
	server.AddHandler("/", handlers.Index)
	server.Start()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <- sigs
		log.Println()
		log.Println("Recieve Signal :",sig)
		done <- true
	}()
	<-done
	log.Println("Application terminated")
}
