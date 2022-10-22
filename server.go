package main 

import (
	"log"
	"net/http"
	"fmt"
)


func main(){
	listen()
}

//const for the hostname and port 
const (
	HOST = "localhost"
	PORT = "8090"
)

// creating our pages 
func pageOne(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "This is page one\n")
}

func pageTwo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is page two\n")
}

func pageThree(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is page three\n")
}

func pageFour(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w,"This is page four\n")
}

// handling pages url with go default http mux
func handler(mux *http.ServeMux) {
	mux.HandleFunc("/one", pageOne)
	mux.HandleFunc("/two", pageTwo)
	mux.HandleFunc("/three", pageThree)
	mux.HandleFunc("/four", pageFour)
}

// serving the server on the provided port and host 
func listen(){

	mux := http.NewServeMux()
	handler(mux)

	err := http.ListenAndServe(HOST+":"+PORT, mux)
	if err != nil {
		log.Fatal(err)
	}
}


