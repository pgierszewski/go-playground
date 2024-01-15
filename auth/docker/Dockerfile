FROM golang:latest

RUN mkdir /app

WORKDIR /app

ADD ../ /app

WORKDIR /app/auth

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

ENTRYPOINT CompileDaemon -polling --build="go build -o main" --command=./main