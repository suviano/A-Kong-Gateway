FROM iron/go

WORKDIR /app

COPY /tmp/autoconfig /app/

ENTRYPOINT ["./autoconfig"]
