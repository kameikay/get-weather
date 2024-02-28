# Get temperature by zip code

## Objective
The goal of this project is to develop a system in Go that takes a postal code as input, identifies the city, and returns the current weather (temperature in Celsius, Fahrenheit, and Kelvin). This system will be deployed on Google Cloud Run.

## Requirements

### Functional

- The system must accept a valid 8-digit zip code.
- After receiving the zip code, the system must identify the corresponding location and return temperatures formatted in Celsius, Fahrenheit and Kelvin.
- The system must respond to the following scenarios:
-- If successful: Respond with HTTP code 200 and a response body containing the temperatures in the three scales.
-- In case of invalid zip code (incorrect format): Respond with HTTP code 422 and an "invalid zip code" error message.
-- In case of zip code not found: Respond with HTTP code 404 and an error message below "can not find zipcode".

### Non-Functional
- The system deployment must be carried out on Google Cloud Run.
- Using the viaCEP API to find the desired location in the provided zip code.
- Use of the WeatherAPI API to query desired temperatures.
- Application of formulas for temperature conversion between Celsius, Fahrenheit and Kelvin scales.


## Technologies and Tools Used

- Programming Language: Go
- External APIs: viaCEP and WeatherAPI
- Hosting: Google Cloud Run
- Containerization: Docker


## Prerequisites

- Go 1.16+ installed
- Docker installed (for running the project via Docker)
- An internet connection to access external APIs (viaCEP and WeatherAPI)


## How to run the project

### Running the project locally

1. Clone the repository
2. Run the following command to start the application:
```bash
go run main.go
```
3. Access the following URL in your browser or using an API client (such as Postman):
```bash
http://localhost:8080/?cep={zipCode}
```
Replace `{zipCode}` with the desired zip code.

### Running the project via Docker

1. Run the following command to build the Docker image:
```bash
docker build -t get-temperature-by-zip-code .
```
2. Run the following command to start the application:
```bash
docker run -p 8080:8080 get-temperature-by-zip-code
```
3. Access the following URL in your browser or using an API client (such as Postman):
```bash
http://localhost:8080/?cep={zipCode}
```
Replace `{zipCode}` with the desired zip code.

### Running the tests

1. Run the following command to start the tests:
```bash
go test ./...
```

## Environment Variables

- `WEATHER_API_KEY`: API key for the [WeatherAPI](https://www.weatherapi.com/) service.

