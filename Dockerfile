FROM scratch
COPY iplookup /
ENTRYPOINT ["/iplookup"]