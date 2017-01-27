package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
	"encoding/json"
    "crypto/tls"
    "golang.org/x/crypto/acme/autocert"
)

func main() {
    fmt.Println("hello")
    listenServer()
}

func listenServer() {

    certManager := autocert.Manager{
        Prompt: autocert.AcceptTOS,
        HostPolicy:autocert.HostWhitelist("air.coretime.cn"),
        Cache: autocert.DirCache("/etc/letsencrypt/csr/"),
    }

	r := mux.NewRouter()
	r.HandleFunc("/", auth).Methods("GET")
	http.Handle("/", r)
    server := &http.Server{
        Addr: ":443",
        TLSConfig:&tls.Config{
            GetCertificate:certManager.GetCertificate,
        },
    }
    server.ListenAndServeTLS("","")
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