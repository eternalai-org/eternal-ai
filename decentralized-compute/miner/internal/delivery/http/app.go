package http

import (
	"fmt"
	"net/http"
	"solo/internal/port"
	"solo/pkg/logger"

	"strings"

	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type httpDelivery struct {
	usecase port.IApi
	router  *mux.Router
	timeout int //seconds
	port    int
}

func NewHttp(uc port.IApi, port int) (*httpDelivery, error) {
	r := mux.NewRouter()
	h := new(httpDelivery)
	h.usecase = uc
	h.router = r
	h.port = port
	h.timeout = 86400
	return h, nil
}

func (h *httpDelivery) Run() {
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "X-Requested-With", "param"})
	h.registerRoutes()
	credentials := handlers.AllowCredentials()
	hCORS := handlers.CORS(credentials, methods, origins, headers)(h.router)
	serverPort := fmt.Sprintf(":%d", h.port)
	timeOut := h.timeout
	srv := &http.Server{
		Handler: handlers.CompressHandler(hCORS),
		Addr:    serverPort,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Duration(timeOut) * time.Second,
		ReadTimeout:  time.Duration(timeOut) * time.Second,
	}

	logger.AtLog.Info(fmt.Sprintf("Server is listening at port %s ...", serverPort))
	if err := srv.ListenAndServe(); err != nil {
		logger.AtLog.Error("httpDelivery.StartServer - Can not start http server", err)
	}
}

func (h *httpDelivery) printRoutes() {
	fmt.Println("Available routers: ")
	r := h.router
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		txt := ""

		tpl, err1 := route.GetPathTemplate()
		if err1 == nil {
			txt += tpl
		}

		met, err2 := route.GetMethods()
		if err2 == nil {
			txt += " [" + strings.Join(met, ", ") + "]"
		}

		fmt.Println(" - ", txt)

		return nil
	})
}
