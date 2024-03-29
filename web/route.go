package web

import (
	"./controllers"
	"chatbot/web/controllers/chatbot/kakao"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Post",
		"GET",
		"/",
		controllers.Index,
	},
	Route{
		"kakaoComponent",
		"GET",
		"/chatbot/kakao",
		kakao.Post,
	},
}
