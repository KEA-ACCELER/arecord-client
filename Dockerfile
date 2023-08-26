FROM golang:alpine

WORKDIR /Backend

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .
RUN go build -o /run-client

WORKDIR /
COPY .env .
VOLUME [ "/acceler" ]

CMD [ "/run-client", "start" ]