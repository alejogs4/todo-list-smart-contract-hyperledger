package todolistsmartcontract

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type TodoListSmartContract struct {
	contractapi.Contract
}

func (tl *TodoListSmartContract) CreateTodoList(ctx contractapi.TransactionContextInterface, title, owner string) error {
	if len(title) == 0 || len(owner) == 0 {
		return ErrInvalidTodoInformation
	}

	timestamp, _ := ctx.GetStub().GetTxTimestamp()
	newTodo := CreateTodo(title, owner, timestamp)

	todoAsJSON, err := json.Marshal(newTodo)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(newTodo.ID, todoAsJSON)
}

func (tl *TodoListSmartContract) RemoveByID(ctx contractapi.TransactionContextInterface, id string) error {
	doesExists, ok := tl.exists(ctx, id)
	if !ok {
		return ErrGettingTodo
	}

	if !doesExists {
		return ErrNotFoundTodo
	}

	return ctx.GetStub().DelState(id)
}

func (tl *TodoListSmartContract) CompleteTodo(ctx contractapi.TransactionContextInterface, id string) error {
	todo, err := tl.GetByID(ctx, id)
	if err != nil {
		return err
	}

	completedTodo := todo.Clone()
	completedTodo.Completed = true
	newTodo, err := json.Marshal(completedTodo)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(completedTodo.ID, newTodo)
}

func (tl *TodoListSmartContract) ChangeOwner(ctx contractapi.TransactionContextInterface, oldOwner, newOwner, id string) error {
	todo, err := tl.GetByID(ctx, id)
	if err != nil {
		return err
	}

	targetTodo := todo.Clone()
	if targetTodo.Owner != oldOwner {
		return ErrInvalidOwner
	}

	targetTodo.Owner = newOwner
	newTodo, err := json.Marshal(targetTodo)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(targetTodo.ID, newTodo)
}

func (tl *TodoListSmartContract) GetAll(ctx contractapi.TransactionContextInterface) ([]*Todo, error) {
	todoListsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}

	defer todoListsIterator.Close()

	var todos []*Todo
	for todoListsIterator.HasNext() {
		response, err := todoListsIterator.Next()
		if err != nil {
			return nil, err
		}

		var todo Todo
		err = json.Unmarshal(response.Value, &todo)
		if err != nil {
			return nil, err
		}

		todos = append(todos, &todo)
	}
	return todos, err
}

func (tl *TodoListSmartContract) GetByID(ctx contractapi.TransactionContextInterface, id string) (*Todo, error) {
	todoAsJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, err
	}

	if todoAsJSON == nil {
		return nil, ErrNotFoundTodo
	}

	var todo Todo
	err = json.Unmarshal(todoAsJSON, &todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (tl *TodoListSmartContract) exists(ctx contractapi.TransactionContextInterface, id string) (bool, bool) {
	todoAsJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, false
	}

	if todoAsJSON == nil {
		return false, true
	}

	return true, true
}
