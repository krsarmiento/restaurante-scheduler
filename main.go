package main

import (
	"fmt"
	"os"
	"log"
	"sync"
	"strconv"
)

var batchSize int = 25
var logFile string = "output.log"

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Batch argument not provided.")
		return
	}
	batch, _ := strconv.Atoi(args[0])

	var wg sync.WaitGroup
	logger := make(chan string)
	establishments := GetEstablishments(batch)

	for _, establishment := range establishments {
		wg.Add(1)
		go EvaluateEstablishment(establishment, logger, &wg)
	}

	go closeChannel(logger, &wg)
	logResults(logger)
}

func logResults(logger chan string) {
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	for message := range logger {
		log.Println(message)
	}
}

func closeChannel(logger chan string, wg *sync.WaitGroup) {
	wg.Wait()
	close(logger)
}
