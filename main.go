package main

import (
	"encoding/csv"
	"log"
	"fmt"
	"net/http"
	"strings"
	"io"
	"strconv"
)

func main() {
	http.HandleFunc("/echo", EchoHandler)
	http.HandleFunc("/invert", InvertHandler)
	http.HandleFunc("/flatten", FlattenHandler)
	http.HandleFunc("/sum", SumHandler)
	http.HandleFunc("/multiply", MultiplyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func EchoHandler(w http.ResponseWriter, r *http.Request){
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return 
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var response string
	for {
		line, err := reader.Read()
		if err == io.EOF{
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for c := 0 ; c < len(line); c++ {
			if (line[c] != "") {
				response = fmt.Sprintf("%s%s,", response, line[c] )		
			}
		}
		response = strings.TrimRight(response,",") + "\n"
	}
	fmt.Fprint(w, response)
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request){
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return 
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var product int = 1
	for {
		line, err := reader.Read()
		if err == io.EOF{
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for c := 0 ; c < len(line); c++ {
			num, err := strconv.Atoi(line[c])
			if (err == nil) {
				product = product * num
			}
		}
	}
	fmt.Fprint(w, product)
}

func SumHandler(w http.ResponseWriter, r *http.Request){
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return 
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var sum int
	for {
		line, err := reader.Read()
		if err == io.EOF{
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for c := 0 ; c < len(line); c++ {
			num, err := strconv.Atoi(line[c])
			if (err == nil) {
				sum = sum + num
			}
		}
	}
	fmt.Fprint(w, sum)
}

func FlattenHandler(w http.ResponseWriter, r *http.Request){
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return 
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var response string
	for {
		line, err := reader.Read()
		if err == io.EOF{
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for c := 0 ; c < len(line); c++ {
			if (line[c] != "") {
				response = fmt.Sprintf("%s%s,", response, line[c] )		
			}
		}
	}
		fmt.Fprint(w, strings.TrimRight(response,","))
}

func InvertHandler(w http.ResponseWriter, r *http.Request){
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return 
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var response string
	var invertedMatrix [][]string
	var row int = 0
	for {
		line, err := reader.Read()
		if err == io.EOF{
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for c := 0 ; c < len(line); c++ {
			if (line[c] != "") {
				if (len(invertedMatrix) > c) {
					invertedMatrix[c] = append(invertedMatrix[c], line[c])
				} else {
					invertedMatrix = append(invertedMatrix, []string{line[c]})
				}
			}
		}
		row = row + 1
	}
	for r := range invertedMatrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(invertedMatrix[r], ","))
	}
	fmt.Fprint(w, response)
}
