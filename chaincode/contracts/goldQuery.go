package contracts

import (
	"encoding/json"
	"fmt"
	// "bytes"

	//"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//Getter Functions
// Get Metal from ID - WORKS FINE
func (spc *GoldContract) GetMetal(ctx contractapi.TransactionContextInterface, id string) (*Metal, error) {
	metalBytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if metalBytes == nil {
		return nil, fmt.Errorf("metal not found")
	}

	var metal Metal
	err = json.Unmarshal(metalBytes, &metal)
	if err != nil {
		return nil, err
	}

	return &metal, err
}

