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

	//Create slice to hold counts
	buckets := make([]int, 12)

	//Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
		//fmt.Println(scanner.Text())
	}
	fmt.Println(buckets)
}

func HashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets

	return bucket
}
