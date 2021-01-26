FROM golang

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

# Download all the dependencies
RUN go mod tidy
RUN go build -o server .

ENV SA_LOGIC_API_URL http://sa-logic:5000

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD "/app/server"
