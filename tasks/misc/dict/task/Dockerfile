FROM debian

RUN apt update
RUN apt install -y socat golang-go
ENV GOFLAGS="-mod=vendor"
ADD . /app
WORKDIR /app

RUN go build -o main

ENTRYPOINT ["/bin/bash"]
CMD ["socat.sh"]
