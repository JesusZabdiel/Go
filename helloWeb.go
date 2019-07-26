package main

import(
	"net/http"
	"io"
	"fmt"
)

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/HelloWorld", func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w,"Hellooooooo mate")
	})
	http.ListenAndServe(":8000", nil)
}

func handler (w http.ResponseWriter, r *http.Request){
	
	fmt.Println("New petition")
	io.WriteString(w,"Hello World")
}