FROM golang:1.27-alpine AS build
WORKDIR /source

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /output/qr-api ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /output/qr-api /qr-api

ENV HOST=0.0.0.0
ENV PORT=8080
EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/qr-api"]
