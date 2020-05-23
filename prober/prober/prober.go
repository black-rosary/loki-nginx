package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Link struct {
	link string
	probabilty int
}

var links = []Link{
	{"http://nginx.test/200", 200},
	{"http://nginx.test/500", 2},
	{"http://nginx.test/400", 1},
	{"http://nginx.test/404", 5},
	{"http://nginx.test/502", 2},

	{"http://nginx.test/css/main.css", 100},
	{"http://nginx.test/js/logic.js", 75},
	{"http://nginx.test/img/logo.png", 50},
	{"http://nginx.test/img/not_found.jpg", 2},

	{"http://nginx.test/photo/a1", 100},
	{"http://nginx.test/photo/a2", 100},
	{"http://nginx.test/photo/a3", 2},

	{"http://nginx.test/index", 50},
	{"http://nginx.test/registration", 20},
	{"http://nginx.test/login", 50},
	{"http://nginx.test/very_slow", 5},
}


func selectLinkByProbabilty() Link {
	var sum, prob, trshld = 0, 0, 0
	for i := 0; i < len(links); i++ {
		sum += links[i].probabilty
	}
	trshld = rand.Intn(sum)
	for i := 0; i < len(links); i++ {
		prob += links[i].probabilty
		if prob > trshld {
			return links[i]
		}
	}
	return links[0]
}

func requestSomething(url string) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Stub_Bot/3.0")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("FAILED %s  \n", url)
	} else {
		fmt.Printf("OK %s \n", url)
	}
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	for {
		for i := 0; i < rand.Intn(20); i++ {
			s := selectLinkByProbabilty()
			go requestSomething(s.link)
		}
		time.Sleep(time.Duration(763) * time.Millisecond)
	}
}