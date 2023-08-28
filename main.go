package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World!")
	case "/test":
		fmt.Fprint(w, "Test Success!")
	default:
		fmt.Fprint(w, "Error!")
	}
	fmt.Printf("This is a %s request!\n", r.Method)
}

func contentType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "<h1>Hello World!</h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Attempting Timeout...")
	// time.Sleep(1 * time.Second)
	fmt.Fprint(w, "<h1>Timeout Unsuccessful!</h1>")
}

func helloWorldDarkMode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Activating Dark Mode...")
	fmt.Fprint(w, "<h1 style=\"background-color:grey;\">Hello World!</h1>")
}

func main() {
	// http.HandleFunc("/", helloWorldPage)
	http.HandleFunc("/content", contentType)
	http.HandleFunc("/timeout", timeout)
	// http.ListenAndServe("", nil)

	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}

	var muxDarkMode http.ServeMux
	server.Handler = &muxDarkMode
	muxDarkMode.HandleFunc("/", helloWorldDarkMode)

	server.ListenAndServe()
}
