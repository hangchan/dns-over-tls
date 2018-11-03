FROM scratch
EXPOSE 53/udp
ENTRYPOINT ["/dns-over-tls"]
COPY ./bin/ /
