FROM --platform=$BUILDPLATFORM golang:1.25.4-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags '-w' -o server ./cmd/main.go

# Stage 1 - Final image
FROM --platform=$BUILDPLATFORM alpine:3.22

WORKDIR /app

RUN apk add --no-cache tzdata

# Set timezone
ENV TZ=Asia/Bangkok

# Copy binary
COPY --from=builder /app/server .

EXPOSE 3452

CMD [ "/app/server","start"]
