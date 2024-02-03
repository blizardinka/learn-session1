# Build the Vue.js frontend
FROM node:14 as frontend-builder
WORKDIR /app/frontend
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install
COPY frontend/ .
RUN npm run build

# Build the Go backend
FROM golang:1.18 as backend-builder
WORKDIR /app
COPY game/go.mod game/go.sum ./
RUN go mod download
COPY game/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Final stage: Put it all together
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=backend-builder /app/server .

# Assuming the Go server serves static files from a "public" directory
# Adjust the target directory depending on your server's static files location
COPY --from=frontend-builder /app/frontend/dist /root/public

# Define the command to run the application
CMD ["./server"]
