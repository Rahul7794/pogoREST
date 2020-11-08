FROM golang:1.14.2-alpine3.11

RUN apk add --update --no-cache alpine-sdk bash git

WORKDIR /app

ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go/vendor
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o pogoREST

RUN chmod +x /app/pogoREST

ENTRYPOINT ["/app/pogoREST"]

