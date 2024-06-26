###########################################################################################
# Build deepchain
###########################################################################################
FROM ubuntu:20.04

ENV GO_VERSION '1.17.8'
ENV GO_ARCH 'linux-amd64'
ENV GO_BIN_SHA '980e65a863377e69fd9b67df9d8395fd8e93858e7a24c9f55803421e453f4f99'
ENV DEBIAN_FRONTEND=noninteractive 
ENV DAEMON_HOME /root/.deepchain
ENV DAEMON_RESTART_AFTER_UPGRADE=true
ENV DAEMON_ALLOW_DOWNLOAD_BINARIES=false
ENV DAEMON_LOG_BUFFER_SIZE=1048
ENV UNSAFE_SKIP_BACKUP=true
ENV DAEMON_NAME deepchain
ENV BUILD_DIR /build
ENV PATH /usr/local/go/bin:/root/.cargo/bin:/root/cargo/env:/root/.deepchain/scripts:$PATH

# Install go and required deps
###########################################################################################
RUN apt-get update && apt-get install -y --no-install-recommends wget ca-certificates \
&& wget -O go.tgz https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz \
&& echo "${GO_BIN_SHA} *go.tgz" | sha256sum -c - \
&& tar -C /usr/local -xzf go.tgz \
&& rm go.tgz \
&& go version 

COPY . /sources
WORKDIR /sources

# Install build tools and compile deepchain
###########################################################################################
RUN apt-get -y install --no-install-recommends \
    make gcc g++ \
    curl \
    gnupg \
    git \
    software-properties-common \
&& mkdir -p /deepchain/bin \
# Compile deepchain for genesis version
###########################################################################################
&& cd /sources \
&& make build CUDA_ENABLED=false \
&& cp ./build/deepchain /deepchain/bin/ \
&& cp ./build/deepchain /usr/local/bin \
&& rm -rf ./build \
# Cleanup 
###########################################################################################
&& apt-get purge -y git \
    make \
    gcc g++ \
    curl \
    gnupg \
    python3.8 \
&& go clean --cache -i \
&& apt-get autoremove -y \
&& apt-get clean 

# Copy startup scripts and genesis
###########################################################################################
WORKDIR /
COPY start_script.sh start_script.sh
COPY entrypoint.sh /entrypoint.sh
RUN wget -O /genesis.json https://raw.githubusercontent.com/deep-foundation/deep-chain/main/genesis.json \
&& chmod +x start_script.sh \
&& chmod +x /entrypoint.sh \
&& deepchain version

#  Start
###############################################################################
EXPOSE 26656 26657 1317 9090 26660
ENTRYPOINT ["/entrypoint.sh"]
CMD ["./start_script.sh"]
