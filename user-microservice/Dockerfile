FROM golang:1.17-alpine
WORKDIR /app
EXPOSE 1323
COPY . .
RUN go mod download
RUN go build -o /user-microservice
CMD ["/user-microservice","-env","production"]