package main

import (
	"log"
	"gold-application-chaincode/contracts"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {

	goldContract, err := contractapi.NewChaincode(&contracts.GoldContract{})
	if err != nil {
		log.Panicf("Error creating gold chaincode: %v", err)
	}

	if err := goldContract.Start(); err != nil {
		log.Panicf("Error starting gold chaincode: %v", err)
	}
}
