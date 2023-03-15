FROM golang:1.20.2-alpine3.17
COPY . /app
WORKDIR /app 
RUN go mod tidy
RUN go build
CMD [ "./server-watcher" ]