# Gunakan base image golang
FROM golang:1.20-alpine

# Set working directory
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Tentukan port yang akan digunakan
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
