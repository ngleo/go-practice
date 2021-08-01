package main

import (
	"fmt"
	"log"
	"net/http"
)

type CounterStore struct {
	counts map[string]int
}

type Command struct {
	name      string
	replyChan chan int
}

func startBackgroundManager() chan Command {
	counts := map[string]int{"channel1": 0, "channel2": 0}
	cmds := make(chan Command)

	go func() {
		for cmd := range cmds {
			if _, ok := counts[cmd.name]; ok {
				counts[cmd.name]++
				fmt.Println(counts)
				cmd.replyChan <- counts[cmd.name]
			} else {
				cmd.replyChan <- -1
			}
		}
	}()
	return cmds
}

type Server struct {
	cmds chan<- Command
}

func main() {
	s := Server{cmds: startBackgroundManager()}
	http.HandleFunc("/inc", s.handleInc)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func (s *Server) handleInc(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	replyChan := make(chan int)

	s.cmds <- Command{name: name, replyChan: replyChan}
	reply := <-replyChan
	fmt.Println(reply)
	if reply >= 0 {
		fmt.Fprintf(w, "ok")
	} else {
		fmt.Fprintf(w, "can't find channel")
	}
}
