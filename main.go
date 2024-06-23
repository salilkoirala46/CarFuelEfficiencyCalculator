package main 

import (
	"context"
	"fmt"
	"log"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Distance   float64 `json:"distance"`
	FuelConsumed float64 `json:"fuelConsumed"`
	FuelPrice  float64 `json:"fuelPrice"`
}

type MyResponse struct {
	FuelEfficiency float64 `json:"fuelEfficiency"`
	TravelCost     float64 `json:"travelCost"`
	Message        string  `json:"message"`
}

func HandleLambdaEvent(ctx context.Context, event MyEvent) (MyResponse, error) {
	log.Println("Received event:", event)

	if event.Distance <= 0 || event.FuelConsumed <= 0 || event.FuelPrice <= 0 {
		return MyResponse{}, fmt.Errorf("invalid input: distance, fuel consumed, and fuel price must be positive values")
	}

	fuelEfficiency := event.Distance / event.FuelConsumed
	travelCost := (event.Distance / fuelEfficiency) * event.FuelPrice
	message := fmt.Sprintf("Your car's fuel efficiency is %.2f miles per gallon, and the travel cost is $%.2f", fuelEfficiency, travelCost)

	response := MyResponse{
		FuelEfficiency: fuelEfficiency,
		TravelCost:     travelCost,
		Message:        message,
	}

	log.Println("Response:", response)
	return response, nil
}


func main () {
	lambda.Start(HandleLambdaEvent)
}