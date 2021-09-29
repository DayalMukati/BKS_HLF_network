package contracts

import (
	// "encoding/json"
	"fmt"
	"bytes"

	//"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)


// // Get record by ID
func (spc *GoldContract) GetbyId(ctx contractapi.TransactionContextInterface, id string, docType string) (string, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, docType, id)
	return GetQueryResultAsString(ctx, queryString)
}

// Get all records by docType
func (spc *GoldContract) GetAll(ctx contractapi.TransactionContextInterface, docType string) (string, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"%s"}}`, docType)
	return GetQueryResultAsString(ctx, queryString)
}

//GetQueryResultAsString - Return result for given query
func GetQueryResultAsString(ctx contractapi.TransactionContextInterface, queryString string) (string, error) {

	fmt.Printf("- getQueryResultAsString queryString:\n%s\n", queryString)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return "", err
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
			return "", fmt.Errorf(err.Error())
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

	return buffer.String(), nil
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
