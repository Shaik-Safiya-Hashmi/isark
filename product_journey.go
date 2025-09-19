package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Product defines the structure for product journey
type Product struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Stage     string `json:"stage"`
	Location  string `json:"location"`
	Timestamp string `json:"timestamp"`
	Details   string `json:"details"`
}

// SmartContract provides functions for managing products
type SmartContract struct {
	contractapi.Contract
}

// AddProduct adds a new product entry
func (s *SmartContract) AddProduct(ctx contractapi.TransactionContextInterface, id string, name string, stage string, location string, timestamp string, details string) error {
	product := Product{
		ID:        id,
		Name:      name,
		Stage:     stage,
		Location:  location,
		Timestamp: timestamp,
		Details:   details,
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, productJSON)
}

// GetProduct retrieves a product by ID
func (s *SmartContract) GetProduct(ctx contractapi.TransactionContextInterface, id string) (*Product, error) {
	productJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if productJSON == nil {
		return nil, fmt.Errorf("product %s does not exist", id)
	}

	var product Product
	err = json.Unmarshal(productJSON, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// GetAllProducts retrieves all products
func (s *SmartContract) GetAllProducts(ctx contractapi.TransactionContextInterface) ([]*Product, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var products []*Product
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var product Product
		err = json.Unmarshal(queryResponse.Value, &product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		panic(fmt.Sprintf("Error creating product_journey chaincode: %v", err))
	}

	if err := chaincode.Start(); err != nil {
		panic(fmt.Sprintf("Error starting product_journey chaincode: %v", err))
	}
}

