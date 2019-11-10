package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var cities = make(map[string][]string)
var cityNames = []string{}
var aliens = make(map[int]string)

func main() {

	if len(os.Args) != 3 {
		log.Fatal("Usage: go run alien-invasion.go <filePath> <numberOfAliens>")
	}
	filePath := os.Args[1]
	numberOfAliens, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Simulating alien invasion with cities from", filePath, "and with", numberOfAliens, "aliens")
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		city := scanner.Text()
		words := strings.Fields(city)
		currentCity := words[0]
		cityNames = append(cityNames, currentCity)
		for _, element := range words[1:] {
			cities[currentCity] = append(cities[currentCity], strings.Split(element, "=")[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created cities from input file:", cities)
	fmt.Println("Initializing aliens locations..")
	for i := 1; i <= numberOfAliens; i++ {
		rand.Seed(time.Now().UnixNano())
		aliens[i] = cityNames[rand.Intn(len(cities))]
		fmt.Println("Alien", i, "is in", aliens[i], "city")
	}

}
