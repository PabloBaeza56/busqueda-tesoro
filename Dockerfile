# Utiliza una imagen base de Go
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /go/src/app

# Copia todos los archivos del directorio actual al directorio de trabajo del contenedor
COPY . .

# Compila el código de Go dentro del contenedor
RUN go build -o app .

# Expone el puerto en el que el servidor Go va a escuchar
EXPOSE 8080

# Comando por defecto para ejecutar cuando se inicie el contenedor
CMD ["./app"]