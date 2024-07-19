FROM debian
WORKDIR /app
COPY . /app
RUN apt update
RUN apt install -y gcc socat
RUN gcc main.c
ENTRYPOINT ["/bin/bash"]
CMD ["socat.sh"]
