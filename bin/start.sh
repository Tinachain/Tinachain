nohup /projects/tinachain/geth --nodiscover  \
--maxpeers 3 \
--identity "tinachain" \
--rpc \
--rpcaddr 0.0.0.0 \
--rpccorsdomain "*" \
--datadir "/projects/tinachain/node" \
--port 30303 \
--rpcapi "db,eth,net,web3" \
--networkid 96579 &
