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
	fmt.Println("Starting invasion..")
	numberOfDeadAlien := 0
	numberOfIterations := 1
	for numberOfIterations <= 10000 {
		fmt.Println("Day", numberOfIterations, "of invasion")
		for i := 1; i <= numberOfAliens; i++ {
			if aliens[i] != "DEAD" {
				aliensCurrentCity := aliens[i]
				possibleDestinations, exists := cities[aliensCurrentCity]
				// If there's no entry for the city in the map or if there's no possible destination out of the city, then alien is considered contained
				if !exists || len(possibleDestinations) == 0 {
					fmt.Println("Alien", i, "has died in", aliensCurrentCity)
					numberOfDeadAlien++
					aliens[i] = "DEAD"
				} else {
					rand.Seed(time.Now().UnixNano())
					// Choose a random destination index from the alien's current city.
					destinationIndex := rand.Intn(len(possibleDestinations))
					newDestination := cities[aliensCurrentCity][destinationIndex]
					// Move alien to new city
					aliens[i] = newDestination
					fmt.Println("Alien number", i, "moved from", aliensCurrentCity, "to", aliens[i])
					// Now check if there's an alien in that city
					for k, v := range aliens {
						if strings.EqualFold(v, newDestination) && k != i {
							fmt.Println("Alien fight! Alien ", k, "is fighting with alien", i, "in", newDestination, "city... The fight ends with the city destroyed and all aliens in this city are now dead.")
							// Kill both aliens
							aliens[i] = "DEAD"
							aliens[k] = "DEAD"
							numberOfDeadAlien = numberOfDeadAlien + 2
							// Detroy city - no way out
							cities[newDestination] = []string{}
							// Remove city as potential destination from other cities
							for city, paths := range cities {
								for index, path := range paths {
									if path == newDestination {
										cities[city][index] = cities[city][len(cities[city])-1]
										cities[city] = cities[city][:len(cities[city])-1]
									}
								}
							}
							break

						}
					}
				}
			}
			if numberOfDeadAlien == numberOfAliens {
				fmt.Println("Congrats! All the aliens are either dead or contained! It's a victory for humanity!... I think")
				// Print whats left of the world
				fmt.Println("Remaining cities and paths are:")
				for city, paths := range cities {
					if len(paths) > 0 {
						fmt.Print("from ", city, " you can go to ")
						for _, path := range paths {
							fmt.Print(path, " ")
						}
						fmt.Println()
					}
				}
				os.Exit(0)
			}
		}
		numberOfIterations++
	}
}
