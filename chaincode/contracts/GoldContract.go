package contracts

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	//"log"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// GoldContract contract for handling writing and reading from the world state
type GoldContract struct {
	contractapi.Contract
}

//Metal
type Metal struct {
	DocType					string		`json:"docType"`
	Id		 				string		`json:"Id"`
	Name					string		`json:"name"`
	ImagePath				string		`json:"imagePath"`
}

// MetalGroup : The asset being tracked on the chain
type Metalgroup struct {
	DocType           		string 		`json:"docType"`
	Id						string 		`json:"Id"`
	Metals		            []Metal 	`json:"metalId"`
	Karatage			    string 		`json:"karatage"`
	Fineness             	int			`json:"fineness"`
	ReferenceId            	int 		`json:"referenceId"`
	ShortName				string		`json:"shortName"`
	Status					string		`json:"status"`
}


var count int 
var metalGroupCount int

var contract GoldContract

// Init and Creator Functions.
func (spc *GoldContract) InitGold(ctx contractapi.TransactionContextInterface) error {
	count = 0
	return nil
}

//Add new Metal 
func (spc *GoldContract) AddMetal(ctx contractapi.TransactionContextInterface, name string, imagePath string) (*Metal, error) {

	id, _ := contract.IDGenerator("metal", name, count)
	metalBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if metalBytes != nil {
		return nil, fmt.Errorf("the metal already exists %s", name)
	}


	//defince structs
	metal := Metal{
		DocType:           "Metal",
		Id:		    		id,
		Name:         		name,
		ImagePath: 			imagePath,
	}

	//convert Golang to jSon format (JSON Byte Array)
	metalBytes, err = json.Marshal(metal)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(id, metalBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &metal, nil
}

//Add new Metal Group
func (spc *GoldContract) AddMetalGroup(ctx contractapi.TransactionContextInterface, metalId string, karatage string, fineness int, referenceId int, shortName string) (*Metalgroup, error) {

	id, _ := contract.IDGenerator("mgroup", shortName, count)
	metalgroupBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if metalgroupBytes != nil {
		return nil, fmt.Errorf("the metal already exists %s", shortName)
	}

	queryString := fmt.Sprintf(`{"selector":{"docType":"%s","Id":"%s"}}`, "Metal", metalId)
	
	metalBytes, err := contract.GetQueryResult(ctx, queryString)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	
	metalArray := []Metal{}
	err = json.Unmarshal([]byte(metalBytes), &metalArray)

	if err != nil {
		return nil, err
	}
	//defince structs
	metalGroup := Metalgroup{
		DocType:           		"MetalGroup",
		Id:						id,
		Metals: 				metalArray,			
		ShortName:         		shortName,
		Karatage:				karatage,
		Fineness:				fineness,
		ReferenceId:			referenceId,
		Status:					"Active",
	}

	//convert Golang to jSon format (JSON Byte Array)
	metalgroupBytes, err = json.Marshal(metalGroup)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(id, metalgroupBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &metalGroup, nil
}

//CyclePeriod:
type CyclePeriod struct {
	DocType					string 		`json:"docType"`
	Name 					string		`json:"name"`
	Id						string		`json:"Id"`
	Graceperiod				int			`json:"graceperiod"`
	MinWeight				int			`json:"minWeight"`
	MinValue				int			`json:"minValue"`
	Status					string		`json:"status"`
}

//Add Cycle Period
func (spc *GoldContract) AddCyclePeriod(ctx contractapi.TransactionContextInterface, name string, graceperiod int, minWeight int, minValue int, status string) (*CyclePeriod, error) {

	id, _ := contract.IDGenerator("cyc", name, count)
	cycleBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if cycleBytes != nil {
		return nil, fmt.Errorf("the cycle period already exists %s", name)
	}

	//defince structs
	cycleperiod := CyclePeriod{
		DocType:           "CyclePeriod",
		Id:		    		id,
		Name:         		name,
		Graceperiod: 		graceperiod,
		MinWeight: 			minWeight,
		MinValue: 			minValue,
		Status:				status,
	}

	//convert Golang to jSon format (JSON Byte Array)
	cycleBytes, err = json.Marshal(cycleperiod)
	if err != nil {
		return nil, err
	}

	//put cycle period data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(id, cycleBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &cycleperiod, nil
}


//Product
type Product struct {
	DocType 				string		`json:"docType"`
	Id						string		`json:"productId"`
	Name					string		`json:"name"`
	ImagePath				string		`json:"imagePath"`
	VideoPath				string		`json:"videoPath"`
}
//Add new Product 
func (spc *GoldContract) AddProduct(ctx contractapi.TransactionContextInterface, name string, imagePath string, videoPath string) (*Product, error) {

	id, _ := contract.IDGenerator("product", name, count)
	productBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if productBytes != nil {
		return nil, fmt.Errorf("the product already exists %s", name)
	}


	//defince structs
	product := Product{
		DocType:           "Product",
		Id:		         	id,
		Name:         		name,
		ImagePath: 			imagePath,
		VideoPath:			videoPath,
	}

	//convert Golang to jSon format (JSON Byte Array)
	productBytes, err = json.Marshal(product)
	if err != nil {
		return nil, err
	}

	//put product data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(id, productBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &product, nil
}

//Collection
type Collection struct {
	DocType					string		`json:"docType"`
	Id						string		`json:"collectionId"`
	CollectionName			string		`json:"collectionName"`
	Img1path				string		`json:"img1path"`
	Img2path				string		`json:"img2path"`
	Img3path				string		`json:"img3path"`
	VideoPath				string		`json:"videoPath"`

}
//Add new collection
func (spc *GoldContract) AddCollection(ctx contractapi.TransactionContextInterface, collection_name string, img1Path string, img2Path string, img3Path string, videoPath string) (*Collection, error) {

	id, _ := contract.IDGenerator("collection", collection_name, count)
	collectionBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if collectionBytes != nil {
		return nil, fmt.Errorf("the collection already exists %s", collection_name)
	}


	//defince structs
	collection := Collection{
		DocType:           	"Collection",
		Id:		         	id,
		CollectionName:     collection_name,
		Img1path: 			img1Path,
		Img2path:			img2Path,
		Img3path:			img3Path,
		VideoPath:			videoPath,
	}

	//convert Golang to jSon format (JSON Byte Array)
	collectionBytes, err = json.Marshal(collection)
	if err != nil {
		return nil, err
	}

	//put collection data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(id, collectionBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &collection, nil
}

//Category
type Category struct {
	DocType					string		`json:"docType"`
	Id						string		`json:"categoryId"`
	CategoryName			string		`json:"categoryName"`
	Img1path				string		`json:"img1path"`
	Img2path				string		`json:"img2path"`
	Img3path				string		`json:"img3path"`
	VideoPath				string		`json:"videoPath"`

}
//Add new categories
func (spc *GoldContract) AddCategory(ctx contractapi.TransactionContextInterface, category_name string, img1Path string, img2Path string, img3Path string, videoPath string) (*Category, error) {

	id, _ := contract.IDGenerator("category", category_name, count)
	categoryBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if categoryBytes != nil {
		return nil, fmt.Errorf("the category already exists %s", category_name)
	}


	//defince structs
	category := Category{
		DocType:           	"Category",
		Id:		         	id,
		CategoryName:       category_name,
		Img1path: 			img1Path,
		Img2path:			img2Path,
		Img3path:			img3Path,
		VideoPath:			videoPath,
	}

	//convert Golang to jSon format (JSON Byte Array)
	categoryBytes, err = json.Marshal(category)
	if err != nil {
		return nil, err
	}

	//put category data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(id, categoryBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &category, nil
}

//Variety
type Variety struct {
	DocType					string		`json:"docType"`
	Id						string		`json:"varietyId"`
	VarietyName				string		`json:"varietyName"`
	Img1path				string		`json:"img1path"`
	Img2path				string		`json:"img2path"`
	Img3path				string		`json:"img3path"`
	VideoPath				string		`json:"videoPath"`

}
//Add new variety
func (spc *GoldContract) AddVariety(ctx contractapi.TransactionContextInterface, variety_name string, img1Path string, img2Path string, img3Path string, videoPath string) (*Variety, error) {

	id, _ := contract.IDGenerator("Variety", variety_name, count)
	varietyBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if varietyBytes != nil {
		return nil, fmt.Errorf("the variety already exists %s", variety_name)
	}


	//defince structs
	variety := Variety{
		DocType:           	"Variety",
		Id:		         	id,
		VarietyName:       	variety_name,
		Img1path: 			img1Path,
		Img2path:			img2Path,
		Img3path:			img3Path,
		VideoPath:			videoPath,
	}

	//convert Golang to jSon format (JSON Byte Array)
	varietyBytes, err = json.Marshal(variety)
	if err != nil {
		return nil, err
	}

	//put variety data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(id, varietyBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &variety, nil
}

//Plans : 
type Plan struct {
	DocType					string		`json:"docType"`
	PlanType				string		`json:"planType"`
	Id						string		`json:"planId"`
	Name					string		`json:"name"`
	CyclePeriod				string		`json:"cyclePeriod"`
	Duration				string		`json:"duration"`
	CalcId					string		`json:"calcId"`

}

//Add new Standard Plan
func (spc *GoldContract) Plan(ctx contractapi.TransactionContextInterface, planType string, name string, cyclePeriod string, duration string, calcId string) (*Plan, error) {

	id, _ := contract.IDGenerator("Plan", planType, count)
	planBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if planBytes != nil {
		return nil, fmt.Errorf("the plan already exists %s", name)
	}


	//defince structs
	plan := Plan{
		DocType:           	"Plan",
		Id:		         	id,
		PlanType:       	planType,
		Name: 				name,
		CyclePeriod:		cyclePeriod,
		Duration:			duration,
		CalcId:			calcId,
	}

	//convert Golang to jSon format (JSON Byte Array)
	planBytes, err = json.Marshal(plan)
	if err != nil {
		return nil, err
	}

	//put plan data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(id, planBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &plan, nil
}
//Flexi Plan:

type FlexiPlan struct {
	DocType					string		`json:"docType"`
	PlanTypeId				string		`json:"planTypeId"`
	Amount					string		`json:"amount"`
	BuySellPriceId			string		`json:"buySellPriceId"`
	AddressId				string		`json:"addressId"`
	Mode					string		`json:"mode"`
	CycleperiodID			string		`json:"cycleperiodID"`
	Duration				string		`json:"duration"`
	TransactionId			string		`json:"transactionId"`
	Gold 					string		`json:"gold"`

}


// Helper function
func (spc *GoldContract) IDGenerator(doctype string, name string, count int) (string, error) {

	docSubstring := doctype[0:3]
	nameSubString := name[0:3]

	s := []string{docSubstring, nameSubString, strconv.Itoa(count)}

	return strings.Join(s, ""), nil
}