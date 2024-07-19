package main

import "black-list/internal/app"

func main() {

	//Создание приложения
	App, error := app.New()
	if error != nil {
		panic(error)
	}

	//Запуск приложения
	error = App.Run()
	if error != nil {
		panic(error)
	}
}
