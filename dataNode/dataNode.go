package main

import (
	"log"
	"net"
	"context"
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"

	pb "../proto"
	"google.golang.org/grpc"
)

// ESTRUCTURAS
type Server struct{}

type NodeInfo struct {
	Id string `json:"id"`
	Ip string `json:"ip"`
	Port string `json:"port"`
}

type Config struct {
	DataNode []NodeInfo `json:"DataNode"`
	NameNode NodeInfo `json:"NameNode"`
}

type Chunk struct {
	paquete bytes  `json:"paquete"`
	nombre string  `json:"nombre"`
	totalPartsNum uint64  `json:"totalPartsNum"`}

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

func obtenerListaIPs() []string{
	var ips []string
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
					ip = v.IP
			case *net.IPAddr:
					ip = v.IP
			}
			ips = append(ips, ip.String())
		}
	}
	return ips
}

func conectarNodo(ip string, port string) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	log.Printf("Intentando iniciar conexión con " + ip + ":" + port)
	host := ip + ":" + port
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		//log.Printf("No se pudo establecer la conexión con " + ip + ":" + strconv.Itoa(port))
		return nil, err
	}
	//log.Printf("Conexión establecida con " + ip + ":" + strconv.Itoa(port))
	return conn, nil
}

func iniciarNodo(port string) {
	// Iniciar servidor gRPC
	log.Printf("Iniciando servidor gRPC en el puerto " + port)
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}
	grpcServer := grpc.NewServer()

	//Registrar servicios en el servidor
	log.Printf("Registrando servicios en servidor gRPC\n")
	pb.RegisterServicioNodoServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

// FUNCIONES DEL SERVER
func (s *Server) ObtenerEstado(ctx context.Context, message *pb.Vacio) (*pb.Estado, error){
	estado := new(pb.Estado)
	estado.Estado = "OK"
	return estado, nil
}


func main(){
	log.Printf("== INICIANDO DATANODE ==")

	// Iniciar variable que mantenga las conexiones establecidas entre nodos
	conexionesNodos := make(map[string]*grpc.ClientConn)

	// Cargar archivo de configuración
	log.Printf("Cargando archivo de configuración")
	var config Config
	config = cargarConfig("config.json")
	log.Printf("Archivo de configuración cargado")

	// Identificar el DataNode correspondiente a la IP de la máquina
	machineIPs := obtenerListaIPs() // Obtener lista de IPs asociadas a la máquina
	for _, dataNode := range config.DataNode { // Iterar sobre las IP configuradas para DataNodes
		_, found := Find(machineIPs, dataNode.Ip)
		if found { // En caso de que la IP configurada coincida con alguna de las IPs de la máquina
			id := dataNode.Id
			ip := dataNode.Ip
			port := dataNode.Port
			conn, err := conectarNodo(ip, port)
			if err != nil{
				// Falla la conexión gRPC 
				log.Fatalf("Error al intentar realizar conexión gRPC: %s", err)
			} else {
				// Registrar servicio gRPC
				c := pb.NewServicioNodoClient(conn)
				estado, err := c.ObtenerEstado(context.Background(), new(pb.Vacio))
				if err != nil {
					//log.Fatalf("Error al llamar a ObtenerEstado(): %s", err)
					log.Printf("DataNode disponible: " + id)
					iniciarNodo(port)
					break
				}
				if estado.Estado == "OK" {
					log.Printf("Almacenando conexión DataNode: " + id)
					conexionesNodos[id] = conn
				}
			}
		}
	}

}
