FROM golang:latest

WORKDIR /app

COPY . .

# RUN go mod init myapp
RUN go get -u github.com/gin-gonic/gin
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/mysql

CMD ["go", "run", "main.go"]
