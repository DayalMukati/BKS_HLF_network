package main

import (
	"gold-application-chaincode/contracts"

	// "simple-payment-application-chaincode/contracts"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	goldContract := new(contracts.GoldContract)

	cc, err := contractapi.NewChaincode(goldContract)

	if err != nil {
		panic(err.Error())
	}

	if err := cc.Start(); err != nil {
		panic(err.Error())
	}
}
