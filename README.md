# Laboratorio 2 - Sistemas Distribuidos

## Consideraciones
Las máquinas virtuales utilizadas y sus roles son:
- 10.10.28.78  --> NameNode **[NN]**
- 10.10.28.79  --> DataNode 1 **[D1]**
- 10.10.28.8  --> DataNode 2 **[D2]**
- 10.10.28.80  --> DataNode 3 **[D3]**

### PROTOC
Se debe agregar la ruta donde se encuentra *protoc-gen-go* al PATH
```console
export PATH="$PATH:$(go env GOPATH)/bin"
```
Para compilar los archivos .proto
```console
protoc -I proto --go_out=plugins=grpc:proto proto/*.proto
```