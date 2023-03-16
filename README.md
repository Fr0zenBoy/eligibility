# Authoraizer Post

Function that authorizes a transaction for a specific account, following some predefined rules.


Rules:

```
1. The transaction amount should not be above limit
2. No transaction should be approved when the card is blocked
3. The first transaction shouldn't be above 90% of the limit
4. There should not be more than 10 transactions on the same merchant
5. Merchant blacklist
6. There should not be more than 3 transactions on a 2 minutes interval
```

## How it works

Sending a dummy transaction in json format for application at http://localhost:8080/api/authoraizer.

### Schema

#### Input

```json
{
  "account": {
    "cardIsActive": "true",
    "limit": "5000",
    "denyList": ["Moes"],
    "isWhitelisted": "true"
  },
  "transaction": {
    "merchant": "MacLarens",
    "amount": "2000",
    "time": "2019-06-19 21:04:00"
  },
  "lastTransactions": [
    {
      "merchant": "MacLarens",
      "amount": "1000",
      "time": "2019-06-19 21:01:00"
    }
  ]
}
```

## You will receive output similar to this.

### Output

```json
{
    "approved": "True",
    "newLimit": "3000.0",
    "deniedReasons": []
}
```

## Prerequisites

To run the application locally and run the tests it is necessary to install python3 and the dependencies in requirements.txt, here i am using python 3.7.x

Run at the root of the project.

OS X & Linux:
```
wip
```

Windows:
```
wip
```


## Running Unittest

Run at the root of the project.


Run unittest fremework in OS X & Linux:
```
wip
```

Run unittest fremework in Windows:
```
wip
```


## Deploy

Run at the root of the project.

### Run Local

OS X & Linux:
```
wip
```

Windows:
```
wip
```

### Docker Run

Run the command:
```
docker build --no-cache -t transaction:latest .

docker run --rm -it --name transaction -p 8080:8080 transaction:latest
```


## Command Line

If you want to test an application via the command line here is an example:

```
curl -H "Content-Type: application/json" --data @body.json http://localhost:8080/api/authoraizer
```

In the project root inside "json_examples" folder you also find examples of json to use.
