package contracts

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"github.com/google/uuid"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// GoldContract contract for handling writing and reading from the world state
type GoldContract struct {
	contractapi.Contract
}

//Metal
type Metal struct {
	DocType					string		`json:"docType"`
	MetalId 				string		`json:"metalId"`
	Name					string		`json:"name"`
	ImagePath				string		`json:"imagePath"`
}

// MetalGroup : The asset being tracked on the chain
type Metalgroup struct {
	DocType           		string 		`json:"docType"`
	MetalgroupID         	string 		`json:"metalgroupId"`
	Metals		            []Metal 	`json:"metals"`
	Karatage			    string 		`json:"karatage"`
	Fineness             	string		`json:"fineness"`
	ReferenceID            	string 		`json:"referenceId"`
	ShortName				string		`json:"shortName"`
}

//Product
type Product struct {
	DocType 				string		`json:"docType"`
	ProductId				string		`json:"productId"`
	Name					string		`json:"name"`
	ImagePath				string		`json:"imagePath"`
	VideoPath				string		`json:"videoPath"`
}

//Collection
type Collection struct {
	DocType					string		`json:"docType"`
	CollectionId			string		`json:"collectionId"`
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

var metalCount int 
var metalGroupCount int

// Init and Creator Functions.
func (spc *GoldContract) InitGold(ctx contractapi.TransactionContextInterface) error {

	return nil
}

//Add new Metal 
func (spc *GoldContract) AddMetal(ctx contractapi.TransactionContextInterface, name string, imagePath string) (*Metal, error) {

	id, _ := uuid.New()
	metalBytes. err := ctx.GetStunb().GetState(id)

	if err != nil {
		return nil, fmt.Error("failed to read data from world state: %v", err)
	}

}
