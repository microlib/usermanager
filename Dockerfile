FROM docker.io/microlib/golang-base:1.7

# This will change depending on each microservice entry point
# This will change to the binary start in production based images
CMD ["/bin/sh", "go run main.go"]
