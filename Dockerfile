
#Using latest golang base image  
FROM golang:latest

#Maintainer
LABEL maintainer="Sai Kiran Ambati <saikiranambati942@gmail.com>"

#setting workdir
RUN mkdir -p /usr/local/go/src/http-keystore-app
WORKDIR /usr/local/go/src/http-keystore-app

#Copying files
ADD . /usr/local/go/src/http-keystore-app

#Build app
RUN go build -o keystore cmd/keystore/main.go

#Ports
EXPOSE 8080

# run the binary
CMD ["./keystore"]