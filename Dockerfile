FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY dbus ./dbus

RUN go build -o /ganesha_exporter

EXPOSE 9587

ENTRYPOINT [ "/ganesha_exporter" ]
