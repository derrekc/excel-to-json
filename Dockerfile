FROM docker.io/golang:1.20-bookworm

USER 0

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
RUN go mod tidy
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .
WORKDIR /dist
COPY recommendations.xlsx .
RUN cp /build/main .

USER 1001

CMD ["/dist/main"]
