FROM scratch
COPY tsctl /
ENTRYPOINT ["/tsctl"]