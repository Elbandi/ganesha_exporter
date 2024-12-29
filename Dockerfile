FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg

RUN go build -o /ganesha_exporter ./cmd/ganesha_exporter

EXPOSE 9587

ENTRYPOINT [ "/ganesha_exporter" ]
