package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Login(w http.ResponseWriter,r *http.Request, ps httprouter.Params)  {
	account := ps.ByName("account")
	//pass := ps.ByName("password")
	fmt.Fprint(w,"Welcome, %s!", account)
}

func Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	account := ps.ByName("account")
	fmt.Fprint(w,"Welcome, %s!", account)
}