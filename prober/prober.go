package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type EntryPoint struct {
	Endpoint string `yaml:"endpoint"` // ServerMux pattern
	Status int `yaml:"status,omitempty"` // Http status
	Size string `yaml:"size,omitempty"` // size in bytes. Possible define a range N..M - the random value in this range will be chosen
	TimeMs string `yaml:"time,omitempty"` // size in milliseconds. Possible define a range N..M - the random value in this range will be chosen
	Prob int `yaml:"weight"` // probability of request
}

type conf struct {
	Endpoints []EntryPoint `yaml:"endpoints"`
}

func (e EntryPoint) getUrl(host string) string {
	return host + e.Endpoint
}

func selectLinkByProbability(links []EntryPoint) EntryPoint {
	var sum, prob, trshld = 0, 0, 0
	for i := 0; i < len(links); i++ {
		sum += links[i].Prob
	}
	trshld = rand.Intn(sum)
	for i := 0; i < len(links); i++ {
		prob += links[i].Prob
		if prob > trshld {
			return links[i]
		}
	}
	return links[0]
}

func requestSomething(url string) {
	time.Sleep(time.Duration(500 + rand.Intn(500)) * time.Millisecond)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Stub_Bot/0.1")
	resp, err := client.Do(req)

	if err == nil {
		resp.Body.Close()
	}
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

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	host := os.Getenv("VIRTUAL_HOST")
	if host == "" {
		host = "localhost"
		log.Println("$VIRTUAL_HOST is not defined. Use localhost as default ")
	}

	port := os.Getenv("VIRTUAL_PORT")
	if port == "" {
		port = "8082"
		log.Println("Error: $VIRTUAL_PORT is not defined. Use 8082 as default")
	}

	c, err := getConf()
	if err != nil {
		log.Fatal("Error get config ", err)
	}

	RunBackyard(port, c.Endpoints)

	maxGoroutines := 16
	guard := make(chan struct{}, maxGoroutines)

	for {
		guard <- struct{}{}
		s := selectLinkByProbability(c.Endpoints)
		go func(url string) {
			requestSomething(url)
			<-guard
		}(s.getUrl("http://" + host))
	}
}
