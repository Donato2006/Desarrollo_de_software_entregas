# Sistema de Gestión de Eventos y Entradas

## Descripción

Sistema web de gestión de eventos y venta de entradas desarrollado como Práctico Integrador para la materia Desarrollo de Software.

La aplicación permite a los usuarios registrarse, iniciar sesión, consultar eventos disponibles, comprar entradas, cancelar compras, transferir entradas a otros usuarios y gestionar listas de espera para eventos agotados.

Además, incorpora un rol Administrador que permite gestionar eventos, consultar reportes de ocupación y administrar el catálogo de conciertos.

---

# Integrantes

* Donato Pallotti
* Tiziano Nicolás

---

# Tecnologías Utilizadas

## Backend

* Golang
* Gin Framework
* GORM
* MySQL
* JWT
* bcrypt
* Docker

## Frontend

* React
* React Router DOM
* Axios
* CSS
* Vite

## DevOps

* Docker
* Docker Compose

---

# Funcionalidades Implementadas

## Cliente

### Registro de usuario

Permite crear una cuenta nueva validando:

* Nombre obligatorio
* Correo válido
* Contraseña mínima

### Inicio de sesión

Autenticación mediante JWT.

### Catálogo de eventos

Visualización de todos los conciertos disponibles.

### Detalle de evento

Visualización completa de un concierto específico.

### Compra de entradas

Permite adquirir entradas para un concierto con cupos disponibles.

### Mis Entradas

Visualización de las entradas activas del usuario.

### Cancelación de entradas

Permite cancelar una entrada adquirida.

### Transferencia de entradas

Permite transferir una entrada a otro usuario registrado.

### Lista de espera

Permite anotarse cuando un concierto no tiene cupos disponibles.

### Asignación automática

Cuando un usuario cancela una entrada:

* Se libera un cupo.
* El primer usuario de la lista de espera recibe automáticamente una entrada.

### Notificación por correo

Cuando una entrada es asignada desde la lista de espera se envía un correo electrónico al usuario beneficiado.

---

## Administrador

### Crear conciertos

Alta de nuevos eventos.

### Modificar conciertos

Actualización de información de eventos existentes.

### Eliminar conciertos

Eliminación de eventos del catálogo.

### Reporte de ocupación

Visualización de:

* Cupo total
* Entradas vendidas
* Cupos disponibles
* Porcentaje de ocupación

### Protección de rutas

Acceso exclusivo mediante rol administrador.

---

# Funcionalidad Extra (Bonus)

## Sistema de Lista de Espera

Cuando un concierto se encuentra agotado:

* Los usuarios pueden anotarse en una cola.
* Se registra el orden de llegada.
* Cuando se libera un cupo, se asigna automáticamente al primer usuario de la lista.
* Se envía una notificación por correo electrónico.

---

# Estructura del Proyecto


/
├── Backend
│   ├── controllers
│   ├── dao
│   ├── domain
│   ├── middleware
│   ├── routes
│   ├── services
│   ├── utils
│   └── main.go
│
├── frontend
│   ├── node_modules
│   ├── public
│   ├── src
│   │   ├── assets
│   │   ├── components
│   │   ├── pages
│   │   ├── services
│   │   ├── styles
│   │   ├── App.jsx
│   │   └── main.jsx
│
├── docker-compose.yml
└── README.md


---

# Requisitos Previos

Instalar:

* Docker Desktop
* Git

Opcional para desarrollo local:

* Go 1.26+
* Node.js 22+
* MySQL 8

---

# Ejecución con Docker

Clonar repositorio:

```bash
git clone <URL_DEL_REPOSITORIO>
cd Desarrollo_de_software_entregas
```

Levantar contenedores:

```bash
docker compose up -d --build
```

Verificar ejecución:

```bash
docker ps
```

Deben aparecer:

conciertos_mysql
conciertos_backend
conciertos_frontend


---

# Acceso a la Aplicación

Frontend:

http://localhost:5173

Backend:

http://localhost:8080

---

# Detener Contenedores

```bash
docker compose down
```

---

# Variables Utilizadas

Base de datos:

MYSQL_ROOT_PASSWORD=12345
MYSQL_DATABASE=conciertos_db

Backend:

DB_DSN=root:12345@tcp(mysql:3306)/conciertos_db?charset=utf8mb4&parseTime=True&loc=Local

---

# Endpoints Principales

## Autenticación

```http
POST /register
POST /login
```

## Conciertos

```http
GET /conciertos
GET /conciertos/:id
POST /conciertos
PUT /conciertos/:id
DELETE /conciertos/:id
```

## Entradas

```http
POST /entradas
GET /mis-entradas
DELETE /entradas/:id
POST /transferir-entrada
```

## Lista de Espera

```http
POST /lista-espera
GET /mis-listas-espera
DELETE /lista-espera/:id
```

---

# Decisiones de Diseño

## 1. Uso de JWT

Se eligió JWT para desacoplar la autenticación del almacenamiento de sesiones en servidor y simplificar la autorización mediante roles.

## 2. Asignación Automática desde Lista de Espera

Se implementó un sistema de cola FIFO (First In First Out) para garantizar que los usuarios reciban los cupos liberados respetando el orden de inscripción.

## 3. Dockerización Completa

Se utilizaron Docker y Docker Compose para garantizar que todo el equipo pueda ejecutar exactamente el mismo entorno de desarrollo independientemente del sistema operativo utilizado.

---

# Testing

Para ejecutar todos los tests del backend:

```bash
go test ./...
```

Para generar un reporte de cobertura:

```bash
go test ./... -coverprofile=coverage.out
```

Para visualizar el porcentaje de cobertura por paquete:

```bash
go tool cover -func=coverage.out
```

Ejemplo de salida:

backend/controllers     coverage: 85.7% of statements
backend/services        coverage: 91.2% of statements
total:                  (statements) 88.4%

Para generar un reporte HTML:

```bash
go tool cover -html=coverage.out
```

Esto abrirá un informe visual indicando qué líneas del código fueron cubiertas por los tests.

## Cobertura Alcanzada

Cobertura total del backend: 80.3%

# Estado del Proyecto

Proyecto desarrollado para el Práctico Integrador de Desarrollo de Software 2026.
