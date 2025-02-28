FROM alpine:3.21.3 AS download-tailwind
RUN apk add --no-cache curl \
&& curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v4.0.8/tailwindcss-linux-x64-musl \
&& chmod +x tailwindcss-linux-x64-musl \
&& mv tailwindcss-linux-x64-musl tailwindcss


FROM golang:1.24-alpine AS install-dependencies
WORKDIR /app
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download


FROM ghcr.io/a-h/templ:latest AS generate-template-stage
COPY --chown=65532:65532 . /app
WORKDIR /app
RUN ["templ", "generate"]


FROM golang:1.24-alpine AS build
RUN apk add --no-cache gcompat build-base
COPY --from=generate-template-stage /app /app
WORKDIR /app
# Copy tailwindcss executable file from tailwindcss stage
COPY --from=download-tailwind /tailwindcss /app/tailwindcss
RUN ./tailwindcss -i cmd/web/styles/input.css -o cmd/web/assets/css/output.css
# Build the application
RUN go build -o main cmd/api/main.go


FROM alpine:3.20.1 AS prod
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE ${PORT}
CMD ["./main"]
