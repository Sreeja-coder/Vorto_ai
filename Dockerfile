FROM golang:1.16.5-alpine3.13
WORKDIR /app
COPY . .
RUN go build -o vorto vorto.go
RUN cd /app
EXPOSE 8080
CMD ["./vorto"]