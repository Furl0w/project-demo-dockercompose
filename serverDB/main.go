package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3031"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Initiating connection with DB\n")
		err := pingDBClient()
		if err != nil {
			fmt.Fprintf(w, err.Error())
			fmt.Fprintf(w, "Failed to connect to DB")
		} else {
			fmt.Fprintf(w, "Succeeded to connect to DB")
		}
		return
	})

	http.ListenAndServe(":"+PORT, nil)
}

func pingDBClient() error {
	var DBPORT string
	if DBPORT = os.Getenv("DBPORT"); DBPORT == "" {
		DBPORT = "27017"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, "mongodb://mongo:"+DBPORT)
	if err != nil {
		return err
	}
	return client.Ping(ctx, readpref.Primary())
}
