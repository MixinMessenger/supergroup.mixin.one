package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/MixinNetwork/supergroup.mixin.one/config"
	"github.com/MixinNetwork/supergroup.mixin.one/durable"
	"github.com/MixinNetwork/supergroup.mixin.one/services"
)

func main() {
	service := flag.String("service", "http", "run a service")
	dir := flag.String("dir", "./config/", "config.yaml dir")
	flag.Parse()

	config.LoadConfig(*dir)
	cfg := config.AppConfig.Database
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panicln(err)
	}
	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxOpenConns(128)
	db.SetMaxIdleConns(4)
	defer db.Close()

	database, err := durable.NewDatabase(context.Background(), db)
	if err != nil {
		log.Panicln(err)
	}

	switch *service {
	case "http":
		log.Println("Http Server Listened Port:", config.AppConfig.Service.HTTPListenPort)
		err := StartServer(database)
		if err != nil {
			log.Println(err)
		}
	default:
		go func() {
			hub := services.NewHub(database)
			err := hub.StartService(*service)
			if err != nil {
				log.Println(err)
			}
		}()
		http.ListenAndServe(fmt.Sprintf(":%d", config.AppConfig.Service.HTTPListenPort+2000), http.DefaultServeMux)
	}
}
