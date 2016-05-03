package main

import "fmt"

func main() {
	n := hashBucket("Go", 12)
	fmt.Println(n)
}

func hashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
/*
    hashBucket returns a hash bucket wrt the value of words first character
    this is a hashing function to produce a hash bucket to distribute hashed
    data. The hash function needs to be optimized wrt table so that the data 
    is evenly distributed across the buckets.  Both the data and number of 
    buckets cnotribute to the optimization. 
*/