FROM golang:1.26rc2-trixie

RUN apt-get update && apt-get install -y \
    curl \
    build-essential \
    git \
    vim && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENV PATH="/root/.local/bin:${PATH}"

WORKDIR /workspace

CMD ["bash"]
