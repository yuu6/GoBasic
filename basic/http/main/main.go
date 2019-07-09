package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloworld)
	http.ListenAndServe(":8000", nil)
}

func helloworld(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}
	switch request.Method{
	case "GET":
		for k, v := range request.URL.Query(){
			fmt.Printf("%s:%s\n",k,v)
		}
		writer.Write([]byte("Receive a GET request \n"))
	case "POST":
		reqBody, err := ioutil.ReadAll(request.Body)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Printf("%s\n",reqBody)
		writer.Write([]byte("Received a POST request\n"))
	default:
		writer.WriteHeader(http.StatusNotImplemented)
		writer.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}
