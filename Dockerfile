FROM golang:alpine
#reference:https://codefresh.io/docs/docs/learn-by-example/golang/golang-hello-world/

## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable

# This container exposes port 8080 to the outside world
EXPOSE 4000

# # Unit tests
 #RUN CGO_ENABLED=0 go test -v
CMD ["/app/main"]
