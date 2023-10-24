FROM golang:1.20.10-bullseye

RUN mkdir /usr/app

WORKDIR /usr/app

RUN apt-get update
RUN apt-get install -y \
        git \
        mercurial

#     \
#        curl \
#        g++
#
#ENV LIBSODIUM_VERSION=1.0.19
#
#RUN \
#    mkdir -p /tmpbuild/libsodium && \
#    cd /tmpbuild/libsodium && \
#    curl -L https://download.libsodium.org/libsodium/releases/libsodium-$LIBSODIUM_VERSION.tar.gz -o $LIBSODIUM_VERSION.tar.gz && \
#    tar xfvz $LIBSODIUM_VERSION.tar.gz && \
#    ls -la /tmpbuild/libsodium && \
#    cd /tmpbuild/libsodium/libsodium-stable && \
#    ./configure && \
#    make && make check && \
#    make install && \
#    mv src/libsodium /usr/local/ && \
#    rm -Rf /tmpbuild/
#
