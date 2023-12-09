# Use the official Golang image as the base image
FROM golang:1.17-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go server source code into the container
COPY /workspaces/go_server .

# Build the Go server binary
RUN go build -o go-server .

# Use a smaller base image for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files from the build image
COPY --from=build /app/go-server .

# Expose the port that the Go server will run on
EXPOSE 8000

# Command to run the Go server
CMD ["./go-server"]

