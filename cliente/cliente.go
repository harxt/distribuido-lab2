package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
	"strconv"
	"bufio"
	"os/exec"
	"strings"

	pb "../proto"
	"google.golang.org/grpc"
)
var err error
// ESTRUCTURAS
type NodeInfo struct {
	Id   string `json:"id"`
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

type Config struct {
	DataNode []NodeInfo `json:"DataNode"`
	NameNode NodeInfo   `json:"NameNode"`
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

func dividirArchivo(path string) [][]byte {
	fileToBeChunked := path
	log.Printf("Cargando archivo: " + path)
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
	chunks := make([][]byte, totalPartsNum)
	log.Printf("Dividiendo archivo en %d chunks.\n", totalPartsNum)
	for i := uint64(0); i < totalPartsNum; i++ {
		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)
		chunks[i] = partBuffer
		//fmt.Printf("Archivo Dividido")
		//fileName := nombre + strconv.FormatUint(i, 10)
		//_, err := os.Create(fileName)
		
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)}
		//ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
		//fmt.Println("Split to :", fileName)
		// ENVIAR EL CHUNK
	}
	return chunks

}

func reunir(nombre string, totalPartsNum uint64) {
	newFileName := nombre
	_,err = os.Create(newFileName)
	
	if err != nil{
		fmt.Println(err)
		os.Exit(1)}
	file, err :=os.OpenFile (newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		}	
	var writePosition int64 = 0
	
	for j := uint64(0) ; j < totalPartsNum; j++ {
		
		currentChunkFileName := nombre + strconv.FormatUint(j,10)
		newFileChunk, err := os.Open(currentChunkFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)}
		defer newFileChunk.Close()
		
		chunkInfo, err := newFileChunk.Stat()
		
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}	
		var chunkSize int64 = chunkInfo.Size()
		chunkBufferBytes := make([]byte, chunkSize)
		
		fmt.Println("añadiendo en la posicion : [", writePosition, "] bytes")
		writePosition = writePosition + chunkSize
		
		reader := bufio.NewReader(newFileChunk)
		_, err =reader.Read(chunkBufferBytes)
		
		if err != nil {
			fmt.Println(err)
			os.Exit(1)}
		n, err := file.Write (chunkBufferBytes)
		
		if err != nil{
			fmt.Println(err)
			os.Exit(1)}
		file.Sync()
		
		chunkBufferBytes = nil
		fmt.Println("Escritos ", n, "bytes")
		fmt.Print("Recombinando la parte ", j, " dentro del archivo", newFileName)
	}
	file.Close()
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

func clearConsole(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func randomDataNode(config *Config) (string, string, string) {
	rand.Seed(time.Now().UnixNano())
	rand_id := rand.Intn(len(config.DataNode))
	id := config.DataNode[rand_id].Id
	ip := config.DataNode[rand_id].Ip
	port := config.DataNode[rand_id].Port
	return id, ip, port
}


func main() {

	log.Printf("= INICIANDO CLIENTE =\n")

	// Cargar archivo de configuración
	log.Printf("Cargando archivo de configuración")
	config := cargarConfig("config.json")
	log.Printf("Archivo de configuración cargado")

	//Recibir como input la operación a realizar
	reader := bufio.NewReader(os.Stdin)
	menuLoop: for {
		fmt.Print("\n=== MENU ===\n")
		fmt.Print("0 ] Salir\n")
		fmt.Print("1 ] Cargar archivo\n")
		fmt.Print("2 ] Descargar archivo\n")
		fmt.Print("Ingresar operación a realizar:\n")
		fmt.Print("=> ")
		op, _ := reader.ReadString('\n')
		op = strings.Replace(op, "\n", "", -1)

		switch op{
		case "0":
			clearConsole()
			log.Printf("\n== CERRANDO CLIENTE ==")
			os.Exit(1)
		case "1":
			clearConsole()
			fmt.Println("Cargar archivo...")
			fmt.Print("Nombre del archivo a cargar: ")
			rutafile, _ := reader.ReadString('\n')
			rutafile = strings.Replace(rutafile, "\n", "", -1)
	
			chunks := dividirArchivo(rutafile)
			log.Printf("Archivo divido")
			fmt.Printf("Primer chunk[0]: %-v \n", chunks[0][0])

			//Conectar DataNode
			break menuLoop
		case "2":
			clearConsole()
			fmt.Println("Descargar archivo...")
			//Conectar Namenode
			break menuLoop
		}
	}

	// Seleccionar un DataNode de forma aleatoria para conectarse
	id, ip, port := randomDataNode(&config)
	log.Printf("DataNode seleccionado: " + id)
	// Conectar con servidor DataNode seleccionado
	conn := conectarNodo(ip, port)
	// Registrar servicio gRPC
	c := pb.NewServicioNodoClient(conn)

	log.Printf("Conectado al nodo " + ip + ":" + port)
	estado, err := c.ObtenerEstado(context.Background(), new(pb.Vacio))
	if err != nil {
		log.Fatalf("Error al llamar a ObtenerEstado(): %s", err)
	}
	log.Printf("Estado del nodo seleccionado: " + estado.Estado)

}
