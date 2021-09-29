package contracts

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// DeleteAsset deletes a given asset from the world state.
func (spc *GoldContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	fmt.Println("Asset id is", id)
	return ctx.GetStub().DelState(id)
}
