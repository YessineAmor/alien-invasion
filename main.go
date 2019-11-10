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
	// Check command line arguments
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run alien-invasion.go <filePath> <numberOfAliens>")
	}
	filePath := os.Args[1]
	numberOfAliens, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Simulating alien invasion with cities from", filePath, "and with", numberOfAliens, "aliens")
	// Open and read lines from input file and populate cityNames and cities
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		city := scanner.Text()
		// Seperate words from current line in an array words
		words := strings.Fields(city)
		// Get city and append it to cityNames array
		currentCity := words[0]
		cityNames = append(cityNames, currentCity)
		// iterate through each destination and append it to the list of paths out of the current city
		for _, element := range words[1:] {
			// Get rid of the direction(north,south,west,east) and append only the city name
			cities[currentCity] = append(cities[currentCity], strings.Split(element, "=")[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created cities from input file:", cities)
	fmt.Println("Initializing aliens locations..")
	// Place aliens randomly on the map
	for i := 1; i <= numberOfAliens; i++ {
		rand.Seed(time.Now().UnixNano())
		aliens[i] = cityNames[rand.Intn(len(cities))]
		fmt.Println("Alien", i, "is in", aliens[i], "city")
	}

}
