package main

import (
	"log"
	"teca_notifications/api"
	"teca_notifications/db"
)

func main() {
	// Conectar a la base de datos
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	// Configurar rutas
	r := api.SetupRoutes(dbConn)

	// Iniciar servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
