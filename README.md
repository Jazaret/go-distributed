# go-distributed

This is a distirbuted system of services written in Go that simulates, monitors and stores sensor data of a power plant.

Communication is done asynchrnously with a service discovery system allowing the sensors to alert the coordinator that the sensor is online and transmitting data.

The services use:

* RabbitMQ as the messaging system
* PostgreSQL for storage
* Go Gorilla toolkit for web sockets
* CanvasJS for html charts
