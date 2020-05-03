package proxy

import (
	"GO-CRUD/common"
	"fmt"
	"net/http"
	"time"
)

func Proxy(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	result1 := make(chan interface{})
	result2 := make(chan interface{})

	changelogs := Changelogs{}

	go common.MakeRequest("https://manychat.com/changelog/get", &changelogs, result1)
	go common.MakeRequest("https://jsonplaceholder.typicode.com/todos/1", &changelogs, result2)

	select {
	case msg1 := <-result1:
		fmt.Println("received 1")
		common.SendJSONresponse(msg1, w)

	case msg2 := <-result2:
		common.SendJSONresponse(msg2, w)
		fmt.Println("received 2")
	}

	fmt.Println(time.Since(start).Seconds(), "seconds")
}
