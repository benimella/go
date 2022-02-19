package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("xxxxxxxxxxxx"))
	})
	http.ListenAndServe(":8080", nil)
}
