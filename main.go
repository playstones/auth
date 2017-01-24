package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
	"encoding/json"
)

func main() {
    fmt.Println("hello")
    listenServer()
}

func listenServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", auth).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func auth(w http.ResponseWriter, r *http.Request)  {
    r.ParseForm()
    fmt.Println(r.Header)
    w.Header().Set("Content-Type", "application/json")
    result := map[string]bool{
        "result": true,
    }

    res, err := json.Marshal(result)
    if err != nil {
        http.Error(w,"Something wrong", 404)
        return
    }
    w.Write(res)
}