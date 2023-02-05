FROM golang:1.19.5 as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", "./build/.air.toml"]
