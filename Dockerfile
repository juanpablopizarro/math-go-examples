FROM golang

WORKDIR /src
COPY . .
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go build -o /server /src/cmd/server/main.go

# Run the outyet command by default when the container starts.
ENTRYPOINT /server

# Document that the service listens on port 8080.
EXPOSE 8080