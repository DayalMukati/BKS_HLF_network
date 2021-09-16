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

//Update Metal Group
func (spc *GoldContract) UpdateMetalGroup(ctx contractapi.TransactionContextInterface, metalGroupId string ,metalId string, karatage string, fineness int, referenceId int, shortName string, status string) (*Metalgroup, error) {

	var metalgroup Metalgroup
	metalgroupBytes, err := ctx.GetStub().GetState(metalGroupId)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if metalgroupBytes == nil {
		return nil, fmt.Errorf("the metal group doesn't exists %s", shortName)
	}

	queryString := fmt.Sprintf(`{"selector":{"docType":"%s","Id":"%s"}}`, "Metal", metalId)
	
	metalBytes, err := contract.GetQueryResult(ctx, queryString)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	if metalBytes == nil {
		return nil, fmt.Errorf("the metal group doesn't exists %s", metalId)
	}
	
	metalArray := []Metal{}
	err = json.Unmarshal([]byte(metalBytes), &metalArray)

	if err != nil {
		return nil, err
	}

	//update with new values
	metalgroup.Metals = metalArray
	metalgroup.ShortName = shortName
	metalgroup.Karatage = karatage
	metalgroup.Fineness = fineness
	metalgroup.ReferenceId = referenceId
	metalgroup.Status = status


	//convert Golang to jSon format (JSON Byte Array)
	metalgroupBytes, err = json.Marshal(metalgroup)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(metalGroupId, metalgroupBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &metalgroup, nil
}

//Update Cycle Period
func (spc *GoldContract) UpdateCyclePeriod(ctx contractapi.TransactionContextInterface, cyclePeriodId string, name string, graceperiod int, minWeight int, minValue int, status string) (*CyclePeriod, error) {

	var cycleperiod CyclePeriod
	cycleBytes, err := ctx.GetStub().GetState(cyclePeriodId)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if cycleBytes == nil {
		return nil, fmt.Errorf("the cycle doesn't exists %s", name)
	}

	err = json.Unmarshal(cycleBytes, &cycleperiod)
	if err != nil {
		return nil, err
	}

	cycleperiod.Name = name
	cycleperiod.Graceperiod = graceperiod
	cycleperiod.MinValue = minValue
	cycleperiod.MinWeight = minWeight
	cycleperiod.Status = status
	

	//convert Golang to jSon format (JSON Byte Array)
	cycleBytes, err = json.Marshal(cycleperiod)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(cyclePeriodId, cycleBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &cycleperiod, nil
}