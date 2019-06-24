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
COPY --from=builder /go/src/app/app /app
CMD ["/app"]
