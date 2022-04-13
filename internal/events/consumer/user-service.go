package consumer

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal"
	"github.com/streadway/amqp"
)

const userQueue = "user-service"

type UserServiceEvents struct {
	*amqp.Channel
}

func NewUserServiceEvents(channel *amqp.Channel, user internal.User) (bool, error) {
	// Declare the queue
	_, err := channel.QueueDeclare(
		userQueue, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // argumentsR
	)
	if err != nil {
		panic(err)
	}

	// subscribe to the queue
	messages, err := channel.Consume(
		userQueue, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Println(err)
		return false, nil
	}

	//build a welcome message
	log.Println("User Service is ready to receive messages")

	// Make a channel to receive messages
	done := make(chan bool)

	go handleUserServiceEvents(messages, user)

	return <-done, nil
}

func handleUserServiceEvents(messages <-chan amqp.Delivery, user internal.User) {
	for message := range messages {

		log.Printf("Received a message: %s", message.Body)

		// populate the message
		KeycloakMessage := KeycloakMessage{}

		// serialize the message
		err := json.Unmarshal(message.Body, &KeycloakMessage)
		if err != nil {
			log.Println(err)
			return
		}
		if KeycloakMessage.OperationType == "CREATE" {
			CreateUser(KeycloakMessage.ResourcePath, user)
		}
		if KeycloakMessage.OperationType == "DELETE" {
			DeleteUser(KeycloakMessage.ResourcePath, user)
		}

	}
}

type KeycloakMessage struct {
	ResourcePath  string `json:"resourcePath"`
	OperationType string `json:"operationType"`
}

func CreateUser(resourcePath string, user internal.User) error {

	result := strings.Split(resourcePath, "/")
	keycloakId := result[1] //<- splits the id from the url   "users/f0576ce0-4335-4b83-883c-f26d6a6c4aac"

	// Create the user
	err := user.Create(context.Background(), keycloakId)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeleteUser(resourcePath string, user internal.User) error {

	result := strings.Split(resourcePath, "/")
	keycloakId := result[1] //<- splits the id from the url   "users/f0576ce0-4335-4b83-883c-f26d6a6c4aac"

	// Create the user
	err := user.Delete(context.Background(), keycloakId)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
