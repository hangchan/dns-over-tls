FROM scratch
EXPOSE 53
ENTRYPOINT ["/dns-over-tls"]
COPY ./bin/ /
