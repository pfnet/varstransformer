FROM --platform=${BUILDPLATFORM} docker.io/library/golang:1.26.4 AS builder

WORKDIR /app/

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags '-w -s' -v -o /function ./

FROM gcr.io/distroless/static:latest
COPY --from=builder /function /usr/local/bin/function
ENTRYPOINT ["function"]
