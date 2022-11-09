package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var abc = 97
var printLimit = 9

// used for indexing and ranking words
type RankedDictionary []struct {
	word string
	rank int
}

// used for indexing and ranking letters
type LetterCount []struct {
	letter byte
	count  int
}

// reads cvs file with "word,index" format
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

// converts full
func loadFullMapAllWords(input [][]string) RankedDictionary {

	var rankedDict RankedDictionary
	for _, value := range input {
		intIndex, err := strconv.Atoi(value[0])
		if err != nil {
			log.Fatal("Unable to convert to int", err)
		}
		wordLength, err := strconv.Atoi(value[3])
		if err != nil {
			log.Fatal("Unable to convert to int", err)
		}

		//to fix an issue where random words are uppercase for some reason
		var lowerWord = strings.ToLower(value[1])
		rankedWord := struct {
			word string
			rank int
		}{lowerWord, intIndex}
		if wordLength == 5 {
			rankedDict = append(rankedDict, rankedWord)
		}

	}
	return rankedDict
}

// converts full
func loadFullMap(input [][]string) RankedDictionary {

	var rankedDict RankedDictionary
	for _, value := range input {
		intIndex, err := strconv.Atoi(value[1])
		if err != nil {
			log.Fatal("Unable to convert to int", err)
		}
		//to fix an issue where random words are uppercase for some reason
		var lowerWord = strings.ToLower(value[0])
		rankedWord := struct {
			word string
			rank int
		}{lowerWord, intIndex}
		rankedDict = append(rankedDict, rankedWord)

	}
	return rankedDict
}

func loadPosRanks(index int, rankedDict RankedDictionary) LetterCount {
	/*
		1. loop thru full map and isolate single letter
		2. create array of first letters
		3. count each instance of a letter
		4. then return a map with a sorted list of most used letter
	*/

	//list of all of the letters in the dictionary
	var letterArray []byte

	//loop thru full data set, adding each instance to the list
	fmt.Println("\nLoc2 - letterArray building:")
	for _, value := range rankedDict {
		fmt.Print(string(value.word[index]))
		//fmt.Println(value.word)
		//fmt.Println(value.word[index])
		//vType := fmt.Sprintf("%T", value.word[index])
		//fmt.Println(vType) // "[]int"

		letterArray = append(letterArray, value.word[index])
	}

	fmt.Println("\n\nLoc2.1 - letterArray:")
	fmt.Println(letterArray[0:printLimit])
	fmt.Println()

	//create a struct (since map isn't easily sortable)
	//with letter value (97-122) and count, then sort it
	//initialize struct
	var letterCount LetterCount
	var i byte
	//initially fill with zeros for everything
	for i = 0; i < 26; i++ {
		singleLetterCount := struct {
			letter byte
			count  int
		}{(i + byte(abc)), 0}
		letterCount = append(letterCount, singleLetterCount)
	}
	fmt.Println("Loc3 - letterCount (empty)")
	fmt.Println(letterCount)
	//add values for each letter
	for _, letterFromArray := range letterArray {
		//adding safety to ensure all letters are in a-z range
		if (letterFromArray < (byte(abc))) || (letterFromArray > (byte(26 + abc))) {
			fmt.Println("ERROR")
			fmt.Printf("\nletterFromArray: %v", letterFromArray)
			fmt.Printf("\nstring(letterFromArray): %v", string(letterFromArray))
			fmt.Println()

			os.Exit(1)
		}

		//debugging out of rang error
		//fmt.Println(letterFromArray)
		//fmt.Println(letterFromArray-byte(abc))
		letterCount[letterFromArray-byte(abc)].count++
	}
	fmt.Println("\nLoc4 - updated letterCount")
	fmt.Println(letterCount)

	sort.Slice(letterCount, func(i, j int) bool {
		return letterCount[i].count > letterCount[j].count
	})

	fmt.Println("\nLoc5 - sorted letterCount")
	fmt.Println(letterCount)

	//replace count with index
	//(this means that "count" is now a misnomer..
	//This could be fixed by making a new struct
	//with a correct name. Maybe for later)
	for i, _ := range letterCount {
		letterCount[i].count = i
	}

	fmt.Println("\nLoc6 - replace count with index")
	fmt.Println(letterCount)

	return letterCount
}

