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

// var dtformated string
var dt = time.Now()
var dtformated = dt.Format("2006.01.02 15:04:05")

//Metal
type Metal struct {
	DocType      string `json:"docType"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	Icon         string `json:"icon"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
}

// MetalGroup : The asset being tracked on the chain
type Metalgroup struct {
	DocType      string  `json:"docType"`
	Id           string  `json:"id"`
	Metals       []Metal `json:"metals"`
	MetalId      string  `json:"metalId"`
	Karatage     string  `json:"karatage"`
	Fineness     float64 `json:"fineness"`
	ReferenceId  int     `json:"referenceId"`
	ShortName    string  `json:"shortName"`
	Status       string  `json:"status"`
	CreatedDate  string  `json:"createdDate"`
	ModifiedDate string  `json:"modifiedDate"`
}

var contract GoldContract

// Init and Creator Functions.
func (spc *GoldContract) InitGold(ctx contractapi.TransactionContextInterface) error {
	return nil
}

//Add new Metal
func (spc *GoldContract) AddMetal(ctx contractapi.TransactionContextInterface, metalData string) (*Metal, error) {

	if len(metalData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(metalData)

	var metal Metal
	err := json.Unmarshal([]byte(metalData), &metal)

	metalBytes, err := ctx.GetStub().GetState(metal.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if metalBytes != nil {
		return nil, fmt.Errorf("the metal already exists %s", metal.Name)
	}

	metal.DocType = "Metal"
	metal.CreatedDate = dtformated
	metal.ModifiedDate = dtformated

	//convert Golang to jSon format (JSON Byte Array)
	metalBytes, err = json.Marshal(metal)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(metal.Id, metalBytes)
	if err != nil {
		return nil, err
	}

	return &metal, nil
}

//Add new Metal Group
func (spc *GoldContract) AddMetalGroup(ctx contractapi.TransactionContextInterface, metalgroupData string) (*Metalgroup, error) {

	if len(metalgroupData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(metalgroupData)

	var metalGroup Metalgroup
	err := json.Unmarshal([]byte(metalgroupData), &metalGroup)

	metalgroupBytes, err := ctx.GetStub().GetState(metalGroup.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if metalgroupBytes != nil {
		return nil, fmt.Errorf("the metal already exists %s", metalGroup.ShortName)
	}
	fmt.Println("second", metalGroup)
	queryString := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Metal", metalGroup.MetalId)

	metalBytes, err := contract.GetQueryResult(ctx, queryString)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	metalArray := []Metal{}
	err = json.Unmarshal([]byte(metalBytes), &metalArray)
	fmt.Println(metalArray)
	if err != nil {
		return nil, err
	}

	metalGroup.DocType = "MetalGroup"
	metalGroup.Status = "Active"
	metalGroup.Metals = metalArray
	metalGroup.CreatedDate = dtformated
	metalGroup.ModifiedDate = dtformated

	//convert Golang to jSon format (JSON Byte Array)
	metalgroupBytes, err = json.Marshal(metalGroup)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(metalGroup.Id, metalgroupBytes)
	if err != nil {
		return nil, err
	}

	return &metalGroup, nil
}

//CyclePeriod:
type CyclePeriod struct {
	DocType      string `json:"docType"`
	Name         string `json:"name"`
	Id           string `json:"id"`
	Graceperiod  int    `json:"graceperiod"`
	MinWeight    int    `json:"minWeight"`
	MinValue     int    `json:"minValue"`
	Status       string `json:"status"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
}

//Add Cycle Period
func (spc *GoldContract) AddCyclePeriod(ctx contractapi.TransactionContextInterface, cycleData string) (*CyclePeriod, error) {

	if len(cycleData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(cycleData)

	var cycleperiod CyclePeriod
	err := json.Unmarshal([]byte(cycleData), &cycleperiod)

	cycleBytes, err := ctx.GetStub().GetState(cycleperiod.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if cycleBytes != nil {
		return nil, fmt.Errorf("the cycle period already exists %s", cycleperiod.Name)
	}

	cycleperiod.DocType = "CyclePeriod"
	cycleperiod.Status = "Active"
	cycleperiod.CreatedDate = dtformated
	cycleperiod.ModifiedDate = dtformated

	//convert Golang to jSon format (JSON Byte Array)
	cycleBytes, err = json.Marshal(cycleperiod)
	if err != nil {
		return nil, err
	}

	//put cycle period data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(cycleperiod.Id, cycleBytes)
	if err != nil {
		return nil, err
	}

	return &cycleperiod, nil
}

//Product
type Product struct {
	DocType      string `json:"docType"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	ImagePath    string `json:"imagePath"`
	VideoPath    string `json:"videoPath"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
}

//Add new Product
func (spc *GoldContract) AddProduct(ctx contractapi.TransactionContextInterface, productData string) (*Product, error) {

	if len(productData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(productData)

	var product Product
	err := json.Unmarshal([]byte(productData), &product)

	productBytes, err := ctx.GetStub().GetState(product.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if productBytes != nil {
		return nil, fmt.Errorf("the product already exists %s", product.Name)
	}

	product.DocType = "Product"
	product.CreatedDate = dtformated
	product.ModifiedDate = dtformated

	//convert Golang to jSon format (JSON Byte Array)
	productBytes, err = json.Marshal(product)
	if err != nil {
		return nil, err
	}

	//put product data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(product.Id, productBytes)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

//Collection
type Collection struct {
	DocType        string `json:"docType"`
	Id             string `json:"id"`
	CollectionName string `json:"collection_name"`
	Img1           string `json:"img1"`
	Img2           string `json:"img2"`
	Img3           string `json:"img3"`
	Video          string `json:"video"`
	Status         string `json:"status"`
	CreatedDate    string `json:"createdDate"`
	ModifiedDate   string `json:"modifiedDate"`
}

//Add new collection
func (spc *GoldContract) AddCollection(ctx contractapi.TransactionContextInterface, collectionData string) (*Collection, error) {

	if len(collectionData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(collectionData)

	var collection Collection
	err := json.Unmarshal([]byte(collectionData), &collection)

	collectionBytes, err := ctx.GetStub().GetState(collection.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if collectionBytes != nil {
		return nil, fmt.Errorf("the collection already exists %s", collection.CollectionName)
	}

	collection.DocType = "Collection"
	collection.Status = "Active"
	collection.CreatedDate = dtformated
	collection.ModifiedDate = dtformated

	//convert Golang to jSon format (JSON Byte Array)
	collectionBytes, err = json.Marshal(collection)
	if err != nil {
		return nil, err
	}

	//put collection data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(collection.Id, collectionBytes)
	if err != nil {
		return nil, err
	}

	return &collection, nil
}

//Category
type Category struct {
	DocType      string `json:"docType"`
	Id           string `json:"id"`
	CategoryName string `json:"category_name"`
	Img1         string `json:"img1"`
	Img2         string `json:"img2"`
	Img3         string `json:"img3"`
	Video        string `json:"video"`
	Status       string `json:"status"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
}

//Add new categories
func (spc *GoldContract) AddCategory(ctx contractapi.TransactionContextInterface, categoryData string) (*Category, error) {

	if len(categoryData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(categoryData)

	var category Category
	err := json.Unmarshal([]byte(categoryData), &category)

	categoryBytes, err := ctx.GetStub().GetState(category.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if categoryBytes != nil {
		return nil, fmt.Errorf("the category already exists %s", category.CategoryName)
	}

	category.DocType = "Category"
	category.Status = "Active"
	category.CreatedDate = dtformated
	category.ModifiedDate = dtformated

	//convert Golang to jSon format (JSON Byte Array)
	categoryBytes, err = json.Marshal(category)
	if err != nil {
		return nil, err
	}

	//put category data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(category.Id, categoryBytes)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

//Variety
type Variety struct {
	DocType      string `json:"docType"`
	Id           string `json:"id"`
	VarietyName  string `json:"variety_name"`
	Img1         string `json:"img1"`
	Img2         string `json:"img2"`
	Img3         string `json:"img3"`
	Video        string `json:"videoPath"`
	Status       string `json:"status"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
}

//Add new variety
func (spc *GoldContract) AddVariety(ctx contractapi.TransactionContextInterface, varietyData string) (*Variety, error) {

	if len(varietyData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(varietyData)

	var variety Variety
	err := json.Unmarshal([]byte(varietyData), &variety)

	varietyBytes, err := ctx.GetStub().GetState(variety.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if varietyBytes != nil {
		return nil, fmt.Errorf("the variety already exists %s", variety.VarietyName)
	}

	variety.DocType = "Variety"
	variety.Status = "Active"
	variety.CreatedDate = dtformated
	variety.ModifiedDate = dtformated

	//convert Golang to jSon format (JSON Byte Array)
	varietyBytes, err = json.Marshal(variety)
	if err != nil {
		return nil, err
	}

	//put variety data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(variety.Id, varietyBytes)
	if err != nil {
		return nil, err
	}

	return &variety, nil
}

// Charges
type Charges struct {
	DocType            string `json:"docType"`
	Id                 string `json:"id"`
	PaymentId          string `json:"paymentId"`
	Mode               string `json:"mode"`
	Status             string `json:"status"`
	Amount             string `json:"amount"`
	Weight             string `json:"weight"`
	DeleveryAgent      []User `json:"deleveryAgent"`
	InstantGoldAppiled string `json:"instantGoldAppiled"`
	CreatedDate        string `json:"createdDate"`
	ModifiedDate       string `json:"modifiedDate"`
}

//Needs to discuss charges and charges
//Add new charges
func (spc *GoldContract) AddCharges(ctx contractapi.TransactionContextInterface, chargesData string) (*Charges, error) {
	if len(chargesData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(chargesData)

	var charges Charges
	err := json.Unmarshal([]byte(chargesData), &charges)

	chargesBytes, err := ctx.GetStub().GetState(charges.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if chargesBytes != nil {
		return nil, fmt.Errorf("the charges already exists %s", charges.Id)
	}
	queryString := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "User", charges.DeleveryAgent)

	agentBytes, err := contract.GetQueryResult(ctx, queryString)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	agentArray := []User{}
	err = json.Unmarshal([]byte(agentBytes), &agentArray)
	fmt.Println(agentArray)
	if err != nil {
		return nil, err
	}

	var dtformated string
	dt := time.Now()
	dtformated = dt.Format("2006.01.02 15:04:05")

	charges.DocType = "Charges"
	charges.DeleveryAgent = agentArray
	charges.CreatedDate = dtformated
	charges.ModifiedDate = dtformated

	//convert Golang to jSon format (JSON Byte Array)
	chargesBytes, err = json.Marshal(charges)
	if err != nil {
		return nil, err
	}

	//put charges data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(charges.Id, chargesBytes)
	if err != nil {
		return nil, err
	}

	return &charges, nil
}

//Diamond
type Diamond struct {
	DocType           string       `json:"docType"`
	Id                string       `json:"id"`
	Shape             string       `json:"shape"`
	Gemstones         string       `json:"gemstones"`
	Clarity           string       `json:"clarity"`
	Color             string       `json:"color"`
	Cut               string       `json:"cut"`
	CertifiyAuthority string       `json:"certify_authority`
	Status            string       `json:"status"`
	Variety_id        string       `json:"variety_id"`
	Category_id       string       `json:"category_id"`
	Collection_id     string       `json:"collection_id"`
	Varieties         []Variety    `json:"varieties"`
	Categories        []Category   `json:"categories"`
	Collections       []Collection `json:"collections"`
	CreatedDate       string       `json:"createdDate"`
	ModifiedDate      string       `json:"modifiedDate"`
}

//Add new Diamond
func (spc *GoldContract) AddDiamond(ctx contractapi.TransactionContextInterface, diamondData string) (*Diamond, error) {

	if len(diamondData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(diamondData)

	var diamond Diamond
	err := json.Unmarshal([]byte(diamondData), &diamond)

	diamondBytes, err := ctx.GetStub().GetState(diamond.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if diamondBytes != nil {
		return nil, fmt.Errorf("the diamond already exists %s", diamond.Shape)
	}
	//Get varity data
	queryStringVariety := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Variety", diamond.Variety_id)

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
	queryStringCategory := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Category", diamond.Category_id)

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
	queryStringCollection := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Collection", diamond.Collection_id)

	collectionBytes, err := contract.GetQueryResult(ctx, queryStringCollection)
	if err != nil {
		return nil, fmt.Errorf("failed to read collection data from world state: %v", err)
	}

	collectionArray := []Collection{}
	err = json.Unmarshal([]byte(collectionBytes), &collectionArray)

	if err != nil {
		return nil, err
	}

	diamond.DocType = "Diamond"
	diamond.Status = "Active"
	diamond.Varieties = varietyArray
	diamond.Categories = categoryArray
	diamond.Collections = collectionArray

	//convert Golang to jSon format (JSON Byte Array)
	diamondBytes, err = json.Marshal(diamond)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(diamond.Id, diamondBytes)
	if err != nil {
		return nil, err
	}

	return &diamond, nil
}

//Calculation
type Calculation struct {
	DocType      string  `json:"docType"`
	Id           string  `json:"id"`
	Sno          int     `json:"sno"`
	Type         string  `json:"type"`
	Percentage   float64 `json:"percentage"`
	Status       string  `json:"status"`
	CreatedDate  string  `json:"createdDate"`
	ModifiedDate string  `json:"modifiedDate"`
}

//Add new calculation
func (spc *GoldContract) AddCalculation(ctx contractapi.TransactionContextInterface, calculationData string) (*Calculation, error) {

	if len(calculationData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(calculationData)

	var calculation Calculation
	err := json.Unmarshal([]byte(calculationData), &calculation)

	calculationBytes, err := ctx.GetStub().GetState(calculation.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if calculationBytes != nil {
		return nil, fmt.Errorf("the calculation already exists %s", calculation.Type)
	}

	var dtformated string
	dt := time.Now()
	dtformated = dt.Format("2006.01.02 15:04:05")

	calculation.DocType = "Calculation"
	calculation.Status = "Active"
	calculation.CreatedDate = dtformated
	calculation.ModifiedDate = dtformated

	//convert Golang to jSon format (JSON Byte Array)
	calculationBytes, err = json.Marshal(calculation)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(calculation.Id, calculationBytes)
	if err != nil {
		return nil, err
	}

	return &calculation, nil
}

//Plans :
type Plan struct {
	DocType      string        `json:"docType"`
	PlanType     string        `json:"planType"`
	Id           string        `json:"id"`
	Mode         string        `json:"mode"`
	Name         string        `json:"name"`
	CyclePeriod  []CyclePeriod `json:"cyclePeriod"`
	Duration     string        `json:"duration"`
	Bonus        string        `json:"bonus"`
	CalcId       []Calculation `json:"calcId"`
	CreatedDate  string        `json:"createdDate"`
	ModifiedDate string        `json:"modifiedDate"`
}

//Add Plan
func (spc *GoldContract) Plan(ctx contractapi.TransactionContextInterface, planData string) (*Plan, error) {

	if len(planData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(planData)

	var plan Plan
	err := json.Unmarshal([]byte(planData), &plan)

	planBytes, err := ctx.GetStub().GetState(plan.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if planBytes != nil {
		return nil, fmt.Errorf("the plan already exists %s", plan.Name)
	}
	//Get calculation
	queryStringCal := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Calculation", plan.CalcId)

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
	queryStringCycle := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "CyclePeriod", plan.CyclePeriod)

	cycleBytes, err := contract.GetQueryResult(ctx, queryStringCycle)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	cycleArray := []CyclePeriod{}
	err = json.Unmarshal([]byte(cycleBytes), &cycleArray)

	if err != nil {
		return nil, err
	}

	plan.DocType = "Plan"
	plan.Mode = "weight"
	plan.CyclePeriod = cycleArray
	plan.CalcId = calArray

	//convert Golang to jSon format (JSON Byte Array)
	planBytes, err = json.Marshal(plan)
	if err != nil {
		return nil, err
	}

	//put plan data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(plan.Id, planBytes)
	if err != nil {
		return nil, err
	}

	return &plan, nil
}

//Buy/Sell
type BuySell struct {
	DocType string       `json:"docType"`
	Id      string       `json:"id"`
	KT24    Buysellprice `json:"kt24"`
	// KT22						struct {
	// 	Buy 				string			`json:"buy"`
	// 	Sell				string			`json:"sell"`
	// } 				`json:"kt22"`
	// KT18						struct {
	// 	Buy 				string			`json:"buy"`
	// 	Sell				string			`json:"sell"`
	// } 				`json:"kt18"`
	// KT14						struct {
	// 	Buy 				string			`json:"buy"`
	// 	Sell				string			`json:"sell"`
	// } 				`json:"kt14"`
	// KT10						struct {
	// 	Buy 				string			`json:"buy"`
	// 	Sell				string			`json:"sell"`
	// } 				`json:"kt10"`
}

type Buysellprice struct {
	Buy  string `json:"buy"`
	Sell string `json:"sell"`
}

//Create Buy/Sell
func (spc *GoldContract) AddBuySell(ctx contractapi.TransactionContextInterface, buysellData string) (*BuySell, error) {
	if len(buysellData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(buysellData)

	var buysell BuySell
	err := json.Unmarshal([]byte(buysellData), &buysell)

	buysellBytes, err := ctx.GetStub().GetState(buysell.Id)
	fmt.Println("buysellData", buysell)
	// if (buysell.Price24ktBuy == "" && buysell.Price24ktSell == ""){
	// 	return nil, fmt.Errorf("Please pass buy sell price")
	// }
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if buysellBytes != nil {
		return nil, fmt.Errorf("the plan already exists %s", buysell.Id)
	}

	buysell.DocType = "BuySell"
	buysell.KT24 = buysell.KT24
	// buysell.KT22 = []Buysellprice{}
	// buysell.KT18 = []Buysellprice{}
	// buysell.KT14 = []Buysellprice{}
	// buysell.KT10 = []Buysellprice{}
	fmt.Println("final buysell", buysell)
	//convert Golang to jSon format (JSON Byte Array)
	buysellBytes, err = json.Marshal(buysell)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(buysell.Id, buysellBytes)
	if err != nil {
		return nil, err
	}

	return &buysell, nil
}

//Video
type Video struct {
	DocType      string `json:"docType"`
	Id           string `json:"id"`
	Language     string `json:"language"`
	Category     string `json:"category"`
	Video        string `json:"video"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
}

//Add new Video
func (spc *GoldContract) AddVideo(ctx contractapi.TransactionContextInterface, videoData string) (*Video, error) {

	if len(videoData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(videoData)

	var video Video
	err := json.Unmarshal([]byte(videoData), &video)

	videoBytes, err := ctx.GetStub().GetState(video.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if videoBytes != nil {
		return nil, fmt.Errorf("the video already exists %s", video.Id)
	}

	video.DocType = "Video"

	//convert Golang to jSon format (JSON Byte Array)
	videoBytes, err = json.Marshal(video)
	if err != nil {
		return nil, err
	}

	//put video data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(video.Id, videoBytes)
	if err != nil {
		return nil, err
	}

	return &video, nil
}

//Subscription
type Subscription struct {
	DocType           string        `json:"docType"`
	Id                string        `json:"id"`
	User              string        `json:"user"`
	Address           []Address     `json:"address"`
	Installments      []Installment `json:"installment"`
	Plan              []Plan        `json:"plan"`
	Status            string        `json:"status"`
	UnpaidSkips       int           `json:"unpaidSkips"`
	SkipCount         int           `json:"skipCount"`
	UnpaidInvestments int           `json:"unpaidInvestments"`
	CreatedDate       string        `json:"createdDate"`
	ModifiedDate      string        `json:"modifiedDate"`
	Custom            string        `json:"custom"`
	CustomPlan        struct {
		Name        string        `json:"name"`
		Mode        string        `json:"mode"`
		PlanType    string        `json:"planType"`
		CyclePeriod []CyclePeriod `json:"cyclePeriod"`
		Duration    string        `json:"duration"`
	}
}

//Custom Plan

//Create Subscription
func (spc *GoldContract) AddSubscription(ctx contractapi.TransactionContextInterface, subscriptionData string) (*Subscription, error) {
	if len(subscriptionData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(subscriptionData)

	var subscription Subscription
	err := json.Unmarshal([]byte(subscriptionData), &subscription)

	subscriptionBytes, err := ctx.GetStub().GetState(subscription.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if subscriptionBytes != nil {
		return nil, fmt.Errorf("the Item already exists %s", subscription.Id)
	}
	fmt.Println(subscription)

	//Get Installment
	queryStringInstallment := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Installment", subscription.Installments)

	instaBytes, err := contract.GetQueryResult(ctx, queryStringInstallment)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	instaArray := []Installment{}
	err = json.Unmarshal([]byte(instaBytes), &instaArray)
	fmt.Println(instaArray)
	if err != nil {
		return nil, err
	}

	//Get Address
	queryStringAddress := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Address", subscription.Address)

	addressBytes, err := contract.GetQueryResult(ctx, queryStringAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	addressArray := []Address{}
	err = json.Unmarshal([]byte(addressBytes), &addressArray)
	fmt.Println(addressArray)
	if err != nil {
		return nil, err
	}

	//Get Plan -- for adding subscription
	if subscription.Plan != nil {
		queryStringPlan := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Plan", subscription.Plan)

		planBytes, err := contract.GetQueryResult(ctx, queryStringPlan)
		if err != nil {
			return nil, fmt.Errorf("failed to read data from world state: %v", err)
		}

		planArray := []Plan{}
		err = json.Unmarshal([]byte(planBytes), &planArray)
		fmt.Println(planArray)
		if err != nil {
			return nil, err
		}
		subscription.Plan = planArray
	}

	//
	//Add flexi subscription
	// customPlan := CustomPlan{}

	if len(subscription.Custom) > 0 {
		//Get CyclePeriod
		queryStringCycle := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "CyclePeriod", subscription.CustomPlan.CyclePeriod)

		cycleBytes, err := contract.GetQueryResult(ctx, queryStringCycle)
		if err != nil {
			return nil, fmt.Errorf("failed to read data from world state: %v", err)
		}

		cycleArray := []CyclePeriod{}
		err = json.Unmarshal([]byte(cycleBytes), &cycleArray)
		fmt.Println(cycleArray)
		if err != nil {
			return nil, err
		}
		subscription.CustomPlan.CyclePeriod = cycleArray
	}

	subscription.DocType = "Subscription"
	subscription.Installments = instaArray
	subscription.Address = addressArray

	//convert Golang to jSon format (JSON Byte Array)
	subscriptionBytes, err = json.Marshal(subscription)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(subscription.Id, subscriptionBytes)
	if err != nil {
		return nil, err
	}

	//Update subscription to user record
	queryStringUser := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "User", subscription.User)

	userBytes, err := contract.GetQueryResult(ctx, queryStringUser)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, fmt.Errorf("line 991: %v", err)
	}

	user.Subscriptions = append(user.Subscriptions, subscription)

	userBytes, err = json.Marshal(user)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(user.Id, userBytes)
	return &subscription, nil

}

//Installment
type Installment struct {
	DocType      string `json:"docType"`
	Id           string `json:"id"`
	PaymentId    string `json:"paymentId"`
	PaymentMode  string `json:"paymentmode"`
	StatusType   string `json:"statustype"`
	Amount       string `json:"amount"`
	Gold         string `json:"gold"`
	Bonus_Saved  string `json:"bonus_saved"`
	Userid       string `json:"user_id"`
	Collector    []User `json:"collector_id"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
}

//Create new Installment
func (spc *GoldContract) AddInstallment(ctx contractapi.TransactionContextInterface, installmentData string) (*Installment, error) {

	if len(installmentData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(installmentData)

	var installment Installment
	err := json.Unmarshal([]byte(installmentData), &installment)

	installmentBytes, err := ctx.GetStub().GetState(installment.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if installmentBytes != nil {
		return nil, fmt.Errorf("the Item already exists %s", installment.Id)
	}
	fmt.Println(installment)

	//Get Collector Agent
	queryStringAgent := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "User", installment.Collector)

	agentBytes, err := contract.GetQueryResult(ctx, queryStringAgent)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	agentArray := []User{}
	err = json.Unmarshal([]byte(agentBytes), &agentArray)
	fmt.Println(agentArray)
	if err != nil {
		return nil, err
	}

	installment.DocType = "Installment"
	installment.Collector = agentArray

	// bank.CreatedDate = dtformated
	// bank.ModifiedDate = dtformated

	installmentBytes, err = json.Marshal(installment)
	if err != nil {
		return nil, err
	}
	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(installment.Id, installmentBytes)
	if err != nil {
		return nil, err
	}

	return &installment, nil

}

//User
type User struct {
	DocType       string         `json:"docType"`
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Mobile        string         `json:"mobile"`
	DOB           string         `json:"dob"`
	Pan           string         `json:"pan"`
	IsWhatsapp    bool           `json:"isWhatsapp"`
	IsInvested    bool           `json:"isInvested"`
	Image         string         `json:"image"`
	Referral      []User         `json:"referral"`
	RefCode       string         `json:"refCode"`
	GBPCode       string         `json:"GBPcode"`
	Level         []Level        `json:"level"`
	Bank          []Bank         `json:"bank"`
	Address       []Address      `json:"address"`
	Subscriptions []Subscription `json:"subscription"`
	CreatedDate   string         `json:"createdDate"`
	ModifiedDate  string         `json:"modifiedDate"`
}

//Add new User
func (spc *GoldContract) AddUser(ctx contractapi.TransactionContextInterface, userData string) (*User, error) {
	if len(userData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(userData)

	var user User
	err := json.Unmarshal([]byte(userData), &user)

	userBytes, err := ctx.GetStub().GetState(user.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if userBytes != nil {
		return nil, fmt.Errorf("the metal already exists %s", user.Name)
	}

	user.DocType = "User"

	//convert Golang to jSon format (JSON Byte Array)
	userBytes, err = json.Marshal(user)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(user.Id, userBytes)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

//Role
type Role struct {
	DocType     string       `json:"docType"`
	Id          string       `json:"id"`
	Role_Name   string       `json:"role_name"`
	Permissions []Permission `json:"permissions"`
	Status      string       `json:"status"`
}

//Add new Role
func (spc *GoldContract) AddRole(ctx contractapi.TransactionContextInterface, roleData string) (*Role, error) {

	if len(roleData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(roleData)

	var role Role
	err := json.Unmarshal([]byte(roleData), &role)

	roleBytes, err := ctx.GetStub().GetState(role.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if roleBytes != nil {
		return nil, fmt.Errorf("the metal already exists %s", role.Role_Name)
	}
	fmt.Println("second", role)
	queryString := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Permission", role.Permissions)

	permissionBytes, err := contract.GetQueryResult(ctx, queryString)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	permissionArray := []Permission{}
	err = json.Unmarshal([]byte(permissionBytes), &permissionArray)
	fmt.Println(permissionArray)
	if err != nil {
		return nil, err
	}

	role.DocType = "Role"
	role.Status = "Active"
	role.Permissions = permissionArray

	//convert Golang to jSon format (JSON Byte Array)
	roleBytes, err = json.Marshal(role)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(role.Id, roleBytes)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

//Permission
type Permission struct {
	DocType         string `json:"docType"`
	Id              string `json:"id"`
	Permission_Name string `json:"permission_name"`
	Status          string `json:"status"`
}

//Add new Permission
func (spc *GoldContract) AddPermission(ctx contractapi.TransactionContextInterface, permissionData string) (*Permission, error) {

	if len(permissionData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(permissionData)

	var permission Permission
	err := json.Unmarshal([]byte(permissionData), &permission)

	permissionBytes, err := ctx.GetStub().GetState(permission.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if permissionBytes != nil {
		return nil, fmt.Errorf("the permission already exists %s", permission.Permission_Name)
	}

	permission.DocType = "Permission"
	permission.Status = "active"

	//convert Golang to jSon format (JSON Byte Array)
	permissionBytes, err = json.Marshal(permission)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(permission.Id, permissionBytes)
	if err != nil {
		return nil, err
	}

	return &permission, nil
}

//Level
type Level struct {
	DocType    string  `json:"docType"`
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Commission float64 `json:"commission"`
}

//Add new Label
func (spc *GoldContract) AddLevel(ctx contractapi.TransactionContextInterface, levelData string) (*Level, error) {

	if len(levelData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(levelData)

	var level Level
	err := json.Unmarshal([]byte(levelData), &level)

	levelBytes, err := ctx.GetStub().GetState(level.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if levelBytes != nil {
		return nil, fmt.Errorf("the level already exists %s", level.Name)
	}

	level.DocType = "Level"

	//convert Golang to jSon format (JSON Byte Array)
	levelBytes, err = json.Marshal(level)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(level.Id, levelBytes)
	if err != nil {
		return nil, err
	}

	return &level, nil
}

//Address
type Address struct {
	DocType          string `json:"docType"`
	Id               string `json:"id"`
	User             []User `json:"user"`
	Pin              string `json:"pin"`
	Landmark         string `json:"landmark"`
	IsDefaultAddress string `json:"isDefaultAddress"`
	Status           string `json:"status"`
}

//Add new Address
func (spc *GoldContract) AddAddress(ctx contractapi.TransactionContextInterface, addressData string) (*Address, error) {

	if len(addressData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(addressData)

	var address Address
	err := json.Unmarshal([]byte(addressData), &address)

	addressBytes, err := ctx.GetStub().GetState(address.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if addressBytes != nil {
		return nil, fmt.Errorf("the Item already exists %s", address.Id)
	}
	fmt.Println(address)

	address.DocType = "Address"
	address.Status = "active"

	// bank.CreatedDate = dtformated
	// bank.ModifiedDate = dtformated

	addressBytes, err = json.Marshal(address)
	if err != nil {
		return nil, err
	}
	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(address.Id, addressBytes)
	if err != nil {
		return nil, err
	}

	queryStringUser := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "User", address.User)

	userBytes, err := contract.GetQueryResult(ctx, queryStringUser)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, fmt.Errorf("line 909: %v", err)
	}

	user.Address = append(user.Address, address)

	userBytes, err = json.Marshal(user)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(user.Id, userBytes)

	return &address, nil

}

//Bank Details
type Bank struct {
	DocType      string `json:"docType"`
	Id           string `json:"id"`
	User         []User `json:"userId"`
	AccountNum   string `json:"Accountnum"`
	IFSC         string `json:"IFSC"`
	Bank         string `json:"Bank"`
	Branch       string `json:"Branch"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
}

//Add Bank Details
func (spc *GoldContract) AddBank(ctx contractapi.TransactionContextInterface, bankData string) (*Bank, error) {

	if len(bankData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(bankData)

	var bank Bank
	err := json.Unmarshal([]byte(bankData), &bank)

	bankBytes, err := ctx.GetStub().GetState(bank.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if bankBytes != nil {
		return nil, fmt.Errorf("the Item already exists %s", bank.Id)
	}
	fmt.Println(bank)

	bank.DocType = "Bank"

	// bank.CreatedDate = dtformated
	// bank.ModifiedDate = dtformated

	bankBytes, err = json.Marshal(bank)
	if err != nil {
		return nil, err
	}
	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(bank.Id, bankBytes)
	if err != nil {
		return nil, err
	}

	queryStringUser := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "User", bank.User)

	userBytes, err := contract.GetQueryResult(ctx, queryStringUser)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, fmt.Errorf("line 909: %v", err)
	}

	user.Bank = append(user.Bank, bank)

	userBytes, err = json.Marshal(user)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(user.Id, userBytes)

	return &bank, nil

}

//Ecom Transaction
type Transaction struct {
	DocType            string         `json:"docType"`
	Id                 string         `json:"id"`
	PaymentId          string         `json:"paymentId"`
	Amount             string         `json:"amount"`
	Mode               string         `json:"mode"`
	Status             string         `json:"status"`
	InstantGoldAppiled []Installment  `json:"instantGoldApplied"`
	Charges            []Charges      `json:"charges"`
	DeleveryAgent      []User         `json:"deleveryAgent"`
	Subscription       []Subscription `json:"subscription"`
	CreatedDate        string         `json:"createdDate"`
	ModifiedDate       string         `json:"modifiedDate"`
}

//Add Ecom Transaction
func (spc *GoldContract) AddTransaction(ctx contractapi.TransactionContextInterface, transactionData string) (*Transaction, error) {

	if len(transactionData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(transactionData)

	var transaction Transaction
	err := json.Unmarshal([]byte(transactionData), &transaction)

	transactionBytes, err := ctx.GetStub().GetState(transaction.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if transactionBytes != nil {
		return nil, fmt.Errorf("the Item already exists %s", transaction.Id)
	}
	fmt.Println(transaction)

	//Get Installment
	queryStringInstallment := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Installment", transaction.InstantGoldAppiled)

	instaBytes, err := contract.GetQueryResult(ctx, queryStringInstallment)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	instaArray := []Installment{}
	err = json.Unmarshal([]byte(instaBytes), &instaArray)
	fmt.Println(instaArray)
	if err != nil {
		return nil, err
	}

	//Get Charges
	queryStringCharges := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Charges", transaction.Charges)

	chargesBytes, err := contract.GetQueryResult(ctx, queryStringCharges)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	chargesArray := []Charges{}
	err = json.Unmarshal([]byte(chargesBytes), &chargesArray)
	fmt.Println(chargesArray)
	if err != nil {
		return nil, err
	}

	//Get Delivery Agent
	queryStringAgent := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "User", transaction.DeleveryAgent)

	agentBytes, err := contract.GetQueryResult(ctx, queryStringAgent)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	agentArray := []User{}
	err = json.Unmarshal([]byte(agentBytes), &agentArray)
	fmt.Println(agentArray)
	if err != nil {
		return nil, err
	}

	//Get Subscription
	queryStringSubscription := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Subscription", transaction.Subscription)

	subscriptionBytes, err := contract.GetQueryResult(ctx, queryStringSubscription)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	subscriptionArray := []Subscription{}
	err = json.Unmarshal([]byte(subscriptionBytes), &subscriptionArray)
	fmt.Println(subscriptionArray)
	if err != nil {
		return nil, err
	}

	transaction.DocType = "ItemDetail"
	transaction.InstantGoldAppiled = instaArray
	transaction.DeleveryAgent = agentArray
	transaction.Subscription = subscriptionArray
	transaction.Charges = chargesArray

	//convert Golang to jSon format (JSON Byte Array)
	transactionBytes, err = json.Marshal(transaction)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(transaction.Id, transactionBytes)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

//Appointment
type Appointment struct {
	DocType         string       `json:"docType"`
	Id              string       `json:"id"`
	MetalGroup      []Metalgroup `json:"metalGroup"`
	Buysellprice    []BuySell    `json:"buySellPrice"`
	Verifier        []User       `json:"verifier"`
	Weight          float64      `json:"weight"`
	Otp             string       `json:"otp"`
	AppointmentDate string       `json:"appointmentDate"`
	AppointmentTime string       `json:"appointmentTime"`
	Status          string       `json:"status"`
	StoreLocation   string       `json:"storeLocation"`
	CreatedDate     string       `json:"createdDate"`
	ModifiedDate    string       `json:"modifiedDate"`
}

//Add Appointment
func (spc *GoldContract) AddAppointment(ctx contractapi.TransactionContextInterface, appointmentData string) (*Appointment, error) {

	if len(appointmentData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(appointmentData)

	var appointment Appointment
	err := json.Unmarshal([]byte(appointmentData), &appointment)

	appointmentBytes, err := ctx.GetStub().GetState(appointment.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if appointmentBytes != nil {
		return nil, fmt.Errorf("the Item already exists %s", appointment.Id)
	}
	fmt.Println(appointment)

	//Get Metalgroup
	queryStringMgroup := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "MetalGroup", appointment.MetalGroup)

	mgroupBytes, err := contract.GetQueryResult(ctx, queryStringMgroup)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	mgroupArray := []Metalgroup{}
	err = json.Unmarshal([]byte(mgroupBytes), &mgroupArray)
	fmt.Println(mgroupArray)
	if err != nil {
		return nil, err
	}

	//Get Buy/Sell price
	queryStringBuysell := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "BuySell", appointment.Buysellprice)

	buysellBytes, err := contract.GetQueryResult(ctx, queryStringBuysell)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	buysellArray := []BuySell{}
	err = json.Unmarshal([]byte(buysellBytes), &buysellArray)
	fmt.Println(buysellArray)
	if err != nil {
		return nil, err
	}

	//Get Verifier
	queryStringVerifier := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "User", appointment.Verifier)

	verifierBytes, err := contract.GetQueryResult(ctx, queryStringVerifier)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	verifierArray := []User{}
	err = json.Unmarshal([]byte(verifierBytes), &verifierArray)
	fmt.Println(verifierArray)
	if err != nil {
		return nil, err
	}

	appointment.DocType = "Appointment"
	appointment.Verifier = verifierArray
	appointment.Buysellprice = buysellArray
	appointment.MetalGroup = mgroupArray

	//convert Golang to jSon format (JSON Byte Array)
	appointmentBytes, err = json.Marshal(appointment)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(appointment.Id, appointmentBytes)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

//Item
type Item struct {
	DocType      string `json:"docType"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	Images       string `json:"images"`
	Video        string `json:"video"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
}

//Add new Item
func (spc *GoldContract) AddItem(ctx contractapi.TransactionContextInterface, itemData string) (*Item, error) {

	if len(itemData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(itemData)

	var item Item
	err := json.Unmarshal([]byte(itemData), &item)

	itemBytes, err := ctx.GetStub().GetState(item.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if itemBytes != nil {
		return nil, fmt.Errorf("the item already exists %s", item.Name)
	}

	item.DocType = "Item"

	//convert Golang to jSon format (JSON Byte Array)
	itemBytes, err = json.Marshal(item)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(item.Id, itemBytes)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

//ItemDetails
type ItemDetail struct {
	DocType      string       `json:"docType"`
	Id           string       `json:"id"`
	Items        []Item       `json:"item"`
	Products     []Product    `json:"product"`
	Collections  []Collection `json:"collection"`
	Categories   []Category   `json:"category"`
	CreatedDate  string       `json:"createdDate"`
	ModifiedDate string       `json:"modifiedDate"`
}

//Add new ItemDetails
func (spc *GoldContract) AddItemDetails(ctx contractapi.TransactionContextInterface, itemdetailData string) (*ItemDetail, error) {

	if len(itemdetailData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(itemdetailData)

	var itemdetail ItemDetail
	err := json.Unmarshal([]byte(itemdetailData), &itemdetail)

	itemdetailBytes, err := ctx.GetStub().GetState(itemdetail.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}
	//check if ID already exists (return the state of the ID by checking the world state)
	if itemdetailBytes != nil {
		return nil, fmt.Errorf("the Item already exists %s", itemdetail.Id)
	}
	fmt.Println(itemdetail)

	//Get Items
	queryStringItem := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Item", itemdetail.Items)

	itemBytes, err := contract.GetQueryResult(ctx, queryStringItem)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	itemArray := []Item{}
	err = json.Unmarshal([]byte(itemBytes), &itemArray)
	fmt.Println(itemArray)
	if err != nil {
		return nil, err
	}

	//Get Products
	queryStringProduct := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Product", itemdetail.Products)

	productBytes, err := contract.GetQueryResult(ctx, queryStringProduct)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	productArray := []Product{}
	err = json.Unmarshal([]byte(productBytes), &productArray)
	fmt.Println(productArray)
	if err != nil {
		return nil, err
	}

	//Get Collections
	queryStringCollection := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Collection", itemdetail.Collections)

	collectionBytes, err := contract.GetQueryResult(ctx, queryStringCollection)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	collectionArray := []Collection{}
	err = json.Unmarshal([]byte(collectionBytes), &collectionArray)
	fmt.Println(collectionArray)
	if err != nil {
		return nil, err
	}

	//Get Category
	queryStringCategory := fmt.Sprintf(`{"selector":{"docType":"%s","id":"%s"}}`, "Category", itemdetail.Categories)

	categoryBytes, err := contract.GetQueryResult(ctx, queryStringCategory)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	categoryArray := []Category{}
	err = json.Unmarshal([]byte(categoryBytes), &categoryArray)
	fmt.Println(categoryArray)
	if err != nil {
		return nil, err
	}

	itemdetail.DocType = "ItemDetail"
	itemdetail.Items = itemArray
	itemdetail.Products = productArray
	itemdetail.Collections = collectionArray
	itemdetail.Categories = categoryArray

	//convert Golang to jSon format (JSON Byte Array)
	itemdetailBytes, err = json.Marshal(itemdetail)
	if err != nil {
		return nil, err
	}

	//put account data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(itemdetail.Id, itemdetailBytes)
	if err != nil {
		return nil, err
	}

	return &itemdetail, nil
}

//Cart
type Cart struct {
	DocType      string   `json:"docType"`
	Id           string   `json:"id"`
	Items        []string `json:"items"`
	Userid       string   `json:"userid"`
	CreatedDate  string   `json:"createdDate"`
	ModifiedDate string   `json:"modifiedDate"`
}

//Add Item to Cart
func (spc *GoldContract) AddTocart(ctx contractapi.TransactionContextInterface, cartData string) (*Cart, error) {

	if len(cartData) == 0 {
		return nil, fmt.Errorf("Please pass the correct data")
	}
	fmt.Println(cartData)

	var cart Cart
	err := json.Unmarshal([]byte(cartData), &cart)

	cartBytes, err := ctx.GetStub().GetState(cart.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from world state: %v", err)
	}

	// //update cart logic
	// if cartBytes != nil {

	// }

	cart.DocType = "Cart"

	//convert Golang to jSon format (JSON Byte Array)
	cartBytes, err = json.Marshal(cart)
	if err != nil {
		return nil, err
	}

	//put video data unto the Ledger (key value pair)
	err = ctx.GetStub().PutState(cart.Id, cartBytes)
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

// Helper function
func (spc *GoldContract) IDGenerator(doctype string, name string, count int) (string, error) {

	docSubstring := doctype[0:4]
	nameSubstring := name[0:3]

	s := []string{docSubstring, nameSubstring, strconv.Itoa(count)}

	return strings.Join(s, ""), nil
}
