FROM scratch

WORKDIR /app

ENV ZTFW_NETWORK="8056c2e21c000001"
ENV ZTFW_FORWARD="127.0.0.1:22"
ENV ZTFW_LISTEN="0.0.0.0:2222"
ENV ZTFW_UDP=""

ENV ZTFW_SERVER=""

COPY --from=builder /app/ztfw /app/ztfw

VOLUME [ "/app/zt-home" ]

ENTRYPOINT ["/usr/bin/ztfw"]
