package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type pen struct {
	Id        string `json/"Id"`
	Model     string `json/"Model"`
	Color     string `json/"Color"`
	Length    string `json/"Length"`
	Width     string `json/"Width"`
	LineWidth string `json/"LineWidth"`
}

type penPrice struct {
	Id        string `json/"Id"`
	BuyPrice  string `json/"BuyPrice"`
	SellPrice string `json/"SellPrice"`
}

type PrivateChaincode struct {
}

// Init
func (t *PrivateChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("PrivateChaincode initialization successful")
	return shim.Success(nil)
}

// Invoke
func (t *PrivateChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Invoke started")

	fn, args := stub.GetFunctionAndParameters()
	switch fn {
	case "createPen":
		return t.createPen(stub, args)
	case "getPen":
		return t.getPen(stub, args)
	case "getPenPrice":
		return t.getPenPrice(stub, args)
	default:
		return shim.Error("Invalid invoke function name")
	}
}

// Create pen
func (t *PrivateChaincode) createPen(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 8 {
		return shim.Error("Expected 8 arguments(id, model, color, length, width, linewidth, buyprice, sellprice)")
	}

	id := args[0]
	model := args[1]
	color := args[2]
	length := args[3]
	width := args[4]
	lineWidth := args[5]
	buyPrice, err1 := strconv.ParseFloat(args[6], 32)
	sellPrice, err2 := strconv.ParseFloat(args[7], 32)

	if err1 != nil || err2 != nil {
		return shim.Error("Error parsing price values")
	}

	pen := &pen{id, model, color, length, width, lineWidth}
	penBytes, err3 := json.Marshal(pen)
	if err3 != nil {
		return shim.Error(err3.Error())
	}

	buyPriceString := fmt.Sprintf("%f", buyPrice)
	sellPriceString := fmt.Sprintf("%f", sellPrice)
	penPrice := &penPrice{id, buyPriceString, sellPriceString}
	penPriceBytes, err4 := json.Marshal(penPrice)
	if err4 != nil {
		return shim.Error(err4.Error())
	}

	err5 := stub.PutPrivateData("collectionData", id, penBytes)
	if err5 != nil {
		return shim.Error(err5.Error())
	}

	err6 := stub.PutPrivateData("collectionPrivate", id, penPriceBytes)
	if err6 != nil {
		return shim.Error(err6.Error())
	}

	jsonPen, err7 := json.Marshal(pen)
	if err7 != nil {
		return shim.Error(err7.Error())
	}

	return shim.Success(jsonPen)
}

func (t *PrivateChaincode) getPen(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	id := args[0]
	pen := pen{}

	penBytes, err1 := stub.GetPrivateData("collectionData", id)
	if err1 != nil {
		return shim.Error(err1.Error())
	}

	err2 := json.Unmarshal(penBytes, &pen)
	if err2 != nil {
		fmt.Println("Error unmarshalling object with id: " + id)
		return shim.Error(err2.Error())
	}

	jsonPen, err3 := json.Marshal(pen)
	if err3 != nil {
		return shim.Error(err3.Error())
	}

	return shim.Success(jsonPen)
}

func (t *PrivateChaincode) getPenPrice(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	id := args[0]
	penPrice := penPrice{}

	penPriceBytes, err1 := stub.GetPrivateData("collectionPrivate", id)
	if err1 != nil {
		return shim.Error(err1.Error())
	}

	err2 := json.Unmarshal(penPriceBytes, &penPrice)
	if err2 != nil {
		fmt.Println("Error unmarshalling object with id: " + id)
		return shim.Error(err2.Error())
	}

	jsonPenPrice, err3 := json.Marshal(penPrice)
	if err3 != nil {
		return shim.Error(err3.Error())
	}

	return shim.Success(jsonPenPrice)
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(PrivateChaincode)); err != nil {
		fmt.Printf("Error starting PrivateChaincode chaincode: %s", err)
	}
}
