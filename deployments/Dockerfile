FROM golang:1.23.6-bookworm

EXPOSE 3000 

COPY cmd cmd
COPY internal internal
COPY ./go.sum .
COPY ./go.mod .

CMD go run ./cmd/service/main.go
