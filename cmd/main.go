package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/zeurd/location/pkg/db"
)

type app struct {
	logger         *log.Logger
	errorLog       *log.Logger
	locationMapper *db.LocationMapper
}

func main() {
	addr := envVar("HISTORY_SERVER_LISTEN_ADDR", ":8080")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mapper := db.NewLocationMapper()

	app := &app{
		errorLog:  errorLog,
		logger:    infoLog,
		locationMapper: mapper,
	}

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	
	ttl := envVar("LOCATION_HISTORY_TTL_SECONDS", "60")
	limit, err := strconv.Atoi(ttl)
	if ttl != "" && err != nil {
		errorLog.Fatal(err)
	}
	if ttl != "" {
		go mapper.Clean(int64(limit), 180)
	}

	infoLog.Printf("start server on %s", addr)
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}

}

func envVar(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
