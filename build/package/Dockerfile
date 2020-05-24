FROM golang:1.14 AS builder

RUN apt-get update \
    && apt-get install -y --no-install-recommends golang-goprotobuf-dev \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


FROM scratch AS server

EXPOSE 50051

COPY calc_server calc_server
CMD ["./calc_server"]
