package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"todo-api-mvc/controller"
	"todo-api-mvc/model"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	fmt.Println("Serving...")
	//Esto defiere la acción al final del programa
	defer db.Close(context.TODO())
	//Si hay error lo muestra, si no cierra la conexión
	log.Fatal(http.ListenAndServe(":3000", mux))
}
