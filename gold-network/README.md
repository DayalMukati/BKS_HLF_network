## Running the test network

You can use the `./network.sh` script to stand up a simple Fabric test network. The test network has two peer organizations with one peer each and a single node raft ordering service. You can also use the `./network.sh` script to create channels and deploy chaincode. For more information, see [Using the Fabric test network](https://hyperledger-fabric.readthedocs.io/en/latest/test_network.html). The test network is being introduced in Fabric v2.0 as the long term replacement for the `first-network` sample.

Before you can deploy the test network, you need to follow the instructions to [Install the Samples, Binaries and Docker Images](https://hyperledger-fabric.readthedocs.io/en/latest/install.html) in the Hyperledger Fabric documentation.

## Deploy chaincode
./network.sh up createChannel -ca -s couchdb
./network.sh deployCC -ccn iac -ccl go -ccp ../chaincode -cci InitGold

./network.sh deployCC -ccn iac -ccl javascript -ccp ../chaincode-js -cci InitGold

## Export path
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/

source ./scripts/setPeerConnectionParam.sh 1 2

source ./scripts/setOrgPeerContext.sh 1

## Init chaincode
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C goldchannel -n iac $PEER_CONN_PARAMS -c '{"function":"InitLedger","Args":[]}'

## Create Metal
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C goldchannel -n iac $PEER_CONN_PARAMS -c '{"function":"AddMetal","Args":["Gold","abc.com"]}'

## Create Metal Group
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C goldchannel -n iac $PEER_CONN_PARAMS -c '{"function":"AddMetalGroup","Args":["metGol6","24k","99","1","24KT"]}'

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C goldchannel -n iac $PEER_CONN_PARAMS -c '{"function":"GetbyId","Args":["metGol0","Metal"]}'

## update metal group
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C goldchannel -n iac $PEER_CONN_PARAMS -c '{"function":"UpdateMetalGroup","Args":["mgr22K0","metSil0","24k","99","1","24KT"]}'

## Add calculation
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C goldchannel -n iac $PEER_CONN_PARAMS -c '{"function":"AddCalculation","Args":["2","Sell/Redeem","1","active"]}'

## Add Collection - ID: collSamp1
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C goldchannel -n iac $PEER_CONN_PARAMS -c '{"function":"AddCollection","Args":["Sample_collection_dev","abc.com","","","google.com","active"]}'

## Add Variety - ID: Varisamp2
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C goldchannel -n iac $PEER_CONN_PARAMS -c '{"function":"AddVariety","Args":["sample dev variety","abc.com","","","google.com","active"]}'


## Add Category - ID: cateSamp3
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C goldchannel -n iac $PEER_CONN_PARAMS -c '{"function":"AddCategory","Args":["Sample Dev Category","abc.com","","","google.com","active"]}'


## Add Diamond
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C goldchannel -n iac $PEER_CONN_PARAMS -c '{"function":"AddDiamond","Args":["circular shaped","2","ok","NavyBlue","good","Develop","active", "Varisamp2","collSamp1","cateSamp3"]}'