FROM alpine:3.21.3 AS tailwindcss

RUN apk add --no-cache curl
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v4.0.8/tailwindcss-linux-x64-musl
RUN chmod +x tailwindcss-linux-x64-musl 
RUN mv tailwindcss-linux-x64-musl tailwindcss


FROM golang:1.24-alpine AS build

RUN apk add --no-cache gcompat build-base

WORKDIR /app

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy application code
COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest 

# Cache the dependencies
RUN go mod verify

# Download and install tailwindcss
COPY --from=tailwindcss /tailwindcss /app/tailwindcss
RUN ./tailwindcss -i cmd/web/styles/input.css -o cmd/web/assets/css/output.css

# Generate templates
RUN templ generate

# Build the application
RUN go build -o main cmd/api/main.go


FROM alpine:3.20.1 AS prod
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE ${PORT}
CMD ["./main"]
