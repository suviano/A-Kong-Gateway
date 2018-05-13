FROM iron/go

WORKDIR /app

COPY autoconfig /app/

ENTRYPOINT ["./autoconfig"]