func getAllPosLetterCounts(rankedDict RankedDictionary) []LetterCount {
	//initialize return object
	var allLetterCount []LetterCount

	//loop through each of the 5 letters and add to allLetterCount
	for i := 0; i < 5; i++ {
		var posLetterCount = loadPosRanks(i, rankedDict)
		allLetterCount = append(allLetterCount, posLetterCount)
	}
	return allLetterCount
}

func getAllWordRanks(indexedDict RankedDictionary, allLetterRanks []LetterCount) RankedDictionary {
	//convert allLettersRanks ([]LetterCount) to []map[byte]int so indexing is easier
	fmt.Println("Loc7.1 - allLetterRanks")
	//iterate thru the 5 letterRanks array
	var allLetterMapArray []map[byte]int
	for _, allLetterValue := range allLetterRanks {
		var oneLetterMap = make(map[byte]int)
		//iterate thru the 26 letters of each
		for _, oneLetterValue := range allLetterValue {
			//fmt.Printf("\n[%d]%d - ",oneLetterValue.count,oneLetterValue.letter)
			oneLetterMap[oneLetterValue.letter] = oneLetterValue.count
		}
		fmt.Println(oneLetterMap)
		allLetterMapArray = append(allLetterMapArray, oneLetterMap)

	}

	fmt.Println("\nLoc7.2 - allLetterMapArray[0][96]")
	//fmt.Println(allLetterMapArray[0])
	fmt.Println(allLetterMapArray[0][96])
	fmt.Println()

	//iterate thru list of words
	fmt.Println("\n\nLoc8 - iterate thru indexedDict")

	//copy indexedDict - we will update the index with the distance, then sort and index again
	//var distanceDict = indexedDict
	for wordIndex, wordValue := range indexedDict {
		var wordDistance int
		if wordIndex <= printLimit {
			fmt.Printf("\n[%v] %v -", wordIndex, wordValue.word)
		}
		for letterIndex, letter := range wordValue.word {
			//fmt.Printf("[%v]%v-%v . ",letterIndex,string(letter), allLetterMapArray[letterIndex])
			if wordIndex <= printLimit {
				fmt.Printf("[%v]%v-%v . ", letterIndex, string(letter), allLetterMapArray[letterIndex][byte(letter)])
			}
			wordDistance = wordDistance + allLetterMapArray[letterIndex][byte(letter)]
		}
		if wordIndex <= printLimit {
			fmt.Printf("      wordDistance = %v", wordDistance)
		}
		//update index with distance
		indexedDict[wordIndex].rank = wordDistance
	}

	fmt.Println("\n\nLoc10 - new indexedDict with distance[0:printLimit]")
	fmt.Println(indexedDict[0:printLimit])

	//sort by distance
	sort.Slice(indexedDict, func(i, j int) bool {
		return indexedDict[i].rank < indexedDict[j].rank
	})

	//fmt.Println("\n\nLoc11 - new indexedDict sorted")
	//fmt.Println(indexedDict)

	return indexedDict
}

func test() {
	records := readCsvFile("data/top_100.csv")
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

	//finding the int value of letters
	//a=97
	//z=122
	//so letterIndex + 97 to get the int value
	fmt.Println("\n\nPosition 1:")
	fmt.Println(string(97))
	fmt.Println(string(122))
	fmt.Println(string(97 + 4)) //printing the 5th (index 4) letter (e)

	fmt.Println("Loc0 - testing byte to string")
	for i := 97; i < 123; i++ {
		fmt.Printf("[%v]%v - ", i, string(i))
	}
}

func getFirstWordRanks() RankedDictionary {

	//Read the csv file
	//records := readCsvFile("data/top_2000.csv")
	records := readCsvFile("data/all_words_update2.csv")

	//parse into map
	//var indexedDict = loadFullMap(records)
	var indexedDict = loadFullMapAllWords(records)
	fmt.Println("Loc1 - indexedDict:[0:printLimit]")
	fmt.Println(indexedDict[0:printLimit])
	fmt.Println()

	//get map for each letter position
	var allLetterRanks = getAllPosLetterCounts(indexedDict)
	fmt.Println("\nLoc7 - allPosLetterRanks")
	fmt.Println(allLetterRanks)

	//rank the words by sum distance
	var allWordRanks = getAllWordRanks(indexedDict, allLetterRanks)
	fmt.Println("\nLoc9 - allWordsRanks[0:printLimit]")
	fmt.Println(allWordRanks[0:printLimit])

	return allWordRanks

}

func main() {
	//test()
	
	var initialRankedDictionary = getFirstWordRanks()
	fmt.Println(initialRankedDictionary[0:printLimit])


}
