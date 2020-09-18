package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"
)

// Values Array of sensor data
type Values struct {
	Name     string  `json:"name"`
	MinTemp  float32 `json:"minTemp"`
	MaxTemp  float32 `json:"maxTemp"`
	Interval int     `json:"interval"`
	Values   []Value `json:"values"`
}

// Value struct with sensor values
type Value struct {
	Message      string  `json:"message"`
	Temperature  float32 `json:"temperature"`
	EnqueuedTime string  `json:"enqueuedTime"`
}

// Reading struct
type Reading struct {
	hour       int
	normal     float32
	outOfRange float32
}

func main() {

	// Step: Generate Data
	channel := make(chan bool)
	go generate(channel)
	<-channel

	// Step: Read data from file
	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Fatal("File not found")
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Error parsing data")
	}

	var v Values
	json.Unmarshal(byteValue, &v)

	tempMap := make(map[int][]float32)
	fmt.Println("Total sensor reading available: ", len(v.Values))

	for _, val := range v.Values {
		t, err := time.Parse("2006-01-02 15:04:05", val.EnqueuedTime)
		if err != nil {
			log.Fatal(err)
		}
		h := t.Hour()
		tempMap[h] = append(tempMap[h], val.Temperature)
	}

	var normal, outOfRange float32
	var readings []Reading

	for x, y := range tempMap {
		normal, outOfRange = 0.0, 0.0
		for _, temp := range y {
			if temp >= v.MinTemp && temp <= v.MaxTemp {
				normal++
			} else {
				outOfRange++
			}
		}
		read := Reading{x, normal, outOfRange}
		readings = append(readings, read)
	}

	sort.Slice(readings, func(i, j int) bool {
		return readings[i].hour < readings[j].hour
	})

	printTable(readings)
	// printChart(readings)
}

func printTable(r []Reading) {
	fmt.Printf("Hour\tTotal\tNormal\tOut of Range\tPercent\n")
	for _, val := range r {
		total := val.normal + val.outOfRange
		percent := val.outOfRange / total * 100
		fmt.Printf("%v\t%v\t%v\t%5v\t\t%5.1f\n", val.hour, total, val.normal, val.outOfRange, percent)
	}
}
