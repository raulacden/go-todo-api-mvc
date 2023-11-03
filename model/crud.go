package model

import (
	"context"
	"fmt"
)

func CreateTodo(name, todo string) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO todo(name,todo) values($1 , $2)", name, todo)
	if err != nil {
		fmt.Println("Error -", err)
		return err
	}
	return nil
}

func DeleteTodo(name string) error {
	_, err := conn.Exec(context.Background(), "DELETE FROM todo where name=$1", name)
	if err != nil {
		fmt.Println("Error -", err)
		return err
	}
	return nil
}
