package main

import (
	"log"
	"context"
	"fmt"
	"math"
	"os"

	pb "../proto"
	"google.golang.org/grpc"
)


func main() {
	log.Printf("= INICIANDO CLIENTE =\n")

	// Conectar con servidor DataNode
	ip := "localhost"
	var conn *grpc.ClientConn
	host := ip + ":9000"
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	c := pb.NewServicioCentralClient(conn)

	_, err = c.Enviar(context.Background(), new(pb.Vacio))
	if err != nil {
		log.Fatalf("Error al llamar a Enviar(): %s", err)
	}


	fileToBeChunked := "./peter.pdf"

	file, err := os.Open(fileToBeChunked)
	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 250 * (1 << 10) 

	// calculate total number of parts the file will be chunked into
	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)
	for i := uint64(0); i < totalPartsNum; i++ {
		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)
		// ENVIAR EL CHUNK
	
	}
}
