# Simple-transfer-system

To run with default settings:
```
docker compose up
go mod tidy
go run ./cmd
```

Make sure `docker compose up` brings up a fresh postgresql container for migration scripts to take effect

## Assumptions made:
* Account balance can only be positive
* Transaction amount must be positive and `source` account 'transfers' to `destination` account
* Calculation and presentation and storage of amount and balance are up to 5 decimal places
* Account id is integer and within range of int64
* Since this service is intended for internal use, error is not masked before sending out in response
* Account balance is limited to 131067 (131072-5) digits before decimal point, and 5 digits after decimal point
* `initial_balance` and `amount` accepts both string and number
