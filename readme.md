# Books

### Requirements
- Docker
- Docker-Compose

### Technologies Used:
- React.js
- Go
- Echo
- Docker
- Kafka
- MongoDB

### Installation
on Linux
```sh
make build
make run
```
on Windows
```sh
docker-compose build
docker-compose up -d
```
React will start at: http://localhost:3000

### Local Installation without Docker
React frontend

_requirements_:
- react.js

```sh
cd frontend
npm install
npm start
```
User microservice

_requirements_:
- go
- mongodb atlas url

rename .env.local.example to .env.local and add atlas url
```sh
cd user-microservice
go mod tidy
go run main.go
```

Books microservice

_requirements_:
- go
- mongodb atlas url

rename .env.local.example to .env.local and add atlas url
```sh
cd books-microservice
go mod tidy
go run main.go
```
### Swagger Documentation

| Docs | Link |
| ------ | ------ |
| user-microservice | http://localhost:1323/swaggerui |
| books-microservice | http://localhost:1322/swaggerui |

### Data Input with CSV

Example csv provided with name Book1.csv

CSV format:

| id | title | story|date|likes (comma seperated userids)|
| ------ | ------ | ------ | ------ | ------ |
|620916dd86cae49baedc7426|abc|story1|15-01-2022 19:04|620916dd86cae49baedc7426,620916dd86cae49baedc7432|
|620916dd86cae49baedc74543|def|story2|12-01-2022 12:04|620916dd86cae49baedc7326,620916dd86cae49baedc7412|

swagger csv input
- link: http://localhost:1322/swaggerui#/books/saveCSVID

API endpoint
- link: http://localhost:1322/saveCSV
- method: POST
- type: multipart/form-data
- name: MyFile