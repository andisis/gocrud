FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY ./bin/docker/gocrud .

EXPOSE 8000

CMD ["./gocrud"]