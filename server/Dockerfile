###############################
# Builder container
###############################

From golang:1.10.2-alpine3.7 as builder
WORKDIR /go/src/app
COPY . .

# Install dependencies...
RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure

# Compile
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


###############################
# Exec container
###############################

From alpine:latest
EXPOSE 8080
RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
COPY --from=builder /go/src/app/app /app
CMD ["/app"]
