package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Print("handler started")
	defer log.Print("handler ended")

	// async work happens here and sends completion to a channel
	// can use close(completion) to throw an error
	completion := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		log.Print("task done")
		completion <- "Here is the result"
	}()

	select {
	// case <-time.After(time.Second * 5):
	// 	fmt.Fprintln(w, "hello")
	case msg, _ := <-completion:
		fmt.Fprintln(w, msg)
	case <-ctx.Done():
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
