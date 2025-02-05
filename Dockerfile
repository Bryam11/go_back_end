# Usa una imagen base de Go para construir la aplicación
FROM golang:1.16 as builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al contenedor
COPY . .

# Descarga las dependencias y construye la aplicación
RUN go mod tidy
RUN go build -o teca_notifications main.go

# Usa una imagen base más pequeña para ejecutar la aplicación
FROM alpine:latest

# Instala las dependencias necesarias
RUN apk --no-cache add ca-certificates

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /root/

# Copia el binario construido desde la imagen builder
COPY --from=builder /app/teca_notifications .

# Expone el puerto en el que la aplicación se ejecutará
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./teca_notifications"]