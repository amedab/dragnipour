FROM golang
RUN git clone https://github.com/amedab/dragnipour.git
WORKDIR dragnipour
RUN go build -o dragnipour main.go
CMD dragnipour
