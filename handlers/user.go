package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Login(w http.ResponseWriter,r *http.Request, ps httprouter.Params)  {
	account := ps.ByName("account")
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