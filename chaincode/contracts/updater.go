package contracts

import (
	"encoding/json"
	"fmt"
	//"bytes"

	//"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)


//Update Metal 
func (spc *GoldContract) UpdateMetal(ctx contractapi.TransactionContextInterface, metalId string, name string, imagePath string) (*Metal, error) {

	var metal Metal
	metalBytes, err := ctx.GetStub().GetState(metalId)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if metalBytes == nil {
		return nil, fmt.Errorf("the metal doesn't exists %s", name)
	}

	err = json.Unmarshal(metalBytes, &metal)
	if err != nil {
		return nil, err
	}

	metal.Name = name
	metal.ImagePath = imagePath
	

	//convert Golang to jSon format (JSON Byte Array)
	metalBytes, err = json.Marshal(metal)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(metalId, metalBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &metal, nil
}