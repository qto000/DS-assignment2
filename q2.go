package main

import (
	"bufio"
	"io"
	"strconv"
	//"fmt"
	"os"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature. 
func sumWorker(nums chan int, out chan int) {
	// TODO: implement me
	// HINT: use for loop over `nums`
	result := 0
	for num := range nums {
		result += num
	}
	out <- result
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	// TODO: implement me
	// HINT: use `readInts` and `sumWorkers`
	// HINT: used buffered channels for splitting numbers between workers
	f, err := os.Open(fileName)
	checkError(err)
	values, err := readInts(f)
	checkError(err)
	outChannel := make(chan int)
	bufSize := len(values) / num
		
	for w := 0; w < num; w++ {
		bufChannel := make(chan int, bufSize)
		sumChannel := 0
		for j := (w * bufSize); j < ((w + 1) * bufSize); j++{
			sumChannel += values[j] 
		}
		
		bufChannel <- sumChannel
		close(bufChannel)
		go sumWorker(bufChannel, outChannel)
	}
	
	result := 0
	for i := 0; i < num; i++ {
		result += <- outChannel
	}

	return result

}
// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}

// func main() {
// 	answers := sum(5, "q2_test1.txt")
// 	fmt.Println(answers)
// }