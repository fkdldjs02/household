package main

import (
	"log"
	"net/http"
	"os"
	"time"

	cors "gopkg.in/gin-contrib/cors.v1"
	pprof "gopkg.in/gin-contrib/pprof.v1"

	"github.com/fkdldjs02/household/api"
	"github.com/fkdldjs02/household/conf"
	"github.com/gin-gonic/gin"
)

func runApp(mode string) {
	caseOne := conf.NewCaseOne(mode)
	env := caseOne.Env
	api.NewApp(caseOne)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	api.SetRouter(r)

	pprof.Register(r, nil)

	server := &http.Server{
		Addr:           env.GetString("Listen"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   35 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Start server\n\tport: %s\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	mode := os.Getenv("APP_ENV")

	switch mode {
	case "release":
		runApp(mode)
	case "develop":
		runApp(mode)
	default:
		runApp("release")
	}
}
