FROM scrolltech/l2geth:scroll-v5.2.0

RUN mkdir -p /l2geth/keystore

ENV RUST_LOG="info"

# TODO: fix this to use the correct chain id
ENV CHAIN_ID=534351


COPY password /l2geth/
COPY genesis.json /l2geth/
COPY genesis-keystore /l2geth/
COPY l2geth_run.sh /bin/

ENTRYPOINT ["sh", "/bin/l2geth_run.sh"]
