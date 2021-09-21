package contracts

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	//"log"
	"strconv"
	"strings"
	"time"

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
	Id						string		`json:"Id"`
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
	Id						string		`json:"Id"`
	CollectionName			string		`json:"collectionName"`
	Img1path				string		`json:"img1path"`
	Img2path				string		`json:"img2path"`
	Img3path				string		`json:"img3path"`
	VideoPath				string		`json:"videoPath"`
	Status					string		`json:"status"`

}
//Add new collection
func (spc *GoldContract) AddCollection(ctx contractapi.TransactionContextInterface, collection_name string, img1Path string, img2Path string, img3Path string, videoPath string, status string) (*Collection, error) {

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
		Status:				status,
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
	Id						string		`json:"Id"`
	CategoryName			string		`json:"categoryName"`
	Img1path				string		`json:"img1path"`
	Img2path				string		`json:"img2path"`
	Img3path				string		`json:"img3path"`
	VideoPath				string		`json:"videoPath"`
	Status					string		`json:"status"`

}
//Add new categories
func (spc *GoldContract) AddCategory(ctx contractapi.TransactionContextInterface, category_name string, img1Path string, img2Path string, img3Path string, videoPath string, status string) (*Category, error) {

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
		Status:				status,
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
	Id						string		`json:"Id"`
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

//Offer
type Offer struct {
	DocType					string		`json:"docType"`
	Id						string		`json:"Id"`
	OfferName				string		`json:"offerName"`
	PercentageOf			string		`json:"percentageOf"`
	Status					string		`json:"status"`

}

//Needs to discuss offer and charges 
//Add new offer
func (spc *GoldContract) AddOffer(ctx contractapi.TransactionContextInterface, offerName string, percentageOf string, status string) (*Offer, error) {
	id, _ := contract.IDGenerator("Offer", offerName, count)
	offerBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if offerBytes != nil {
		return nil, fmt.Errorf("the offer already exists %s", offerName)
	}


	//defince structs
	offer := Offer{
		DocType:           	"Offer",
		Id:		         	id,
		OfferName:       	offerName,
		PercentageOf: 		percentageOf,
		Status:				status,
	}

	//convert Golang to jSon format (JSON Byte Array)
	offerBytes, err = json.Marshal(offer)
	if err != nil {
		return nil, err
	}

	//put offer data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(id, offerBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &offer, nil
}

//Diamond
type Diamond struct {
	DocType					string			`json:"docType"`
	Id						string			`json:"Id"`
	Shape					string			`json:"shape"`
	Gemstones				string			`json:"gemstones"`
	Clarity					string			`json:"clarity"`
	Color					string			`json:"color"`
	Cut 					string			`json:"cut"`
	CertifiyAuthority		string			`json:"certify_authority`
	Status					string			`json:"status"`
	Variety_id				string			`json:"variety_id"`
	Category_id				string			`json:"category_id"`
	Collection_id			string			`json:"collection_id"`
	Varieties				[]Variety		`json:"varieties"`
	Categories				[]Category		`json:"categories"`
	Collections 			[]Collection	`json:"collections"`
}
//Add new Diamond
func (spc *GoldContract) AddDiamond(ctx contractapi.TransactionContextInterface, shape string, gemstones string, clarity string, color string, cut string, certify_authority string, status string, variety_id string, category_id string, collection_id string) (*Diamond, error) {

	id, _ := contract.IDGenerator("diamond", shape, count)
	diamondBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if diamondBytes != nil {
		return nil, fmt.Errorf("the diamond already exists %s", shape)
	}
	//Get varity data
	queryStringVariety := fmt.Sprintf(`{"selector":{"docType":"%s","Id":"%s"}}`, "Variety", variety_id)
	
	varietyBytes, err := contract.GetQueryResult(ctx, queryStringVariety)
	if err != nil {
		return nil, fmt.Errorf("failed to read varity data from world state: %v", err)
	}

	varietyArray := []Variety{}
	err = json.Unmarshal([]byte(varietyBytes), &varietyArray)

	if err != nil {
		return nil, err
	}
	//Get category data
	queryStringCategory := fmt.Sprintf(`{"selector":{"docType":"%s","Id":"%s"}}`, "Category", category_id)
	
	categoryBytes, err := contract.GetQueryResult(ctx, queryStringCategory)
	if err != nil {
		return nil, fmt.Errorf("failed to read category data from world state: %v", err)
	}

	categoryArray := []Category{}
	err = json.Unmarshal([]byte(categoryBytes), &categoryArray)

	if err != nil {
		return nil, err
	}

	//Get collection data
	queryStringCollection := fmt.Sprintf(`{"selector":{"docType":"%s","Id":"%s"}}`, "Collection", collection_id)
	
	collectionBytes, err := contract.GetQueryResult(ctx, queryStringCollection)
	if err != nil {
		return nil, fmt.Errorf("failed to read collection data from world state: %v", err)
	}

	collectionArray := []Collection{}
	err = json.Unmarshal([]byte(collectionBytes), &collectionArray)

	if err != nil {
		return nil, err
	}

	//defince structs
	diamond := Diamond{
		DocType:           			"Diamond",
		Id:							id,
		Shape: 						shape,			
		Gemstones:         			gemstones,
		Clarity:					clarity,
		Color:						color,
		Cut:						cut,
		CertifiyAuthority:			certify_authority,
		Status:						"active",
		Variety_id:					variety_id,
		Collection_id:				collection_id,
		Category_id:				category_id,
		Varieties:					varietyArray,
		Categories:					categoryArray,
		Collections:				collectionArray,
	}

	//convert Golang to jSon format (JSON Byte Array)
	diamondBytes, err = json.Marshal(diamond)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(id, diamondBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &diamond, nil
}

//Calculation
type Calculation struct {
	DocType					string			`json:"docType"`
	Id						string			`json:"Id"`
	Sno 					int 			`json:"sno"`
	Type 					string			`json:"type"`
	Percentage				float64 		`json:"percentage"`
	Status					string			`json:"status"`
	CreatedDate				string			`json:"createdDate"`
	ModifiedDate			string			`json:"modifiedDate"`
}
//Add new calculation
func (spc *GoldContract) AddCalculation(ctx contractapi.TransactionContextInterface, sno int, Type string, percentage float64, status string) (*Calculation, error) {

	id, _ := contract.IDGenerator("calculation", Type, count)
	calculationBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if calculationBytes != nil {
		return nil, fmt.Errorf("the calculation already exists %s", Type)
	}

	var dtformated string
	dt := time.Now()
	dtformated = dt.Format("2006.01.02 15:04:05")

	//defince structs
	calculation := Calculation{
		DocType:           "Calculation",
		Id:		    		id,
		Sno:         		sno,
		Type:				Type,
		Percentage:			percentage,
		Status:				status,
		CreatedDate:		dtformated,
		ModifiedDate:		dtformated,
	}

	//convert Golang to jSon format (JSON Byte Array)
	calculationBytes, err = json.Marshal(calculation)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(id, calculationBytes)
	if err != nil {
		return nil, err
	}

	count += 1
	return &calculation, nil
}
//Plans : 
type Plan struct {
	DocType					string				`json:"docType"`
	PlanType				string				`json:"planType"`
	Id						string				`json:"planId"`
	Mode					string				`json:"mode"`
	Name					string				`json:"name"`
	CyclePeriod				[]CyclePeriod		`json:"cyclePeriod"`
	Duration				string				`json:"duration"`
	Bonus					string				`json:"bonus"`
	CalcId					[]Calculation		`json:"calcId"`

}

//Add Plan
func (spc *GoldContract) Plan(ctx contractapi.TransactionContextInterface, planType string, name string, duration string, bonus string, cyclePeriod string,  calcId string) (*Plan, error) {

	id, _ := contract.IDGenerator("plan", planType, count)
	planBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if planBytes != nil {
		return nil, fmt.Errorf("the plan already exists %s", name)
	}
	//Get calculation
	queryStringCal := fmt.Sprintf(`{"selector":{"docType":"%s","Id":"%s"}}`, "Calculation", calcId)
	
	calBytes, err := contract.GetQueryResult(ctx, queryStringCal)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	
	calArray := []Calculation{}
	err = json.Unmarshal([]byte(calBytes), &calArray)

	if err != nil {
		return nil, err
	}
	//Get Cycle Period
	queryStringCycle := fmt.Sprintf(`{"selector":{"docType":"%s","Id":"%s"}}`, "CyclePeriod", cyclePeriod)
	
	cycleBytes, err := contract.GetQueryResult(ctx, queryStringCycle)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	
	cycleArray := []CyclePeriod{}
	err = json.Unmarshal([]byte(cycleBytes), &cycleArray)

	if err != nil {
		return nil, err
	}
	//defince structs
	plan := Plan{
		DocType:           	"Plan",
		Mode:				"weight",
		Id:		         	id,
		PlanType:       	planType,
		Name: 				name,
		Duration:			duration,
		Bonus:				bonus,
		CyclePeriod:		cycleArray,
		CalcId:				calArray,
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


// Helper function
func (spc *GoldContract) IDGenerator(doctype string, name string, count int) (string, error) {

	docSubstring := doctype[0:4]
	nameSubString := name[0:4]

	s := []string{docSubstring, nameSubString, strconv.Itoa(count)}

	return strings.Join(s, ""), nil
}