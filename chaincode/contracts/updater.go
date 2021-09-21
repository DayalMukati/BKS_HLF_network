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

//Update Product
func (spc *GoldContract) UpdateProduct(ctx contractapi.TransactionContextInterface, productId string, name string, imagePath string, videoPath string) (*Product, error) {

	var product Product
	productBytes, err := ctx.GetStub().GetState(productId)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if productBytes == nil {
		return nil, fmt.Errorf("the product doesn't exists %s", name)
	}

	err = json.Unmarshal(productBytes, &product)
	if err != nil {
		return nil, err
	}

	product.Name = name
	product.ImagePath = imagePath
	product.VideoPath = videoPath
	

	//convert Golang to jSon format (JSON Byte Array)
	productBytes, err = json.Marshal(product)
	if err != nil {
		return nil, err
	}

	//put product data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(productId, productBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &product, nil
}

//Update Collection
func (spc *GoldContract) UpdateCollection(ctx contractapi.TransactionContextInterface, collectionId string, collection_name string, img1Path string, img2Path string, img3Path string, videoPath string, status string) (*Collection, error) {

	var collection Collection
	collectionBytes, err := ctx.GetStub().GetState(collectionId)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if collectionBytes == nil {
		return nil, fmt.Errorf("the collection doesn't exists %s", collection_name)
	}

	err = json.Unmarshal(collectionBytes, &collection)
	if err != nil {
		return nil, err
	}

	collection.CollectionName = collection_name
	collection.Img1path = img1Path
	collection.Img2path = img2Path
	collection.Img3path = img3Path
	collection.VideoPath = videoPath
	collection.Status = status
	

	//convert Golang to jSon format (JSON Byte Array)
	collectionBytes, err = json.Marshal(collection)
	if err != nil {
		return nil, err
	}

	//put collection data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(collectionId, collectionBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &collection, nil
}

//Update Category
func (spc *GoldContract) UpdateCategory(ctx contractapi.TransactionContextInterface, categoryId string, category_name string, img1Path string, img2Path string, img3Path string, videoPath string, status string) (*Category, error) {

	var category Category
	categoryBytes, err := ctx.GetStub().GetState(categoryId)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if categoryBytes == nil {
		return nil, fmt.Errorf("the category doesn't exists %s", category_name)
	}

	err = json.Unmarshal(categoryBytes, &category)
	if err != nil {
		return nil, err
	}

	category.CategoryName = category_name
	category.Img1path = img1Path
	category.Img2path = img2Path
	category.Img3path = img3Path
	category.VideoPath = videoPath
	category.Status = status
	

	//convert Golang to jSon format (JSON Byte Array)
	categoryBytes, err = json.Marshal(category)
	if err != nil {
		return nil, err
	}

	//put category data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(categoryId, categoryBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &category, nil
}