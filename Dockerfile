FROM golang:1.20


# Set the working directory in the container
WORKDIR /app

# Copy the application code to the container
COPY . .

# Download the required dependencies
RUN go mod download

# Build the application
RUN go build

# Expose port 8080 for the Gin web server to listen on
EXPOSE 8080

# Specify the command to run when the container starts
CMD cd /app && go run .