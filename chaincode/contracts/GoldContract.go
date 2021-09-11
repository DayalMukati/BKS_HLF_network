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
	Id		 				string		`json:"metalId"`
	Name					string		`json:"name"`
	ImagePath				string		`json:"imagePath"`
}

// MetalGroup : The asset being tracked on the chain
type Metalgroup struct {
	DocType           		string 		`json:"docType"`
	Id			         	string 		`json:"metalgroupId"`
	Metals		            []string 	`json:"metals"`
	Karatage			    string 		`json:"karatage"`
	Fineness             	int			`json:"fineness"`
	ReferenceId            	int 		`json:"referenceId"`
	ShortName				string		`json:"shortName"`
}

//Product
type Product struct {
	DocType 				string		`json:"docType"`
	Id						string		`json:"productId"`
	Name					string		`json:"name"`
	ImagePath				string		`json:"imagePath"`
	VideoPath				string		`json:"videoPath"`
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


//Category
type Category struct {
	DocType					string		`json:"docType"`
	CategoryId				string		`json:"categoryId"`
	CategoryName			string		`json:"categoryName"`
	Img1path				string		`json:"img1path"`
	Img2path				string		`json:"img2path"`
	Img3path				string		`json:"img3path"`
	VideoPath				string		`json:"videoPath"`

}

//Variety
type Variety struct {
	DocType					string		`json:"docType"`
	VarietyId				string		`json:"varietyId"`
	VarietyName				string		`json:"varietyName"`
	Img1path				string		`json:"img1path"`
	Img2path				string		`json:"img2path"`
	Img3path				string		`json:"img3path"`
	VideoPath				string		`json:"videoPath"`

}
//StandardPlan : 
type StandardPlan struct {
	DocType					string		`json:"docType"`
	PlanTypeId				string		`json:"planTypeId"`
	PlanId					string		`json:"planId"`
	Amount					string		`json:"amount"`
	BuySellPriceId			string		`json:"buySellPriceId"`
	AddressId				string		`json:"addressId"`
	Mode					string		`json:"mode"`
	TransactionId			string		`json:"transactionId"`
	Gold 					string		`json:"gold"`
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

//CyclePeriod:

type CyclePeriod struct {
	DocType					string 		`json:"docType"`
	Name 					string		`json:"name"`
	Graceperiod				string		`json:"graceperiod"`
	MinWeight				string		`json:"minWeight"`
	MinValue				string		`json:"minValue"`
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
		Id:		         	id,
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
func (spc *GoldContract) AddMetalGroup(ctx contractapi.TransactionContextInterface, metals string, karatage string, fineness int, referenceId int, shortName string) (*Metalgroup, error) {

	id, _ := contract.IDGenerator("mgroup", shortName, count)
	metalgroupBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if metalgroupBytes != nil {
		return nil, fmt.Errorf("the metal already exists %s", shortName)
	}


	//defince structs
	metalGroup := Metalgroup{
		DocType:           		"MetalGroup",
		Id:			         	id,
		Metals: 				[]string{metals},			
		ShortName:         		shortName,
		Karatage:				karatage,
		Fineness:				fineness,
		ReferenceId:			referenceId,
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

//Add new collection
func (spc *GoldContract) AddCollection(ctx contractapi.TransactionContextInterface, collection_name string, image1Path string, image2Path string, image3Path string, videoPath string) (*Collection, error) {

	id, _ := contract.IDGenerator("collection", collection_name, count)
	collectionBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if collectionBytes != nil {
		return nil, fmt.Errorf("the product already exists %s", name)
	}


	//defince structs
	collection := Collection{
		DocType:           	"Collection",
		Id:		         	id,
		CollectionName:     collection_name,
		Img1path: 			img1path,
		Img2path:			img2path,
		Img3path:			img3path,
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


//Add new categories


//Add new variety


// Helper function
func (spc *GoldContract) IDGenerator(doctype string, name string, count int) (string, error) {

	docSubstring := doctype[0:3]
	nameSubString := name[0:3]

	s := []string{docSubstring, nameSubString, strconv.Itoa(count)}

	return strings.Join(s, ""), nil
}