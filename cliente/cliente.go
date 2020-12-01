package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"encoding/json"
	"io/ioutil"
	"time"

	pb "../proto"
	"google.golang.org/grpc"
)

// ESTRUCTURAS
type NodeInfo struct {
	Id string `json:"id"`
	Ip string `json:"ip"`
	Port string `json:"port"`
}

type Config struct {
	DataNode []NodeInfo `json:"DataNode"`
	NameNode NodeInfo `json:"NameNode"`
}

// FUNCIONES
func cargarConfig(file string) Config {
    var config Config
    configFile, err := ioutil.ReadFile(file)
    if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(configFile, &config)
    return config
}

func dividir(libro string) uint64 {
	fileToBeChunked := libro

	file, err := os.Open(fileToBeChunked)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 250 * (1 << 10)

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)
	for i := uint64(0); i < totalPartsNum; i++ {
		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)
		fmt.Printf("Archivo Dividido")
		// ENVIAR EL CHUNK
	}
	return totalPartsNum

}

func conectarNodo(ip string, port string) *grpc.ClientConn {
	var conn *grpc.ClientConn
	log.Printf("Intentando iniciar conexión con " + ip + ":" + port)
	host := ip + ":" + port
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	return conn
}


func main() {

	log.Printf("= INICIANDO CLIENTE =\n")

	// Cargar archivo de configuración
	log.Printf("Cargando archivo de configuración")
	config := cargarConfig("config.json")
	log.Printf("Archivo de configuración cargado")
	
	// Seleccionar un DataNode de forma aleatoria para conectarse
	rand.Seed(time.Now().UnixNano())
	rand_id := rand.Intn(len(config.DataNode))
	id := config.DataNode[rand_id].Id
	ip := config.DataNode[rand_id].Ip
	port := config.DataNode[rand_id].Port
	log.Printf("DataNode seleccionado: " + id)
	
	// Conectar con servidor DataNode seleccionado
	conn := conectarNodo(ip, port)

	// Registrar servicio gRPC
	c := pb.NewServicioNodoClient(conn)

	log.Printf("Conectado al nodo " +  ip + ":" + port)
	estado, err := c.ObtenerEstado(context.Background(), new(pb.Vacio))
	if err != nil {
		log.Fatalf("Error al llamar a ObtenerEstado(): %s", err)
	}
	log.Printf("Estado del nodo seleccionado: " + estado.Estado)

}
