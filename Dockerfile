FROM bearstech/debian:stretch

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    nmap \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY bin/nmap-aas /usr/local/bin

EXPOSE 8888

CMD ["nmap-aas"]