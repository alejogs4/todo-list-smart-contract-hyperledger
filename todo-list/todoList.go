package main

import (
	"log"

	"github.com/alejogs4/todo-list/todolistsmartcontract"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	todoListChainCode, err := contractapi.NewChaincode(&todolistsmartcontract.TodoListSmartContract{})
	if err != nil {
		log.Panicf("Error creating todo-list chaincode: %v", err)
	}

	if err := todoListChainCode.Start(); err != nil {
		log.Panicf("Error starting todo-list chaincode: %v", err)
	}
}
