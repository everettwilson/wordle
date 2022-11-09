package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Dictionary struct {
	//6-D slice
	//0: index
	//1: first letter
	//2: second letter ...
	dict [][][][][][]int
}

func loadDictionary(input [][]string) (returnDict Dictionary) {
	//looping over range: index (not used) and value ([word][index]string)
	localDict := make([][][][][][]int, 0)
	for _, value := range input {
		intIndex, err := strconv.Atoi(value[1])
		if err != nil {
			log.Fatal("Unable to parse file as CSV for "+filePath, err)
		}

		//in progress here
		//localDict = append(localDict, 
	}
	return
}

func loadMap (input [][]string) (map[int]string){
	
	dictMap = make(map[int]string)
	for _, value := range input {
		intIndex, err := strconv.Atoi(value[1])
		if err != nil {
			log.Fatal("Unable to parse file as CSV for "+filePath, err)
		}
		dictMap[intIndex] = value[0]

	}
	return dictMap
}

type DictionaryOld struct {
	WordList   []Word
	DictLength int
}

type Word struct {
	Characters string
	Rank       int
	//Char1 int
	//Char2 int
	//Char3 int
	//Char4 int
	//Char5 int
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func convertToDictionary(fullFile [][]string) (dict Dictionary) {
	//loop through each line of the fullFile

	for _, v := range fullFile {
		//convert string to int for index
		intIndex, err := strconv.Atoi(v[1])
		if err != nil {
			// ... handle error
			panic(err)
		}
		//create a new Word object with each line
		word := Word{Characters: v[0], Rank: intIndex}
		fmt.Printf("\nword: %+v", word)

		//add new word to word list
	}
	return
}

func test() {
	records := readCsvFile("data/top_10.csv")
	recordsType := fmt.Sprintf("%T", records)
	fmt.Println(recordsType) // "[]int"

	fmt.Println("\nPrint all:")
	fmt.Println(records)

	fmt.Println("\nPrint line:")
	fmt.Println(records[0])

	fmt.Println("\nPrint individual component of the line")
	fmt.Println(records[0][0])

	fmt.Println("\nPrint individual character of the line")
	fmt.Println(records[0][0][1])
	fmt.Println(string(records[0][0][1]))

	fmt.Printf("\nrecords length: %d", int(len(records)))
	fmt.Printf("\nrecords cap: %d", int(cap(records)))

}

func main() {
	//test()

	//Read the csv file
	records := readCsvFile("data/top_10.csv")

	//parse into map
	dictMap = loadMap(records)
	fmt.Println(dictMap)
}
