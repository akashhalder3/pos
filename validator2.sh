#!/bin/bash

set -exu
set -o pipefail

# NETWORK_DIR is where all files for the testnet will be stored,
# including logs and storage
NETWORK_DIR=./network
NODE_DIR=$NETWORK_DIR/node0
# We use an empty password. Do not do this in production
geth_pw_file="$NODE_DIR/geth_password.txt"
echo "" > "$geth_pw_file"

# The prysm bootstrap node is set after the first loop, as the first
# node is the bootstrap node. This is used for consensus client discovery
PRYSM_BOOTSTRAP_NODE=enr:-MK4QPeAb8Cv3jXoxeufznBt3DDrZqB9WCE1Ukm-FcshVXPnTC_GrUMaOvo0KoqBnw374Q6WC69Z56ulWYypnwMmQ1qGAY2r4c4jh2F0dG5ldHOIAAAAgAEAAACEZXRoMpBa8xKTIAAAk___________gmlkgnY0gmlwhATwaU-Jc2VjcDI1NmsxoQOjgo9rtH9wyMLGHOAaOH6Z73BX8UYlfSDdlrz8njYyeYhzeW5jbmV0cw-DdGNwgjLIg3VkcIIu4A

# Change this number for your desired number of nodes
NUM_NODES=64

# Port information. All ports will be incremented upon
# with more validators to prevent port conflicts on a single machine
GETH_BOOTNODE_PORT=30301

GETH_HTTP_PORT=8545
GETH_WS_PORT=8546
GETH_AUTH_RPC_PORT=8547
GETH_METRICS_PORT=8548
GETH_NETWORK_PORT=8549

PRYSM_BEACON_RPC_PORT=4000
PRYSM_BEACON_GRPC_GATEWAY_PORT=4100
PRYSM_BEACON_P2P_TCP_PORT=4200
PRYSM_BEACON_P2P_UDP_PORT=4300
PRYSM_BEACON_MONITORING_PORT=4400

PRYSM_VALIDATOR_RPC_PORT=7000
PRYSM_VALIDATOR_GRPC_GATEWAY_PORT=7100
PRYSM_VALIDATOR_MONITORING_PORT=7200

# Set Paths for your binaries. Configure as you wish, particularly
# if you're developing on a local fork of geth/prysm
GETH_BINARY=/usr/bin/geth
GETH_BOOTNODE_BINARY=./dependencies/go-ethereum/build/bin/bootnode
PRYSM_CTL_BINARY=./dependencies/prysm/out/prysmctl
PRYSM_BEACON_BINARY=./dependencies/prysm/out/beacon-chain
PRYSM_VALIDATOR_BINARY=./dependencies/prysm/out/validator

# Start geth execution client for this node
$GETH_BINARY \
      --networkid=${CHAIN_ID:-32382} \
      --http \
      --http.api=eth,net,web3 \
      --http.addr=0.0.0.0 \
      --http.corsdomain="*" \
      --ws \
      --ws.api=eth,net,web3 \
      --ws.addr=0.0.0.0 \
      --ws.origins="*" \
      --authrpc.vhosts="*" \
      --authrpc.addr=0.0.0.0 \
      --authrpc.jwtsecret=$NODE_DIR/execution/jwtsecret \
      --datadir=$NODE_DIR/execution \
      --password=$geth_pw_file \
      --verbosity=3 \
      --syncmode=full \
      --nat extip:20.244.97.158 > "$NODE_DIR/logs/geth.log" 2>&1 &

sleep 5

# Start prysm consensus client for this node
$PRYSM_BEACON_BINARY \
      --datadir=$NODE_DIR/consensus/beacondata \
      --min-sync-peers=0 \
      --genesis-state=$NODE_DIR/consensus/genesis.ssz \
      --bootstrap-node=$PRYSM_BOOTSTRAP_NODE \
      --interop-eth1data-votes \
      --chain-config-file=$NODE_DIR/consensus/config.yml \
      --contract-deployment-block=0 \
      --chain-id=${CHAIN_ID:-32382} \
      --rpc-host=0.0.0.0 \
      --grpc-gateway-host=0.0.0.0 \
      --execution-endpoint=http://0.0.0.0:8551 \
      --accept-terms-of-use \
      --jwt-secret=$NODE_DIR/execution/jwtsecret \
      --suggested-fee-recipient=0x123463a4b065722e99115d6c222f267d9cabb524 \
      --minimum-peers-per-subnet=0 \
      --enable-debug-rpc-endpoints \
      --p2p-host-ip=20.244.97.158 \
      --minimum-peers-per-subnet=0 \
      --monitoring-port=$PRYSM_BEACON_MONITORING_PORT \
      --verbosity=info \
      --slasher \
      --enable-debug-rpc-endpoints > "$NODE_DIR/logs/beacon.log" 2>&1 &

sleep 5

# Start prysm validator for this node. Each validator node will manage 1 validator
# $PRYSM_VALIDATOR_BINARY \
#       --beacon-rpc-provider=localhost:$PRYSM_BEACON_RPC_PORT \
#       --datadir=$NODE_DIR/consensus/validatordata \
#       --accept-terms-of-use \
#       --interop-num-validators=$NUM_NODES \
#       --interop-start-index=0 \
#       --chain-config-file=$NODE_DIR/consensus/config.yml > "$NODE_DIR/logs/validator.log" 2>&1 &
done