FROM scrolltech/l2geth:scroll-v5.3.0

RUN mkdir -p /l2geth/keystore

ENV RUST_LOG="info"
ENV CHAIN_ID=53077

COPY password /l2geth/
COPY genesis.json /l2geth/
COPY genesis-keystore /l2geth/
COPY l2geth_run.sh /bin/

ENTRYPOINT ["sh", "/bin/l2geth_run.sh"]
