FROM debian
WORKDIR /app
COPY . /app
RUN apt update
RUN apt install -y socat
ENTRYPOINT ["/bin/bash"]
CMD ["socat.sh"]
