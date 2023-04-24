package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"startwithmongo/app"
	"startwithmongo/repository"
	"startwithmongo/util/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE       = "golang"
	SERVER_ADDRESS = "SERVER_ADDRESS"
	SERVER_PORT    = "SERVER_PORT"
	MONGO_URI      = "MONGO_URI"
	MONGO_USERNAME = "MONGO_INITDB_ROOT_USERNAME"
	MONGO_PASSWORD = "MONGO_INITDB_ROOT_PASSWORD"
)

var (
	//Application Configuration
	serverAddr = "localhost"
	serverPort = "4000"
	//MongoDB Configuration
	mongoURI  = "mongodb://localhost:27071"
	mongoUser = "mongoadmin"
	mongoPass = "secret"
)

func initializeVariables() {
	servA, b := os.LookupEnv(SERVER_ADDRESS)
	if b {
		serverAddr = servA
	}
	servP, b := os.LookupEnv(SERVER_PORT)
	if b {
		serverPort = servP
	}
	mongo, b := os.LookupEnv(MONGO_URI)
	if b {
		mongoURI = mongo
	}
	mongoU, b := os.LookupEnv(MONGO_USERNAME)
	if b {
		mongoUser = mongoU
	}
	mongoP, b := os.LookupEnv(MONGO_PASSWORD)
	if b {
		mongoPass = mongoP
	}

}


func main() {
	logger.Info().Println("Starting Application")
	initializeVariables()

	// Create mongo client configuration
	co := options.Client().ApplyURI(mongoURI)
	co.Auth = &options.Credential{
		Username: mongoUser,
		Password: mongoPass,
	}
	// Establish database connection
	client, err := mongo.NewClient(co)
	if err != nil {
		logger.Error().Fatal(err)
	}
	// CAPIRE CHE PROBLEMA COMPORTA QUESTA VARIAZIONE CON IL CONTEXT SENZA TIME OUT
	//ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	//defer cancel()
	ctx := context.Background()

	err = client.Connect(ctx)
	if err != nil {
		logger.Error().Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			logger.Error().Fatal(err)
		}
	}()
	logger.Info().Printf("Database connection established")

	Datab := client.Database(DATABASE)

	(&repository.DB).Set(Datab, ctx)

	// Initialize a new instance of application containing the dependencies.
	serverURI := fmt.Sprintf("%s:%s", serverAddr, serverPort)
	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     logger.Error(),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info().Printf("Starting server on %s", serverURI)
	err = srv.ListenAndServe()
	logger.Error().Fatal(err)
}
