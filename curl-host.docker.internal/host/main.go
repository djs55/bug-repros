package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

const addr = "127.0.0.1:8080"

func main() {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on " + addr)
	h := &handler{}
	go func() {
		_ = http.Serve(l, h)
	}()
	log.Println("Hit enter to stop.")
	b := make([]byte, 1)
	_, _ = os.Stdin.Read(b)
	_ = l.Close()
	log.Printf("Handled %d responses\n", h.n)
}

type handler struct {
	n int
}

func (h *handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello"))
	if err != nil {
		log.Printf("error writing response: %v", err)
	}
	h.n++
	if h.n%100 == 0 {
		log.Printf("Handled %d responses\n", h.n)
	}
}
