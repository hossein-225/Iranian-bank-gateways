# BitPay Gateway Integration

This Go package provides a simple way to integrate the BitPay payment gateway into your Golang application. It allows you to send payment requests and verify transactions using the BitPay API.

## Requirements

- **API Key**: You need to obtain an API key from your BitPay account.

## Installation

To install the package, use `go get`:

```bash
go get github.com/hossein-225/Iranian-bank-gateways/gateways/bitpay
```

## Usage

### 1. Initialize the BitPay Service

You need to initialize the `BitPayIR` service by passing your API key.

```go
import "github.com/hossein-225/Iranian-bank-gateways/gateways/bitpay"

func main() {
    service, err := bitpay.NewService("your-api-token")
    if err != nil {
        log.Fatalf("Failed to initialize BitPay service: %v", err)
    }
}
```

### 2. Send a Payment Request

To send a payment request, use the `Request` method. The method requires a `BitPayRequest` struct that contains the payment details such as the amount, order ID, name, email, description, and callback URL.

```go
request := &bitpay.BitPayRequest{
    Amount:      1000, // Amount in Toman (e.g., 1000 Toman) - Necessary
    OrderID:     "order12345", // Optional
    Name:        "Hossein Hosseini", // Optional
    Email:       "hossein@example.com", // Optional
    Description: "Purchase of product XYZ", // Optional
    CallbackURL: "https://your-callback-url.com", // Necessary
}

paymentURL, err := service.Request(context.Background(), request)
if err != nil {
    log.Fatalf("Failed to request payment: %v", err)
}

fmt.Println("Redirect the user to:", paymentURL)
```

After a successful request, you will receive a payment URL that the user needs to be redirected to complete the transaction.

### 3. Verify a Transaction

Once the user completes the payment, BitPay will redirect them to your callback URL with `trans_id` and `id_get` parameters. You can then verify the transaction using these parameters.

```go
result, err := service.Verify(context.Background(), "trans_id_value", "id_get_value")
if err != nil {
    log.Fatalf("Transaction verification failed: %v", err)
}

fmt.Printf("Verification result: %v", result)
```

The `Verify` method will return a map containing transaction details if successful.

## Example

Here is a full example:

```go
package main

import (
    "fmt"
    "log"
    "context"
    "github.com/hossein-225/Iranian-bank-gateways/gateways/bitpay"
)

func main() {
    // Initialize BitPay service
    service, err := bitpay.NewService("your-api-token")
    if err != nil {
        log.Fatalf("Failed to initialize BitPay service: %v", err)
    }

    // Send a payment request
    request := &bitpay.BitPayRequest{
        Amount:      1000, // Amount in Toman - Necessary
        OrderID:     "order12345", // Optional
        Name:        "Hossein Hosseini", // Optional
        Email:       "hossein@example.com", // Optional
        Description: "Purchase of product XYZ", // Optional
        CallbackURL: "https://your-callback-url.com", // Necessary
    }

    paymentURL, err := service.Request(context.Background(), request)
    if err != nil {
        log.Fatalf("Failed to request payment: %v", err)
    }

    fmt.Println("Redirect the user to:", paymentURL)

    // After the payment, verify the transaction
    result, err := service.Verify(context.Background(), "trans_id_value", "id_get_value")
    if err != nil {
        log.Fatalf("Transaction verification failed: %v", err)
    }

    fmt.Printf("Verification result: %v", result)
}
```

## Error Handling

Errors are handled internally by the package. If a request fails, the error message will be returned from the BitPay API, and you can handle it accordingly in your application.

## License

This package is open-source and is licensed under the [MIT License](https://opensource.org/licenses/MIT).
