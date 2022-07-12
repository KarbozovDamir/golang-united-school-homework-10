package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", getNameParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", getBadUrl).Methods(http.MethodGet)
	router.HandleFunc("/data", postData).Methods(http.MethodPost)
	router.HandleFunc("/headers", postHeaders).Methods(http.MethodPost)
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func getNameParam(w http.ResponseWriter, r *http.Request) {
	bodyReq := mux.Vars(r)["PARAM"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello, %s!", bodyReq)))

}

func getBadUrl(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func postData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("I got message:\n%s", body)))
}

func postHeaders(w http.ResponseWriter, r *http.Request) {
	a, err1 := strconv.Atoi(r.Header.Get("a"))
	b, err2 := strconv.Atoi(r.Header.Get("b"))
	if err1 != nil && err2 != nil {
		return
	}
	w.Header().Set("a+b", strconv.Itoa(a+b))
	w.WriteHeader(http.StatusOK)
}
