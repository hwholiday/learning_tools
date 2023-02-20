FROM alpine:3.17

ARG ROCKSDB_VERSION=v7.9.2

RUN apk update && \
    apk add --no-cache zlib-dev bzip2-dev lz4-dev snappy-dev zstd-dev gflags-dev && \
    apk add --no-cache build-base linux-headers git bash perl && \
    mkdir /usr/src && \
    cd /usr/src && \
    git clone --depth 1 --branch ${ROCKSDB_VERSION} https://github.com/facebook/rocksdb.git && \
    cd /usr/src/rocksdb && \
    sed -i 's/install -C/install -c/g' Makefile && \
    make -j4 shared_lib && \
    make install-shared && \
    apk del build-base linux-headers git bash perl && \
    rm -rf /usr/src/rocksdb

WORKDIR /