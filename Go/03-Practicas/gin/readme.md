# Go con gin

En este proyecto creare una api con Go y su framework Gin

## Dependencias

- Gin go ` go get -u github.com/gin-gonic/gin`
- Air Go
- Taildind por plan CDN

## Air

Air es una dependencia de go que sirve para recargar las paginas de Go

1. Instalacion
   `go install github.com/air-verse/air@latest`

2. Configurar air
   `air init` Esto instala un archivo `air.toml` para configurar

3. Ejecutar air con ejecutando -> `air `

## thunder client

Puede ejecutar sus consultas de

- Get (/todos) `http://localhost:4000/api/todos/`
- Get by ID (/todos/:id)`http://localhost:4000/api/todos/1`
- POST (/todos) `http://localhost:4000/api/todos/`
- DELETE (/todos/:id) `http://localhost:4000/api/todos/1`
- PUT (/todos/:id) `http://localhost:4000/api/todos/1`
