package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("timestamp:", time.Now().String())
		w.WriteHeader(http.StatusOK)
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
		// using dynamic TLS key pair
		// TLSConfig: &tls.Config{
		// 	GetCertificate: func(chi *tls.ClientHelloInfo) (*tls.Certificate, error) {
		// 		cert, err := tls.LoadX509KeyPair("tls.crt", "tls.key")
		// 		if err != nil {
		// 			return nil, err
		// 		}
		// 		return &cert, nil
		// 	},
		// },
	}

	// using static TLS key pair
	if err := server.ListenAndServeTLS("tls.crt", "tls.key"); err != nil {
		panic(err)
	}

	// using dynamic TLS key pair
	// if err := server.ListenAndServeTLS("", ""); err != nil {
	// 	panic(err)
	// }
}
