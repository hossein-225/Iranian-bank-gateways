
# ZarinPal Payment Gateway Integration

This Go package provides a simple way to integrate the ZarinPal payment gateway into your Golang application. It allows you to handle payment requests, verify transactions, and more using the ZarinPal API.

## Requirements

- **API Key**: Provided by ZarinPal for accessing the payment gateway.

## Installation

To install the package, use `go get`:

```bash
go get github.com/hossein-225/Iranian-bank-gateways/gateways/zarinpal
```

## Usage

### 1. Initialize the ZarinPal Service

You need to initialize the `ZarinPalService` by passing your API key and the callback URL:

```go
import "github.com/hossein-225/Iranian-bank-gateways/gateways/zarinpal"

func main() {
    service, err := zarinpal.NewService("your-api-key")
    if err != nil {
        log.Fatalf("Failed to initialize ZarinPal service: %v", err)
    }
}
```

### 2. Send a Payment Request

To send a payment request, use the `Request` method. The method requires a `PaymentRequestDto` struct that contains the payment details such as the amount, description, email, and mobile number.

```go
request := zarinpal.PaymentRequestDto{
    Amount:      100000, // Amount in Rials
    Description: "Purchase of product XYZ",
    Email:       "example@example.com", // Optional
    Mobile:      "09120000000", // Optional
    Currency:    "IRR", // Default currency - Optional
    OrderID:     "order12345", // Optional
}

paymentURL, err := service.Request(context.Background(), request)
if err != nil {
    log.Fatalf("Failed to request payment: %v", err)
}

fmt.Println("Redirect the user to:", paymentURL)
```

### 3. Verify a Transaction

Once the user completes the payment, ZarinPal will redirect them to your callback URL with transaction parameters. You can then verify the transaction using the `Verify` method.

```go
verifyResponse, err := service.Verify(context.Background(), "authorityValue", 100000, "order12345")
if err != nil {
    log.Fatalf("Transaction verification failed: %v", err)
}

fmt.Printf("Verification result: %+v
", verifyResponse)
```

### 4. Inquire About a Transaction

To inquire about the status of a transaction, use the `Inquiry` method by passing the `authority` code.

```go
inquiryResponse, err := service.Inquiry(context.Background(), "authorityValue")
if err != nil {
    log.Fatalf("Transaction inquiry failed: %v", err)
}

fmt.Printf("Inquiry result: %+v
", inquiryResponse)
```

### 5. Unverified Transactions

To get a list of unverified transactions, use the `Unverified` method.

```go
unverifiedResponse, err := service.Unverified(context.Background())
if err != nil {
    log.Fatalf("Failed to fetch unverified transactions: %v", err)
}

fmt.Printf("Unverified transactions: %+v
", unverifiedResponse)
```

## Example

Here is a full example of requesting, verifying, and inquiring about a payment:

```go
package main

import (
    "fmt"
    "log"
    "context"
    "github.com/hossein-225/Iranian-bank-gateways/gateways/zarinpal"
)

func main() {
    // Initialize ZarinPal service
    service, err := zarinpal.NewService("your-api-key")
    if err != nil {
        log.Fatalf("Failed to initialize ZarinPal service: %v", err)
    }

    // Send a payment request
    request := zarinpal.PaymentRequestDto{
        Amount:      100000,
        Description: "Purchase of product XYZ",
        Email:       "example@example.com",
        Mobile:      "09120000000",
        Currency:    "IRR",
        OrderID:     "order12345",
    }

    paymentURL, err := service.Request(context.Background(), request)
    if err != nil {
        log.Fatalf("Failed to request payment: %v", err)
    }

    fmt.Println("Redirect the user to:", paymentURL)

    // Verify the transaction
    verifyResponse, err := service.Verify(context.Background(), "authorityValue", 100000, "order12345")
    if err != nil {
        log.Fatalf("Transaction verification failed: %v", err)
    }

    fmt.Printf("Verification result: %+v
", verifyResponse)

    // Inquire about the transaction
    inquiryResponse, err := service.Inquiry(context.Background(), "authorityValue")
    if err != nil {
        log.Fatalf("Transaction inquiry failed: %v", err)
    }

    fmt.Printf("Inquiry result: %+v
", inquiryResponse)

    // Fetch unverified transactions
    unverifiedResponse, err := service.Unverified(context.Background())
    if err != nil {
        log.Fatalf("Failed to fetch unverified transactions: %v", err)
    }

    fmt.Printf("Unverified transactions: %+v
", unverifiedResponse)
}
```

## Error Handling

Errors are handled internally by the package. If a request fails, the error message will be returned, and you can handle it accordingly in your application.

## License

This package is open-source and is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Notes

- Ensure that your IP address and callback URL are properly configured with ZarinPal to allow access to their API.
- Make sure that you handle the transaction reference IDs and response codes securely.
