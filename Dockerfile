#syntax=docker/dockerfile:1

# we use base image having standard go toolchain 
FROM golang:1.19 

# define our working directory in our image 
WORKDIR /app 

# copy go.mod and go.sum into our image
COPY go.mod go.sum ./

# download go dependencies from our go.mod and go.sum files
RUN go mod download

#Copy source code
COPY main.go ./

RUN go build -o /docker-stratos

CMD ["/docker-stratos"]
