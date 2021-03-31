FROM scratch

COPY minirootfs minirootfs

COPY dkr dkr

COPY dkrd dkrd

CMD ["./dkrd"]
