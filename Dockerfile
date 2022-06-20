# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.17-alpine

# create work dir inside image
WORKDIR /app

# copy files to image dir @ /app
COPY go.mod ./
COPY go.sum ./

# install mod dependencies
RUN go mod download

# copy repo files to image
COPY . ./

# build executable
RUN go build -o /bank-server

# EXPOSE port to connect to image
EXPOSE 8080

ENV GIN_MODE=release

# run executable
CMD ["/bank-server"]