
# Saman Payment Gateway Integration

This Go package provides a simple way to integrate the Saman payment gateway into your Golang application. It allows you to handle payment requests, verify transactions, and more using the Saman Bank Payment Gateway API.

## Requirements

- **Terminal ID**: Provided by Saman Bank.

## Installation

To install the package, use `go get`:

```bash
go get github.com/hossein-225/Iranian-bank-gateways/gateways/saman
```

## Usage

### 1. Initialize the Saman Payment Service

You need to initialize the `PaymentService` by passing your terminal ID:

```go
import "github.com/hossein-225/Iranian-bank-gateways/gateways/saman"

func main() {
    service, err := saman.NewPaymentService("your-terminal-id")
    if err != nil {
        log.Fatalf("Failed to initialize Saman payment service: %v", err)
    }
}
```

### 2. Send a Payment Request

To send a payment request, use the `SendRequest` method. The method requires the amount, transaction ID (ResNum), customer’s phone number, and redirect URL.

```go
paymentURL, err := service.SendRequest(context.Background(), 100000, "order12345", "09120000000", "https://your-callback-url.com")
if err != nil {
    log.Fatalf("Failed to request payment: %v", err)
}

fmt.Println("Redirect the user to:", paymentURL)
```

After a successful request, you will receive a payment URL that the user needs to be redirected to complete the transaction.

### 3. Verify a Transaction

Once the user completes the payment, Saman will redirect them to your callback URL with transaction parameters. You can then verify the transaction using the `Verify` method.

```go
verifyResponse, err := service.Verify(context.Background(), "refNumValue")
if err != nil {
    log.Fatalf("Transaction verification failed: %v", err)
}

fmt.Printf("Verification result: %+v
", verifyResponse)
```

### 4. Reverse a Transaction

If you need to reverse a transaction (e.g., if the transaction is not completed), you can use the `Reverse` method.

```go
reverseResponse, err := service.Reverse(context.Background(), "refNumValue")
if err != nil {
    log.Fatalf("Transaction reversal failed: %v", err)
}

fmt.Printf("Reversal result: %+v
", reverseResponse)
```

## Example

Here is a full example of requesting, verifying, and reversing a payment:

```go
package main

import (
    "fmt"
    "log"
    "context"
    "github.com/hossein-225/Iranian-bank-gateways/gateways/saman"
)

func main() {
    // Initialize Saman payment service
    service, err := saman.NewPaymentService("your-terminal-id")
    if err != nil {
        log.Fatalf("Failed to initialize Saman payment service: %v", err)
    }

    // Send a payment request
    paymentURL, err := service.SendRequest(context.Background(), 100000, "order12345", "09120000000", "https://your-callback-url.com")
    if err != nil {
        log.Fatalf("Failed to request payment: %v", err)
    }

    fmt.Println("Redirect the user to:", paymentURL)

    // Verify the transaction
    verifyResponse, err := service.Verify(context.Background(), "refNumValue")
    if err != nil {
        log.Fatalf("Transaction verification failed: %v", err)
    }

    fmt.Printf("Verification result: %+v
", verifyResponse)

    // Reverse the transaction if necessary
    reverseResponse, err := service.Reverse(context.Background(), "refNumValue")
    if err != nil {
        log.Fatalf("Transaction reversal failed: %v", err)
    }

    fmt.Printf("Reversal result: %+v
", reverseResponse)
}
```

## Error Handling

Errors are handled internally by the package. If a request fails, the error message will be returned, and you can handle it accordingly in your application.

## License

This package is open-source and is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Notes

- Ensure that your IP address and callback URL are properly configured with Saman Bank to allow access to their API.
- Make sure that you handle the transaction reference IDs and response codes securely.
