# Etapa 1: Construcción
FROM golang:latest AS builder

# Establece el directorio de trabajo en la imagen de construcción
WORKDIR /go/src/app

# Copia y descarga solo los archivos de dependencias para caché eficiente
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código al directorio de trabajo del contenedor
COPY . .

# Compila el código Go
RUN go build -o app .

# Etapa 2: Imagen de producción
FROM alpine:latest

# Instala dependencias necesarias para ejecutar la aplicación
RUN apk --no-cache add ca-certificates

# Copia el binario compilado desde la etapa de construcción
WORKDIR /root/
COPY --from=builder /go/src/app/app .

# Expone el puerto de la aplicación
EXPOSE 8080

# Comando por defecto para ejecutar la aplicación
CMD ["./app"]
