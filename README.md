# Eligibility

A check system that validates that the client can be supported within the platform

## How it works

Sending a dummy transaction in json format for application at http://localhost:8080/api/eligibility.

### Schema

#### Input

```json
{
  "numeroDoDocumento": "14041737706",
  "tipoDeConexao": "bifasico",
  "classeDeConsumo": "comercial",
  "modalidadeTarifaria": "convencional",
  "historicoDeConsumo": [
    3878,
    9760,
    5976,
    2797,
    2481,
    5731,
    7538,
    4392,
    7859,
    4160,
    6941,
    4597 
  ]
}
```

## You will receive output similar to this.

### Output

```json
{
   "elegivel": true,
   "economiaAnualDeCO2": 5553.24,
}
```

## Prerequisites

To run the application locally and run the tests it is necessary to install golang.

Run at the root of the project.

OS X & Linux:
```
go run .
```

## Running Unittest

Run at the root of the project.


Run unittest fremework in OS X & Linux:
```
go test ./pkg...
```

## Deploy

Run at the root of the project.

### Run Local

OS X & Linux:
```
go run .
```

Binary:
```
go build .

sudo chmod +x eligibility

./eligibility
```

### Docker Build

Run the command:
```
docker build --tag docker-eligibility .
```

Muilt Stage:
``` 
docker build -t docker-eligibility:multistage -f Dockerfile.multistage .
```

### Docker Run

```
 docker run --rm --name eligibility -p 8080:8080 docker-eligibility
```

Muilt Stage:
```
docker run --rm --name eligibility -p 8080:8080 eligibility:multistage
```
## Command Line

If you want to test an application via the command line here is an example:

```
curl -H "Content-Type: application/json" --data @body.json http://localhost:8080/api/eligibility
```

In the project root inside "json_exampl
