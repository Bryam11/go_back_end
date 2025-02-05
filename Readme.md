# Teca Notifications

Teca Notifications is an application to manage notifications, tasks, and user comments.

## Requirements

- Go 1.16 or higher
- PostgreSQL

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/tu_usuario/teca_notifications.git
   cd teca_notifications
   
2. Install the dependencies:
   ```sh
   go mod tidy
   
3. Configura la base de datos en el archivo `db/database.go`:
4. Genera la documentación de la API:
   ```sh
   swag init -g main.go
   ```
esto generará un archivo `docs/docs.go` que contiene la documentación de la API.

1. Ejecuta la aplicación:
   ```sh
   go run main.go
   ```
   La aplicación estará disponible en `http://localhost:8080`.
2. Accede a la documentación de la API en `http://localhost:8080/swagger/index.html`.
## 3. Estructura del proyecto

```plaintext
📦 tu-proyecto
├── 📂 db
│   ├── database.go      # Configuración de la base de datos
│   └── migrations       # Archivos de migración
├── 📂 handlers          # Manejadores de rutas o lógica de negocio
├── 📂 models            # Modelos de datos
├── 📂 routes            # Configuración de rutas de la API
├── main.go             # Punto de entrada de la aplicación
├── go.mod              # Archivo de módulos de Go
├── go.sum              # Archivo de sumas de dependencias
└── README.md           # Este archivo de documentación




   
   