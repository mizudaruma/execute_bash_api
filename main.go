package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func hoge(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)

	if err != nil {
		fmt.Println("error!!", err)
		return
	}

	fmt.Println(buf.String())

}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	http.HandleFunc("/test", hoge)

	server.ListenAndServe()
}
