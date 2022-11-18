FROM golang:latest
WORKDIR /
COPY . /
EXPOSE 8080
CMD ["go","run", "main.go"]
