package infrastructure

import "api-hexagonal-cars/src/cars/infrastructure/adapters"
 
var (
	mysql *MySQL
	rabbitmq *adapters.Rabbit
)

func GoMySQL() {
	mysql = NewMySQL()
}

func GetMySQL() *MySQL {
	return mysql 
}

func GoRabbitMQ(){
	rabbitmq = adapters.NewRabbitMq()
}

func GetRabbitMQ() *adapters.Rabbit {
	return rabbitmq
}