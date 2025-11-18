package database

import (
	"backend/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhitespaceTrimming(t *testing.T) {
	var db = NewDatabase()

	var input = types.ToDo{
		Title: " t ",
		Description: " d ",
	}

	query, err := db.AddTodo(input)

	assert.Nil(t, err)

	var expectedOutput = types.ToDo{
	Title: "t",
	Description: "d",
	}

	assert.Equal(t, expectedOutput, query, "The two words should be the same.")

}

func TestEmptyFields(t *testing.T) {
	var db = NewDatabase()

	var input1 = types.ToDo{
		Title: "",
		Description: " d ",
	}
	var input2 = types.ToDo{
		Title: "d",
		Description: "",
	}

	_, err1 := db.AddTodo(input1)
	_, err2 := db.AddTodo(input2)

	assert.NotNil(t, err1)
	assert.NotNil(t, err2)
}

func TestAllTodosReturned(t *testing.T) {
	var db = NewDatabase()

	var input1 = types.ToDo{
		Title: "a",
		Description: "a",
	}

	var input2 = types.ToDo{
		Title: "aa",
		Description: "aa",
	}

	var input3 = types.ToDo{
		Title: "aaa",
		Description: "aaa",
	}

	var list = types.ToDoList{input1, input2, input3}

	for _, val := range list {
		db.AddTodo(val);
	}

	assert.ElementsMatch(t, db.GetAll(), list)	
}

func TestEmptyDBReturnsEmpty(t *testing.T) {
	var db = NewDatabase();

	todos := db.GetAll();
	expected := types.ToDoList{}

	assert.Equal(t, todos, expected)
}



	
