FROM alpine:latest
EXPOSE 53
RUN apk add — no-cache ca-certificates
RUN update-ca-certificates
ENTRYPOINT ["/dns-over-tls"]
COPY ./bin/ /
