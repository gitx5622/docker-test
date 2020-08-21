#  latest golang image. builder is alias
FROM golang:alpine as builder
# install git
RUN apk update && apk add --no-cache git
#set working directory on container
WORKDIR /app
# copy go.mod go.sum files from local to container
COPY go.mod go.sum ./
# downloads all the dependencies
RUN go mod download
# copy source from local to working directory on container
COPY . .
# build the app. After building the app executable is stored in main ( -o main)
# go help build will give more details about each parameters
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#next stage#
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# copy the pre-built binary from 1st/previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/db/migrations ./db/migrations

# run the executable. User either CMD or ENTRYPOINT
ENTRYPOINT ["./main"]
