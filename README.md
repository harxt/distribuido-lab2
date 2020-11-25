# Laboratorio 2 - Sistemas Distribuidos

## Librerias Externas
- log

### PROTOC
Comando para compilar los archivos .proto
```console
protoc -I proto --go_out=plugins=grpc:proto proto/*.proto
```