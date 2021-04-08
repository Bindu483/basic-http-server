package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hellow world")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hellow world"))
		return
	})
	http.HandleFunc("/bindu", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("uday corn"))
		return
	})
	http.HandleFunc("/greeting", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL.String())
		fmt.Println(request.Method)
		fmt.Println(request.URL.Query().Get("msg"))
		writer.Write([]byte(fmt.Sprintf(`
<!DOCTYPE html>
<html>
    <head>
        <!-- head definitions go here -->
    </head>
    <body>
        <h1> Hello your message was %s </h1>
    </body>
</html>
`, request.URL.Query().Get("msg"))))
		return
	})
	err := http.ListenAndServe("0.0.0.0:8000", nil)

	if err != nil {
		fmt.Println("couldnt start server", err)
	}
}
