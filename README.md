# Simple-transfer-system

To run with default settings:
```
docker compose up
go mod tidy
go run ./cmd
```

Make sure `docker compose up` brings up a fresh postgresql container for migration scripts to take effect

## Endpoints
* `POST` `localhost:8080/accounts`
    * Request Body:
    ```
    {
        "account_id":210,
        "initial_balance":1000
    }
    ```
    * Response:
    ```
    {
        "code": 0,
        "data": null,
        "message": "success"
    }
    ```
* `GET` `localhost:8080/accounts/:account_id`
    * Response:
    ```
    {
        "code": 0,
        "data": {
            "account_id": 210,
            "balance": "1000.00000"
        },
        "message": "success"
    }
    ```
* `POST` `localhost:8080/transactions`
    * Request Body:
    ```
    {
        "source_account_id":123,
        "destination_account_id":456,
        "amount":"111.3333"
    }
    ```
    * Response:
    ```
    {
        "code": 0,
        "data": null,
        "message": "success"
    }
    ```

* Unsuccessful requests' response code will not be 0

## Assumptions made:
* Account balance can only be positive
* Transaction amount must be positive and `source` account 'transfers' to `destination` account
* Calculation and presentation and storage of amount and balance are up to 5 decimal places
* Account id is integer and within range of int64
* Since this service is intended for internal use, error is not masked before sending out in response
* Account balance is limited to 131067 (131072-5) digits before decimal point, and 5 digits after decimal point
* `initial_balance` and `amount` accepts both string and number

## Improvement opportunities
* Use interface for constructors to allow mock and unit testing
* Use mocksql to mock database behavior
* Use `wire` code-gen for dependency injection
