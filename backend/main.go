package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jason2924/scanner/backend/config"
	"github.com/Jason2924/scanner/backend/databases"
	"github.com/Jason2924/scanner/backend/handlers"
	"github.com/Jason2924/scanner/backend/server"
)

func main() {
	// get config from env file
	conf, erro := config.Load("./", ".env.local", "env")
	if erro != nil {
		log.Fatalln("Failed occured while setting config", erro)
	}
	fmt.Println(conf)
	// connect database
	mysqlDtbs := databases.NewMysqlDatabase(&conf.Mysql)
	mysqlDtbs.Connect()
	if erro := mysqlDtbs.Ping(context.Background()); erro != nil {
		log.Fatalln("Failed occured while connecting mysql database", erro)
	}
	// create route
	ngin := handlers.Initialize(mysqlDtbs)
	// set graceful shutdown
	serv := server.NewServer(conf.Port, ngin)
	sigChan := make(chan os.Signal, 1)
	// create the background and listen and serve
	go func() {
		if erro := serv.Start(); erro != nil && erro != http.ErrServerClosed {
			log.Fatalln("Failed occured while starting server", erro)
		}
	}()
	// the signal channel to listen the Interrupt and Termination signals
	// SIGINT = Interrupt | SIGTERM = Termination
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down...")
	// set timeout for closing all connection
	ctxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// close all connection
		// cron.Stop()
		// cach.Close()
		// dtbs.Close()
		// log.Finish()
		cancel()
	}()
	// shutdown the server
	if erro := serv.Stop(ctxt); erro == context.DeadlineExceeded {
		log.Println("Halted active connections")
	}
	close(sigChan)
}
