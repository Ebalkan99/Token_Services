# Base image
FROM golang:latest

# Çalışma dizini oluştur
WORKDIR /app

# Gerekli dosyaları kopyala
COPY . .

# Gerekli paketleri kur
RUN go mod download

# Uygulamayı derle
RUN go build -o main .

# Uygulamayı çalıştır
CMD ["./main"]
