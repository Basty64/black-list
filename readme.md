# Веб приложение для создания и администрирования чёрного списка пользователей

Используемый стэк: go 1.22 / postgresql / redis / html / css / java-script

создать миграцию
> goose postgres "host=localhost port=5433 user=admin password=admin dbname=black_list sslmode=disable" up 
> 

запуск контейнера с базой данных
> docker run --name black_list -p 5433:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=black_list -d postgres:14
> 

handlers:

authorization
    	POST /api/v0/registration - регистрация пользователя
     
    	GET /api/v0/authentication - аутентификация пользователя

admin role
    	GET /api/v0/list - список работников
     
	GET /api/v0/list/{id} - информация о работнике по айди
 
	POST /api/v0/list - добавить работника в список
 
	DELETE /api/v0/list/{id} - удалить работника из списка
 
	PATCH /api/v0/list/{id} - обновить работника в списке
 
	GET /api/v0/test-frontend - посмотреть демо фронта  

### ---updates soon---

