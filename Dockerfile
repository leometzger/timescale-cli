FROM scratch
COPY timescale /
ENTRYPOINT ["/tsctl"]