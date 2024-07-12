FROM golang:1.22.5-bullseye as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

FROM builder as downloader

COPY wait-for-it.sh /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh

FROM downloader as final

RUN --mount=type=cache,target="/root/.cache/go-build" go build -o bin .

CMD ["/app/bin"]