package router

import (
	"log"
	"net/http"

	"time"

	"github.com/titusdmoore/goCms/internal/config"
)

type Router struct {
	mux    *http.ServeMux
	routes []Route
}

type Route struct {
	route   string
	handler func(http.ResponseWriter, *http.Request)
}

func (router *Router) RegisterRoute(route string, handler func(http.ResponseWriter, *http.Request)) {
	router.routes = append(router.routes, Route{
		route:   route,
		handler: handler,
	})
}

func (router *Router) Serve(config config.Config) {
	// Register the routes to the net router mux
	for _, route := range router.routes {
		router.mux.HandleFunc(route.route, route.handler)
	}

	server := http.Server{
		Addr:    config.Router.Port,
		Handler: HandlerWithLogging(router.mux),
	}

	log.Printf("Server listening on port %s\n", config.Router.Port)

	server.ListenAndServe()
}

func HandlerWithLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		next.ServeHTTP(writer, request)
		log.Println(request.Method, request.URL.Path, time.Since(start))
	})
}

func InitializeRouter() (Router, error) {
	mux := http.NewServeMux()
	var routes []Route

	router := Router{
		mux:    mux,
		routes: routes,
	}

	return router, nil
}
