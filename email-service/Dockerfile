# email-service/Dockerfile

# We use the official golang image, which contains all the 
# correct build tools and libraries. Notice `as builder`,
# this gives this container a name that we can reference later on. 
FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/email-service

# Copy the current code into our workdir
COPY . .

RUN go mod download

# Build the binary, with a few flags which will allow
# us to run this binary in Alpine. 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o email-service main.go

FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/email-service .

CMD ["./email-service"]