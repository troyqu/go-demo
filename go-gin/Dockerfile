FROM golang:1.14-alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
EXPOSE 8888
ENV GOPROXY=https://goproxy.io
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o go-gin
ENTRYPOINT ["./go-gin"]