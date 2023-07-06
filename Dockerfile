FROM golang:1.20-alpine
WORKDIR /opt/notifications-service
COPY . .
RUN apk update && apk upgrade && go build -o notifications-service && rm -rf main.go go.mod go.sum providers .idea .gitignore .git LICENSE Dockerfile README.md
EXPOSE 3000
CMD ["./notifications-service"]
