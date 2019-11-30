package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	//scan the pages
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	//Set the split function for the scanning operations.
	scanner.Split(bufio.ScanWords)

	//Create a slice of slice of string to hold slice of buckets
	buckets := make(map[int]map[string]int)
	for i := 0; i < 12; i++ {
		buckets[i] = make(map[string]int)
	}

	//loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := HashTable(word, 12)
		buckets[n][word]++
	}

	//Print len of each bucket
	for k, v := range buckets[6] {
		fmt.Println(v, "\t - ", k)
	}

	//Print the word in one of the buckets
	fmt.Println(buckets[6])

	//Print len of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}
}

func HashTable(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
