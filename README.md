# matrix-qr-api-go

API en **Go + Fiber v3** que recibe una matriz rectangular, calcula su **factorización QR** y obtiene sus estadísticas desde `matrix-stats-api-node`.

> En construcción: por ahora solo expone `GET /health`.

## Requisitos

- Go 1.27 o superior
- Docker (opcional, para correr el contenedor)

## Correr en local

```bash
go run ./cmd/api
curl http://127.0.0.1:8080/health   # {"status":"ok"}
```

## Correr con Docker

```bash
docker build -t matrix-qr-api-go .
docker run --rm -p 8080:8080 matrix-qr-api-go
curl http://localhost:8080/health   # {"status":"ok"}
```

## Variables de entorno

| Variable | Default | Descripción |
|---|---|---|
| `HOST` | `127.0.0.1` | Dirección en la que escucha el servidor. La imagen Docker usa `0.0.0.0`. |
| `PORT` | `8080` | Puerto HTTP. |

## Repositorios del sistema

| Repositorio | Rol | Tecnología |
|---|---|---|
| [matrix-qr-api-go](https://github.com/basiliok/matrix-qr-api-go) | API de factorización QR (punto de entrada del frontend) | Go + Fiber |
| [matrix-stats-api-node](https://github.com/basiliok/matrix-stats-api-node) | API de estadísticas sobre las matrices Q y R | Node.js + Express + TypeScript |
| [matrix-web-react](https://github.com/basiliok/matrix-web-react) | Frontend | React + TypeScript |
