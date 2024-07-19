FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN echo -en "2GIS.CTF{b4ckd00r_upl0ad3d_fl4g_st0l3n}" > /etc/flag.txt

RUN go build -o secstor .

CMD ["/app/secstor"]
