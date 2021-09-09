package contracts

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// GoldContract contract for handling writing and reading from the world state
type GoldContract struct {
	contractapi.Contract
}

// MetalGroup : The asset being tracked on the chain
type Metalgroup struct {
	DocType           		string 		`json:"docType"`
	MetalgroupID         	string 		`json:"metalgroupID"`
	MetalgroupName          string 		`json:"metalgroupname"`
	Karatage			    string 		`json:"karatage"`
	Fineness             	string		`json:"fineness"`
	ReferenceID            	string 		`json:"referenceID"`
	ShortName				string		`json:"shortname"`
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



// Init and Creator Functions.
func (spc *GoldContract) InitGold(ctx contractapi.TransactionContextInterface) error {

	return nil
}

