FROM golang:latest

# working directory on the container
WORKDIR /app

# copy go.mod and go.sum files from your machine to container path 
COPY go.mod go.sum ./

# download all the dependencies. 
RUN go mod download

# copy everything from source(local machine) to destination ( container working directory)
COPY . .

# go build will compile and build the files. -o means output. 
# the below command will build and generate the executable by the name of "main" and place it in container working directory
RUN go build -o main .

# run the executable. User either CMD or ENTRYPOINT
ENTRYPOINT ["./main"]