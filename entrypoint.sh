#!/bin/sh

if [ ! -d "/root/.deepchain/" ]
then
  mkdir /root/.deepchain/
  mkdir /root/.deepchain/config/
  /usr/local/bin/deepchain init ${NODE_MONIKER}
fi

if [ ! -f "/root/.deepchain/config/node_key.json" ]
then
  /usr/local/bin/deepchain init ${NODE_MONIKER}
  cp /genesis.json /root/.deepchain/config/
fi

if [ ! -f "/root/.deepchain/config/genesis.json" ]
then
  cp /genesis.json /root/.deepchain/config/
fi

if [ "$2" = 'init' ]; then
  return 0
else
  exec "$@"
fi
