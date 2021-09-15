## Running the test network

You can use the `./network.sh` script to stand up a simple Fabric test network. The test network has two peer organizations with one peer each and a single node raft ordering service. You can also use the `./network.sh` script to create channels and deploy chaincode. For more information, see [Using the Fabric test network](https://hyperledger-fabric.readthedocs.io/en/latest/test_network.html). The test network is being introduced in Fabric v2.0 as the long term replacement for the `first-network` sample.

Before you can deploy the test network, you need to follow the instructions to [Install the Samples, Binaries and Docker Images](https://hyperledger-fabric.readthedocs.io/en/latest/install.html) in the Hyperledger Fabric documentation.

## Deploy chaincode
./network.sh deployCC -ccn iac -ccl go -ccp ../chaincode -cci InitGold

## Export path
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/

source ./scripts/setPeerConnectionParam.sh 1 2

source ./scripts/setOrgPeerContext.sh 1

## Init chaincode
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C goldchannel -n iac $PEER_CONN_PARAMS -c '{"function":"InitGold","Args":[]}'

## Create Metal
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C goldchannel -n iac4 $PEER_CONN_PARAMS -c '{"function":"AddMetal","Args":["Gold","abc.com"]}'

## Create Metal Group
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C goldchannel -n iac4 $PEER_CONN_PARAMS -c '{"function":"AddMetalGroup","Args":["metGol0","24k","99","1","24KT"]}'

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C goldchannel -n iac4 $PEER_CONN_PARAMS -c '{"function":"GetbyId","Args":["metGol0","Metal"]}'