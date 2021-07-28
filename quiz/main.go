package main

import "flag"

var csvFilePath string

func init() {
	flag.StringVar(&csvFilePath, "csv", "problems.csv", "Absolute path to the csv-file that contains a quiz (in the format <problem>,<solution>)")
}

func main() {
	flag.Parse()
}
