FROM redis:latest

RUN apt-get update && apt-get install -y \
    tar \
    wget \
    && rm -rf /var/lib/apt/lists/*

RUN wget https://dl.min.io/client/mc/release/linux-amd64/mc -O /usr/local/bin/mc \
    --tries=3 --timeout=30 && \
    chmod +x /usr/local/bin/mc  

RUN mc --version

RUN mkdir -p /usr/local/etc/redis \
    && chown -R redis:redis /usr/local/etc/redis

COPY redis.conf /usr/local/etc/redis/redis.conf

CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]
