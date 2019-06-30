FROM golang:1.12

WORKDIR $GOPATH/src/github.com/kenfdev/opa-api-auth-go
COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 1323

# Run the executable
CMD ["opa-api-auth-go"]