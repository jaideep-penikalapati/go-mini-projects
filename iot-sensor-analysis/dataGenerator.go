package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func generate(channel chan bool) {
	fmt.Println("Creating file: ")
	go createFile(channel)
	fmt.Println("File created: ")
}

func createFile(channel chan bool) {
	rand.Seed(time.Now().UnixNano())

	var msg string
	var counter int
	var temp float32

	f, err := os.Create("data.json")
	if err != nil {
		log.Fatal("Not able to create file")
	}

	defer f.Close()

	w := bufio.NewWriter(f)
	var t = time.Now()
	minTemp := 20
	msg = fmt.Sprintf("{\n\"name\": \"IOTSensor1\",\n\"tempMin\":20,\n\"tempMax\":30,\n\"interval\":60,\n\"values\": [\n")
	w.WriteString(msg)

	for i := 1; i <= 1200; i++ {
		temp = float32(minTemp) + float32(rand.Intn(20))*rand.Float32()
		msg = fmt.Sprintf("{\"messageId\":%v,\"temperature\":%v,\"enqueuedTime\":\"%v\"},\n", counter, temp, t.Format("2006-01-02 15:04:05"))
		t = t.Add(time.Duration(20+rand.Intn(20)) * time.Second)
		if i == 1200 {
			msg = msg[:len(msg)-2]
		}

		w.WriteString(msg)
		counter++
	}
	w.WriteString("\n\t]\n}")

	w.Flush()
	channel <- true
}
