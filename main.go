package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "/Users/brendanwilde/Documents/code/go/go_1billionLineChallenge/measurements.txt"
	
	//the last buffer's end position
	//first time it's the size of the buffer
	//after that it's n returned from last read
	n := int64(8) 



	bufferExtraBytes, err := getBufferExtraBytes(filePath, n)
	if err != nil {
		fmt.Printf("Error seeking to last buffer end pos: %v, %v", n, err)
		os.Exit(1)
	}
	fmt.Println("bufferExtraBytes: ", bufferExtraBytes)

}

func getBufferExtraBytes(filePath string, bufferEnd int64) (extraBytes int64, err error){
	
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a buffer to hold a single byte
	singleByteBuf := make([]byte, 1)
	extraBytes = int64(0)

	for {
		// Seek to the current position in the file
		_, err := file.Seek(bufferEnd+extraBytes, 0)
		_ = err
		if err != nil {
			return 0, fmt.Errorf("error: getBufferExtraBytes: 1. seeking newline: %v", err)
		}

		// Read a single byte into the buffer
		_, err = file.Read(singleByteBuf)
		_ = err
		if err != nil {
			return 0, fmt.Errorf("error: getBufferExtraBytes: 2. reading at seek point: %v", err)
		}

		fmt.Printf("Pos: %v Byte read: %v \n", bufferEnd+extraBytes, string(singleByteBuf[0]))

		// Check if the byte is a newline (ASCII 10)
		if singleByteBuf[0] == 10 {
			fmt.Println("Found newline character: ascii", singleByteBuf[0])
			return int64(extraBytes), nil
		}

		extraBytes++
	}

}