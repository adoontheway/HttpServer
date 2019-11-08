package HttpServer

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type IHttpServer interface {
	AddHandler(addr string, handler httprouter.Handle)
	Start()
	Stop()
}

type httpServer struct {
	router *httprouter.Router
	addr string
}

func NewHttpServer(addr string) IHttpServer  {
	return &httpServer{
		router:httprouter.New(),
		addr:addr,
	}
}

func (s *httpServer)AddHandler(pattern string, handler httprouter.Handle)  {
	s.router.GET(pattern, handler)
}

func (s *httpServer)Start()  {
	log.Fatal(http.ListenAndServe(s.addr, s.router))
}

func (s *httpServer)Stop()  {

}