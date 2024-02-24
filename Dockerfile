# https://docs.docker.com/language/golang/build-images/#create-a-dockerfile-for-the-application

FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY main.go packages ./

# Build
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build .

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

CMD ["main"]