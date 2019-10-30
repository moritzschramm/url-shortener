package main

import (

	"net/http"
	"log"
	"database/sql"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8484"
)

func StartServer(db *sql.DB) {

	router := httprouter.New()

	SetupRoutes(router, db)

	server := negroni.New()
	server.Use(negroni.NewLogger())
	server.Use(negroni.NewRecovery())
	server.UseHandler(router)

	log.Println("Starting server on http://" + SERVER_HOST + ":" + SERVER_PORT)
	log.Fatal(http.ListenAndServe(":"+SERVER_PORT, server))
}

func SetupRoutes(router *httprouter.Router, db *sql.DB) {

	router.NotFound = http.HandlerFunc( func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "404.html")
	})

	// loop through links
	links, err := GetLinks(db)
	if err != nil {
		log.Fatal("Error while setting up routes: " + err.Error())
	}

	for _, link := range links {

		if !strings.HasPrefix(link.ShortUrl, "/") {

			link.ShortUrl = "/" + link.ShortUrl
		}

		router.GET(link.ShortUrl, redirectTo(link.TargetUrl))
	}
}

func redirectTo(targetUrl string) httprouter.Handle {

	return func(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		http.Redirect(res, req, targetUrl, 302)
	}
}

