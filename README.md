# API_Call_Digesto
Simple example script for calling [Digesto API](https://blog.digesto.com.br/api-digesto-operações-adicione-o-poder-do-big-data-ao-seu-sistema-jur%C3%ADdico-a962f86a2520).

## Features
- Muiltithread Asyncronous request
- Deals with JSON pagination
- JSON returns as .csv

## How to run
- You must have a .csv file with a single collumn containg all parameters for requesting the API
- You must save it as ```requests.csv``` on ```data``` folder
- The returnded file will be save on the same ```data``` folder with the name: ```response.csv```
- Create a .env file with your authorization token naming it "AUTH"
```bash
AUTH = myApiKey

# Run
- go run main.go

## Note
Although this Caller is pre-set to a specific API, you just have to change:
- API constant on ```main.go```
- The ```models/response.go``` struct to match the specific return of the API

PS: Some of the API may work on ```GET``` endpoints and not on ```POST``` endpoints, if it is you case you nedd to change all functions on ```request``` folder
