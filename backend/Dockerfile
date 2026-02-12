# ========= STAGE 1: Builder (Dapur) =========
FROM golang:1.23-alpine AS builder

# Install git untuk download dependency jika diperlukan
RUN apk add --no-cache git

WORKDIR /app

# 1. Copy file 'belanjaan' library
COPY go.mod go.sum ./
RUN go mod download

# 2. Copy seluruh source code project
COPY . .

# 3. Build aplikasi jadi file binary bernama 'weather-backend'
RUN CGO_ENABLED=0 GOOS=linux go build -o weather-backend .

# ========= STAGE 2: Runner (Meja Saji) =========
FROM alpine:latest

# Tambahkan CA-Certificates supaya bisa akses API OpenWeather via HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 4. Ambil file binary dari stage builder
COPY --from=builder /app/weather-backend .

# 5. PENTING: Ambil file config agar os.ReadFile tidak error
COPY --from=builder /app/.apiConfig .

# 6. Expose port sesuai yang ada di main.go kamu
EXPOSE 8080

# 7. Jalankan aplikasinya
CMD ["./weather-backend"]