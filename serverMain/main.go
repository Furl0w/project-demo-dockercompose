package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World from path: %s\n", r.URL.Path)
		return
	})

	http.HandleFunc("/checkDB", checkDB)

	http.ListenAndServe(":"+PORT, nil)
}

func checkDB(w http.ResponseWriter, r *http.Request) {
	var dbServicePort, dbServiceName string
	if dbServicePort = os.Getenv("DB_SERVICE_PORT"); dbServicePort == "" {
		dbServicePort = "3031"
	}
	if dbServiceName = os.Getenv("DB_SERVICE_NAME"); dbServiceName == "" {
		dbServiceName = "serverdb"
	}
	request := "http://" + dbServiceName + ":" + dbServicePort + "/"
	response, err := http.Get(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		p := make([]byte, response.ContentLength)
		response.Body.Read(p)
		fmt.Fprintf(w, "Response : %s", string(p))
	}
}
