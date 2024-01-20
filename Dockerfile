from golang:1.21.6 as build

WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o angry-customer main.go

FROM scratch
COPY --from=build /app/angry-customer app/angry-customer

ENTRYPOINT ["/app/angry-customer"]
