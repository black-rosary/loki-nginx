package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func randRange(start int, end int) int {
	if start >= end {
		return start
	}
	return start + rand.Intn(end - start)
}

func parseRange(r string) int {
	var l, start, end = 0,0,0
	if r == "" {
		return 0
	}
	s := strings.Split(r, "..")
	if len(s) > 0 {
		start, _ = strconv.Atoi(s[0])
		if len(s) > 1 {
			end, _ = strconv.Atoi(s[1])
		}
		l = randRange(start, end)
	}
	return l
}

func delay(ms string) {
	delayMs := parseRange(ms)
	if delayMs == 0 {
		return
	}
	time.Sleep(time.Duration(delayMs) * time.Millisecond)
}

func getStatus(status int) int {
	if status == 0 {
		return http.StatusOK
	}
	if status < 100 || status > 599 {
		log.Printf("status %d not supported. Ignored", status)
		return http.StatusOK
	}
	return status
}

func getBytes(size string) []byte {
	l := parseRange(size)
	ret := make([]byte, l)
	for i := range ret {
		ret[i] = 'x'
	}
	return ret
}

func handler(e EntryPoint) func(w http.ResponseWriter, req *http.Request)  {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(getStatus(e.Status))
		w.Write(getBytes(e.Size))
		delay(e.TimeMs)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found"))
}

func RunBackyard(port string, entries []EntryPoint) {
	http.HandleFunc("/", hello)
	for _, entry := range entries {
		if entry.Endpoint != "" && entry.Endpoint != "\\" {
			http.HandleFunc(entry.Endpoint, handler(entry))
		}
	}
	go func() {
		log.Fatal(http.ListenAndServe(":" + port, nil))
	}()
}