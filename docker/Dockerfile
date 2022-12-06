
############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache vips-dev gcc musl-dev gcompat ffmpeg
WORKDIR /waminderbot
COPY . .

# Fetch dependencies.
RUN go mod download
# Install pkger
RUN go install github.com/markbates/pkger/cmd/pkger@latest
# Build the binary.
RUN pkger && go build -o /app/waminderbot

#############################
## STEP 2 build a smaller image
#############################
FROM alpine
RUN apk update && apk add --no-cache vips-dev ffmpeg
WORKDIR /app
# Copy compiled from builder.
RUN mkdir db 
COPY --from=builder /app/waminderbot ./waminderbot
COPY --from=builder /waminderbot/.env ./.env
# Run the binary.
EXPOSE 8000
ENTRYPOINT ["./waminderbot"]