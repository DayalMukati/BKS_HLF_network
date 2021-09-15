package contracts

import (
	"encoding/json"
	"fmt"
	"bytes"

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


//GetQueryResult - Return result for given query
func (spc *GoldContract) GetQueryResult(ctx contractapi.TransactionContextInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultAsString queryString:\n%s\n", queryString)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")
	fmt.Println("=======before loop=======")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		fmt.Println("=======inside loop======")
		fmt.Println(queryResponse)
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		fmt.Println(queryResponse.Key)
		fmt.Println(queryResponse.Value)
		// buffer.WriteString("{")
		buffer.WriteString(string(queryResponse.Value))
		// buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("getQueryResultAsString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}
