FROM scratch
EXPOSE 8080
ENTRYPOINT ["/dns-over-tls"]
COPY ./bin/ /