# Use the official Golang image as the base image
FROM golang:1.24-alpine
# Set the working directory inside the container
WORKDIR /app
# Copy the source code into the container
COPY . .
# Build the Gin app
RUN go build -o main .
# Expose the port that the Gin app listens on
EXPOSE 9000
# Run the Gin app when the container starts
CMD ["./main"]
