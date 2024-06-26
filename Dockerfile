FROM golang:alpine

WORKDIR /app

# copy the package manifest and checksum file
COPY go.mod go.sum ./

# install the deps
RUN go mod download

# copy the source files
COPY *.go ./

# build the app
RUN go build -o echo_app .

# expose port 1323
EXPOSE 1323

# run the app
CMD [ "./echo_app" ]