FROM golang
RUN go install github.com/amedab/dragnipour@latest
CMD ~/go/bin/dragnipour
