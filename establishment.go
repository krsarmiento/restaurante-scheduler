package main

import (
	"math/rand"
	"time"
	"strconv"
	"sync"
)

type Establishment struct {
	id int
	isOpen bool
	opensIn string
}

func GetEstablishments(batch int) []Establishment {
	var establishments []Establishment
	starts := (batch-1) * batchSize
	ends := batchSize * batch
	var minutes int
	var opensIn string
	for i := starts; i < ends; i++ {
		minutes = rand.Intn(30)
		opensIn = strconv.Itoa(minutes) + " minutes."
		restaurant := Establishment{
			i,
			false,
			opensIn,
		}
		establishments = append(establishments, restaurant)
	}
	return establishments
}

func EvaluateEstablishment(establishment Establishment, logger chan string, wg *sync.WaitGroup) bool {
	defer wg.Done()
	logger <- ">> Evaluating establishment #" + strconv.Itoa(establishment.id)
	isOpen := isEstablishmentOpen(establishment.id, "establecimiento")
	establishment.isOpen = isOpen
	saveInDatabase(establishment)
	logger <- "<< Evaluation finished #" + strconv.Itoa(establishment.id)
	return true
}

func isEstablishmentOpen(id int, scheduleType string) bool {
	//TODO: EXECUTE SCHEDULE'S QUERY
	//IF WE GET RECORDS WE RETURN TRUE
	//FALSE OTHERWISE
	seconds := rand.Intn(10)
	time.Sleep(time.Duration(seconds) * time.Second)
	return true
}

func saveInDatabase(restaurant Establishment) bool {
	//TODO: Save restaurant in database
	time.Sleep(time.Duration(1) * time.Second)
	return true
}