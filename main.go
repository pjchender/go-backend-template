package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pjchender/go-backend-template/internal/grpc"
	"github.com/pjchender/go-backend-template/internal/router"
	log "github.com/sirupsen/logrus"

	"github.com/pjchender/go-backend-template/global"
	"github.com/pjchender/go-backend-template/internal/database"
	"github.com/pjchender/go-backend-template/pkg/setup"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

func init() {
	var err error
	err = setup.Logger()
	if err != nil {
		log.Fatalf("init.setupLogger failed: %v", err)
	}

	// setupEnv should invoke before setupSetting()
	err = setup.Env()
	if err != nil {
		log.Fatalf("init.setupEnv failed: %v", err)
	}

	err = setup.Settings()
	if err != nil {
		log.Fatalf("init.setupSetting failed: %v", err)
	}
}

func main() {
	// init database
	db, err := database.New(global.DatabaseSetting.DSN, global.GormSetting)
	if err != nil {
		log.Fatalf("[main] database.New failed: %v", err)
	}
	db.AutoMigrate()

	// start gin server
	engine := router.New(db)
	go startHTTPServer(engine)
	go startGRPCServer(db)

	// shut down server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")
}

func startHTTPServer(engine *gin.Engine) {
	var httpHandler http.Handler = engine

	// start tls https connection
	go func() {
		isSSLEnabled := *global.HTTPServerSetting.SSL.Enabled
		if !isSSLEnabled {
			return
		}

		addrTLS := fmt.Sprintf("%s:%d", global.HTTPServerSetting.SSL.ListenAddr, global.HTTPServerSetting.SSL.Port)
		fmt.Println("[main] Started Listening for TLS connection on " + addrTLS)

		err := http.ListenAndServeTLS(addrTLS, global.HTTPServerSetting.SSL.CertFile,
			global.HTTPServerSetting.SSL.CertKey,
			httpHandler)
		if err != nil {
			log.Fatalf("[main] run - http.ListenAndServeTLS failed: %v", err)
		}
	}()

	// start plain http connection
	go func() {
		serverPort, err := strconv.Atoi(global.HTTPServerSetting.Port)
		if err != nil {
			log.Fatal("[main] run - strconv.Atoi failed: ", err)
		}

		addr := fmt.Sprintf("%s:%d", global.HTTPServerSetting.ListenAddr, serverPort)

		server := http.Server{
			Addr:    addr,
			Handler: httpHandler,
		}
		err = server.ListenAndServe()
		if err != nil {
			log.Fatalf("[main] run - http.ListenAndServe failed: %v", err)
		}

		fmt.Println("Started Listening for plain HTTP connection on " + addr)
	}()
}

func startGRPCServer(db *database.GormDatabase) {
	var err error
	serverPort, err := strconv.Atoi(global.GRPCSetting.Server.Port)
	if err != nil {
		log.Fatal("[main] run - strconv.Atoi failed: ", err)
	}

	fmt.Printf("[main] gRPC server started at %v\n", global.GRPCSetting.Server.Port)
	grpcAddr := fmt.Sprintf("%s:%d", global.GRPCSetting.Server.ListenAddr, serverPort)
	err = grpc.Serve(db, grpcAddr, global.GRPCSetting.Server.CertFile, global.GRPCSetting.Server.CertKey)
	if err != nil {
		log.Fatalln("[main] grpc.Serve failed: ", err)
	}
}
