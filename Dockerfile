# syntax=docker/dockerfile:1

FROM golang:1.20-alpine as BUILDER

# Set destination for COPY
WORKDIR /go/src/bee-go-demo

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

RUN go mod vendor

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/src/bee-go-demo/main ./cmd/server/


FROM alpine
EXPOSE 80
WORKDIR /bee-go-demo
COPY --from=BUILDER /go/src/bee-go-demo/main /bee-go-demo
CMD ["/bee-go-demo/main"]