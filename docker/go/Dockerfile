FROM golang:1.23.8-bullseye

WORKDIR /app

RUN apt-get update && apt-get install -y \
    net-tools \
    iputils-ping \
    dnsutils \
    curl \
    wget \
    tcpdump \
    netcat-openbsd \
    mariadb-client && \
    rm -rf /var/lib/apt/lists/*

COPY . .

RUN go mod tidy && go mod download

EXPOSE 50051 8000 8080

WORKDIR /app/cmd/ordersystem

CMD ["go", "run", "main.go", "wire_gen.go"]