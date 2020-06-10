package main

import (
	_ "fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type EntryPoint struct {
	Endpoint string `yaml:"endpoint"` // ServerMux pattern
	Status int `yaml:"status,omitempty"` // Http status
	Size string `yaml:"size,omitempty"` // size in bytes. Possible define a range N..M - the random value in this range will be chosen
	TimeMs string `yaml:"time,omitempty"` // size in milliseconds. Possible define a range N..M - the random value in this range will be chosen
}

type conf struct {
	Endpoints []EntryPoint `yaml:"endpoints"`
}

func getConf() (*conf, error) {
	c := &conf{}
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

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

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	c, err := getConf()
	if err != nil {
		log.Fatal("Error ", err)
	}

	http.HandleFunc("/", hello)
	for _, entry := range c.Endpoints {
		if entry.Endpoint != "" && entry.Endpoint != "\\" {
			http.HandleFunc(entry.Endpoint, handler(entry))
		}
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
