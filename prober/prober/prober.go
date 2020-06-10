package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Link struct {
	Url  string `yaml:"url"`
	Prob int `yaml:"weight"`
}

type conf struct {
	Urls []Link `yaml:"urls"`
}

func selectLinkByProbability(links []Link) Link {
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

	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("FAILED %s  \n", url)
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

	c, err := getConf()
	if err != nil {
		log.Fatal("Error ", err)
	}

	maxGoroutines := 5
	guard := make(chan struct{}, maxGoroutines)

	for {
		guard <- struct{}{}
		s := selectLinkByProbability(c.Urls)
		go func(url string) {
			requestSomething(url)
			<-guard
		}(s.Url)
	}
}
