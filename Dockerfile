FROM golang:alpine
COPY . ./videoTransmission
WORKDIR /go/videoTransmission/src
RUN go build main.go
CMD ["./main"]

