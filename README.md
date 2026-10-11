# matrix-qr-api-go

API en **Go + Fiber v3** que recibe una matriz rectangular, calcula su **factorización QR** y obtiene sus estadísticas desde `matrix-stats-api-node`.

> En construcción. Ya existen `GET /health`, la documentación en `/docs` y la factorización QR (`internal/matrix`). Próximo paso: el endpoint `POST /api/v1/qr`.

Desplegada en Azure App Service: https://app-matrix-qr-api-go-h8cmdmddgkezh2gm.centralus-01.azurewebsites.net/health

## Requisitos

- Go 1.27 o superior
- Docker (opcional, para correr el contenedor)

## Correr en local

```bash
go run ./cmd/api
curl http://127.0.0.1:8080/health   # {"status":"ok","service":"matrix-qr-api-go"}
```

## Estructura

```
cmd/api/           # main: configuración, rutas y arranque del servidor
internal/matrix/   # validación de la matriz y factorización QR (sin dependencias de Fiber)
docs/              # especificación Swagger generada por swag (no se edita a mano)
```

## Factorización QR

- **Algoritmo:** reflexiones de Householder, implementadas en `internal/matrix/qr.go`.
- **Forma reducida:** para una matriz m×n y k = min(m, n), Q es m×k (columnas ortonormales) y R es k×n (triangular superior). Es la forma por defecto de `numpy.linalg.qr`.
- **Resultado único:** la diagonal de R es ≥ 0, los valores bajo la diagonal son ceros exactos y no se devuelve `-0`.
- **Validación:** la matriz no puede estar vacía, todas las filas deben tener la misma cantidad de columnas y el tamaño máximo es 100×100.

## Documentación (Swagger)

Con el servidor levantado: http://127.0.0.1:8080/docs/index.html

Cada endpoint se documenta con anotaciones de [swag](https://github.com/swaggo/swag) sobre su handler. Después de agregar o cambiar anotaciones, regenera el paquete `docs/`:

```bash
go tool swag init --dir ./cmd/api --generalInfo main.go --output ./docs --outputTypes go
```

## Correr con Docker

```bash
docker build -t matrix-qr-api-go .
docker run --rm -p 8080:8080 matrix-qr-api-go
curl http://localhost:8080/health   # {"status":"ok","service":"matrix-qr-api-go"}
```

## Despliegue

Cada push a `main` ejecuta `.github/workflows/build-and-push-image.yml`, que construye la imagen y la sube a Azure Container Registry con los tags `latest` y el SHA del commit. Un webhook del registro hace que App Service descargue la imagen nueva y reinicie.

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
