# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.17-alpine AS builder
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
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bank-server .


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/bank-server ./ 

EXPOSE 8080

ENV GIN_MODE=release

CMD ["./bank-server"]
