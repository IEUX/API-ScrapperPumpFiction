# API-ScrapperPumpFiction
A quick scrapper made in go and an API to serve the data powered by GinGonic

## How to launch the project

First create two directory ```/logs``` and ```/database``` to permit the code to write files in it.

Then launch the scrapper with the command ```go run scrapper.go```

You can now launch the API by run ```go run API/server.go```

**API Documentation**

```GET /stations?city="[Name of the city]"```

There you are 👍
