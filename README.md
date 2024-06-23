# CarFuelEfficiencyCalculator

This AWS Lambda function calculates the fuel efficiency of a car and estimates the travel cost based on the distance traveled, fuel consumed, and fuel price.

## Features

- Calculate fuel efficiency in miles per gallon.
- Estimate the travel cost based on the fuel price and distance.

## Getting Started

### Prerequisites

- Go (version 1.20 or later)
- AWS CLI configured with appropriate permissions
- An AWS IAM role with permissions to execute Lambda functions

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/yourusername/CarFuelEfficiencyCalculator.git
   cd CarFuelEfficiencyCalculator

2. **Build the Go application:**
    export GOOS=linux
    export GOARCH=amd64
    go build -o bootstrap main.go

3. **Package the application:**
    zip function.zip bootstrap

4. **Deploy the Lambda function:**
    aws lambda create-function --function-name CarFuelEfficiencyCalculator \
    --zip-file fileb://function.zip \
    --handler bootstrap \
    --runtime provided.al2023 \
    --role arn:aws:iam::YOUR_IAM_ROLE_ARN

### Usage

1. **Create a test event (event.json):**

    {
    "distance": 300,
    "fuelConsumed": 10,
    "fuelPrice": 3.50
    }

2. **Invoke the Lambda function:**

    aws lambda invoke \
    --function-name CarFuelEfficiencyCalculator \
    --cli-binary-format raw-in-base64-out \
    --payload file://event.json \
    response.json

3. **Check the response:**

    {
    "fuelEfficiency": 30.0,
    "travelCost": 35.0,
    "message": "Your car's fuel efficiency is 30.00 miles per gallon, and the travel cost is $35.00"
    }




