package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
)

func SendJSONresponse(response interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func MakeRequest(url string, respInterface interface{}, channel chan interface{}) {
	resp, fetchErr := http.Get(url)

	if fetchErr != nil {
		log.Fatal("failed to fetch")
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		log.Fatal("failed to read response")
	}

	jsonResult := respInterface

	jsonErr := json.Unmarshal(body, &jsonResult)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	channel <- jsonResult

}

func checkLink(link string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get(link)
	if err != nil {
		c <- link + " is down"
		return
	}
	c <- link + " is up!"
}

func concurrencyTest() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)
	var wg sync.WaitGroup

	for _, link := range links {
		wg.Add(1)
		go checkLink(link, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for msg := range c {
		fmt.Println(msg)
	}

	fmt.Println(runtime.NumCPU())

}
