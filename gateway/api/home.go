package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type homeApiHandle struct{}

func RegisterHomeApiHandle(router *mux.Router) *homeApiHandle {
	handler := &homeApiHandle{}
	router.HandleFunc("/api", handler.home)
	return handler
}

func (*homeApiHandle) home(res http.ResponseWriter, req *http.Request) {
	enableCors(res)
	res.Write([]byte("This home of api"))
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
