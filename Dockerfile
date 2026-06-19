FROM --platform=${BUILDPLATFORM} docker.io/library/golang:1.26.2 AS builder

WORKDIR /app/

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags '-w -s' -v -o /function ./

FROM gcr.io/distroless/static-debian13:latest
COPY --from=builder /function /usr/local/bin/function
ENTRYPOINT ["function"]
