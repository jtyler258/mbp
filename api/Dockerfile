FROM golang:1.17-alpine

WORKDIR /app

COPY . .
# RUN go mod tidy

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8080

ENTRYPOINT CompileDaemon --build="go build -o /mbp-api" --command=/mbp-api
# RUN go build -o /mbp-api

# EXPOSE 8080

# CMD [ "/mbp-api" ]
# CMD [ "go run server.go"]