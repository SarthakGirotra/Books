FROM golang:1.17-alpine
WORKDIR /app
EXPOSE 1322
COPY . .
RUN go mod download
RUN go build -o /books-microservice
CMD ["/books-microservice","-env","production"]