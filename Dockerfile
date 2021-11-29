FROM golang:1.13

WORKDIR /src/

COPY container.go  /src/
RUN export GO111MODULE=on
RUN go get github.com/gin-gonic/gin

RUN go build -o /bin/demo

EXPOSE 8080

ENTRYPOINT ["/bin/demo"]
