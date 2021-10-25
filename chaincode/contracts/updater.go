package contracts

import (
	"encoding/json"
	"fmt"
	//"bytes"

	//"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//Update Metal
func (spc *GoldContract) UpdateMetal(ctx contractapi.TransactionContextInterface, metalId string, name string, icon string) (*Metal, error) {

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
	metal.Icon = icon
	metal.UpdatedAt = dtformated

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

	return &metal, nil
}

//Update Metal Group
func (spc *GoldContract) UpdateMetalGroup(ctx contractapi.TransactionContextInterface, metalGroupId string ,metalId string, karatage string, fineness float64, referenceId int, shortName string, status string) (*Metalgroup, error) {

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
	metalgroup.UpdatedAt = dtformated

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
	return &metalgroup, nil
}

//Update Cycle Period
func (spc *GoldContract) UpdateCyclePeriod(ctx contractapi.TransactionContextInterface, cyclePeriodId string, name string, graceperiod int, minWeight int, minValue int) (*CyclePeriod, error) {

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
	cycleperiod.UpdatedAt = dtformated

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
	return &cycleperiod, nil
}

//Update Product
func (spc *GoldContract) UpdateProduct(ctx contractapi.TransactionContextInterface, productId string, name string, images string, video string) (*Product, error) {

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
	product.Images = images
	product.Video = video
	product.UpdatedAt = dtformated

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

	return &product, nil
}

//Update Collection
func (spc *GoldContract) UpdateCollection(ctx contractapi.TransactionContextInterface, collection_id string, img1 string, img2 string, img3 string, video string, collection_name string) (*Collection, error) {

	var collection Collection
	collectionBytes, err := ctx.GetStub().GetState(collection_id)

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
	collection.Img1 = img1
	collection.Img2 = img2
	collection.Img3 = img3
	collection.Video = video
	collection.UpdatedAt = dtformated

	//convert Golang to jSon format (JSON Byte Array)
	collectionBytes, err = json.Marshal(collection)
	if err != nil {
		return nil, err
	}

	//put collection data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(collection_id, collectionBytes)
	if err != nil {
		return nil, err
	}

	return &collection, nil
}

//Update Category
func (spc *GoldContract) UpdateCategory(ctx contractapi.TransactionContextInterface, category_id string, category_name string, img1 string, img2 string, img3 string, video string) (*Category, error) {

	var category Category
	categoryBytes, err := ctx.GetStub().GetState(category_id)

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
	category.Img1 = img1
	category.Img2 = img2
	category.Img3 = img3
	category.Video = video
	category.UpdatedAt = dtformated

	//convert Golang to jSon format (JSON Byte Array)
	categoryBytes, err = json.Marshal(category)
	if err != nil {
		return nil, err
	}

	//put category data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(category_id, categoryBytes)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

//Update Variety
func (spc *GoldContract) UpdateVariety(ctx contractapi.TransactionContextInterface, variety_id string, variety_name string, img1 string, img2 string, img3 string, video string) (*Variety, error) {

	var variety Variety
	varietyBytes, err := ctx.GetStub().GetState(variety_id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if varietyBytes == nil {
		return nil, fmt.Errorf("the variety doesn't exists %s", variety_name)
	}

	err = json.Unmarshal(varietyBytes, &variety)
	if err != nil {
		return nil, err
	}

	variety.VarietyName = variety_name
	variety.Img1 = img1
	variety.Img2 = img2
	variety.Img3 = img3
	variety.Video = video
	variety.UpdatedAt = dtformated

	//convert Golang to jSon format (JSON Byte Array)
	varietyBytes, err = json.Marshal(variety)
	if err != nil {
		return nil, err
	}

	//put category data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(variety_id, varietyBytes)
	if err != nil {
		return nil, err
	}

	return &variety, nil
}

// //Update Charges
// func (spc *GoldContract) UpdateCharges(ctx contractapi.TransactionContextInterface, variety_id string, variety_name string, img1 string, img2 string, img3 string, video string) (*Variety, error) {

// 	var variety Variety
// 	varietyBytes, err := ctx.GetStub().GetState(variety_id)

// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read data from world state: %v", err)
// 	}
// 	//check if ID already exists (return the state of the ID by checking the world state)
// 	if varietyBytes == nil {
// 		return nil, fmt.Errorf("the variety doesn't exists %s", variety_name)
// 	}

// 	err = json.Unmarshal(varietyBytes, &variety)
// 	if err != nil {
// 		return nil, err
// 	}

// 	variety.VarietyName = variety_name
// 	variety.Img1 = img1
// 	variety.Img2 = img2
// 	variety.Img3 = img3
// 	variety.Video = video
// 	variety.UpdatedAt = dtformated

// 	//convert Golang to jSon format (JSON Byte Array)
// 	varietyBytes, err = json.Marshal(variety)
// 	if err != nil {
// 		return nil, err
// 	}

// 	//put category data unto the Ledger (key value pair)
// 	err = ctx.GetStub().PutState(variety_id, varietyBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &variety, nil
// }