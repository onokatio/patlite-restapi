FROM golang:1.16-alpine

WORKDIR /app

COPY *.go ./

RUN go build -o /patlite main.go

EXPOSE 8085

CMD [ "/patlite" ]
