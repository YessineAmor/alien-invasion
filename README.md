# alien-invasion
[![Go Report Card](https://goreportcard.com/badge/github.com/YessineAmor/alien-invasion)](https://goreportcard.com/report/github.com/YessineAmor/alien-invasion)

Alien invasion simulation in Go.

Aliens are about to invade the earth so me being the expert, I've been tasked with creating a simulation for the invasion.

The program takes a file containing the names of cities followed by 1-4 directions (north, south, west, east). Each one represents a road to another city that lies in that direction.

A number of aliens will also need to be provided.

## Usage:
```console
$ go run main.go <filePath> <numberOfAliens> 
```
### Example:
```console
$ go run main.go test/input.txt 4
```
