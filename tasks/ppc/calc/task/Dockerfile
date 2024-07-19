FROM python:3.8-bookworm
COPY ./requirements.txt /app/requirements.txt
WORKDIR /app
RUN pip install -r requirements.txt
COPY . /app
RUN apt update
RUN apt install -y socat
ENTRYPOINT ["/bin/bash"]
CMD ["socat.sh"]