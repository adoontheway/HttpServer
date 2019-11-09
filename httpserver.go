package HttpServer

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/adoontheway/HttpServer/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type HttpConfig struct {
	Port     int32  `json:"port"`
	LogLevel int32  `json:"log_level"`
	DB       string `json:"db"`
	Redis    string `json:"redis"`
	LogPath  string `json:"log_path"`
}

type IHttpServer interface {
	AddHandler(addr string, handler httprouter.Handle)
	Start()
	Stop()
}

type httpServer struct {
	router *httprouter.Router
	addr   string
}

func ReadConfig(filepath string) (*HttpConfig, error) {
	configfile, err := os.Open(filepath)
	if err != nil {
		//log.Fatal(err)
		utils.Zapper.Error(err.Error())
		return nil, err
	}
	defer configfile.Close()

	byteValue, err := ioutil.ReadAll(configfile)
	if err != nil {
		//log.Fatal(err)
		utils.Zapper.Error(err.Error())
		return nil, err
	}
	var config HttpConfig
	json.Unmarshal(byteValue, &config)
	return &config, nil
}

func NewHttpServer(addr string) IHttpServer {
	return &httpServer{
		router: httprouter.New(),
		addr:   addr,
	}
}

func (s *httpServer) AddHandler(pattern string, handler httprouter.Handle) {
	s.router.GET(pattern, handler)
}

func (s *httpServer) Start() {
	log.Fatal(http.ListenAndServe(s.addr, s.router))
}

func (s *httpServer) Stop() {

}
