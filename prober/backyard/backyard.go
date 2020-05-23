package main

import (
	_ "fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func status(status int) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(status)
	}
}

func randomDelay(delay_ms int) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(time.Duration(rand.Intn(delay_ms)) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}
}

func randomSize(size int) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		b := make([]byte, rand.Intn(size))
		for i := range b {
			b[i] = 'm'
		}
		w.Write(b)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Ready"))
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/200", status(http.StatusOK))
	http.HandleFunc("/301", status(http.StatusMovedPermanently))
	http.HandleFunc("/400", status(http.StatusBadRequest))
	http.HandleFunc("/404", status(http.StatusNotFound))
	http.HandleFunc("/500", status(http.StatusInternalServerError))
	http.HandleFunc("/502", status(http.StatusBadGateway))
	http.HandleFunc("/css/main.css", randomSize(500))
	http.HandleFunc("/js/logic.js", randomSize(700))
	http.HandleFunc("/img/logo.png", randomSize(15000))
	http.HandleFunc("/img/not_found.jpg", randomSize(500))
	http.HandleFunc("/photo/a1", randomSize(10000))
	http.HandleFunc("/photo/a2", randomSize(10000))
	http.HandleFunc("/photo/a3", status(http.StatusNotFound))
	http.HandleFunc("/index", randomDelay(200))
	http.HandleFunc("/registration", randomDelay(500))
	http.HandleFunc("/login", randomDelay(500))
	http.HandleFunc("/very_slow", randomDelay(1000))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
