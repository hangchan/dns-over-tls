FROM alpine:latest
EXPOSE 53/udp
RUN apk --no-cache add ca-certificates
RUN update-ca-certificates
ENTRYPOINT ["/dns-over-tls"]
COPY ./bin/ /
