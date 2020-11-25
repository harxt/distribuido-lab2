package main

import (
		
		"fmt"
		
		"math"
		"os"
		
)

func main() {

		fileToBeChunked := "./bigfile.zip" // change here!

		file, err := os.Open(fileToBeChunked)

		if err != nil {
				fmt.Println(err)
				os.Exit(1)
		}

		defer file.Close()

		fileInfo, _ := file.Stat()

		var fileSize int64 = fileInfo.Size()

		const fileChunk = 250 * (1 << 10) // 1 MB, change this to your requirement

		// calculate total number of parts the file will be chunked into

		totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

		fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

		for i := uint64(0); i < totalPartsNum; i++ {

				partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
				partBuffer := make([]byte, partSize)

				file.Read(partBuffer)

				// ENVIAR EL CHUNK

		}}
