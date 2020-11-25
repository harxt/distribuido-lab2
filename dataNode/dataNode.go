package main

import (
	"log"
	"net"
	"context"

	pb "../proto"
	"google.golang.org/grpc"
)

// ESTRUCTURAS
type Server struct{}

// FUNCIONES DEL SERVER
func (s *Server) Enviar(ctx context.Context, message *pb.Vacio) (*pb.Vacio, error){
	return new(pb.Vacio), nil
}

func main(){

	log.Printf("Iniciando servidor gRPC")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	//Registrar servicios en el servidor
	log.Printf("Registrando servicios en servidor gRPC\n")
	pb.RegisterServicioCentralServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}