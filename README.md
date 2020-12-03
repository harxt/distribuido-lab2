# Laboratorio 2 - Sistemas Distribuidos

## Consideraciones
Las máquinas virtuales utilizadas y sus roles son:
- 10.10.28.78  --> NameNode **[NN-1]**
- 10.10.28.79  --> DataNode 1 **[DN-1]**
- 10.10.28.8  --> DataNode 2 **[DN-2]**
- 10.10.28.80  --> DataNode 3 **[DN-3]**

Las máquinas virtuales cuentan con el repositorio, donde debe realizarse toda la ejecución en la ruta */home/distribuido-lab2/*

## Ejecución
1. Ejecutar los DataNodes en sus respectivas VM.
```console
make Datanode
```

2. Ejecutar el NameNode en sus respectiva VM
```console
make Namenode
```

3. Ejecutar el cliente en cualquiera de las VM
```console
make Cliente 
```

4. Al ejecutar el cliente este presenta un menu con las opciones de subir o descargar archivos.

### PROTOC
Se debe agregar la ruta donde se encuentra *protoc-gen-go* al PATH
```console
export PATH="$PATH:$(go env GOPATH)/bin"
```
Para compilar los archivos .proto
```console
protoc -I proto --go_out=plugins=grpc:proto proto/*.proto
```