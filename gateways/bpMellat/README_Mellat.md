
# Mellat Payment Gateway Integration

This Go package provides a simple way to integrate the Mellat payment gateway into your Golang application. It allows you to handle payment requests, verify transactions, and more using the Mellat Bank Payment Gateway API.

## Requirements

- **Terminal ID**: Provided by Mellat Bank.
- **Username and Password**: Credentials provided by Mellat Bank for accessing the payment gateway.

## Installation

To install the package, use `go get`:

```bash
go get github.com/hossein-225/Iranian-bank-gateways/gateways/bpmellat
```

## Usage

### 1. Initialize the Mellat Service

You need to initialize the `BpMellat` service by passing your terminal ID, username, and password:

```go
import "github.com/hossein-225/Iranian-bank-gateways/gateways/bpmellat"

func main() {
    service, err := bpmellat.NewService("your-terminal-id", "your-username", "your-password")
    if err != nil {
        log.Fatalf("Failed to initialize Mellat service: %v", err)
    }
}
```

### 2. Send a Payment Request

To send a payment request, use the `BpPayRequest` method. The method requires a `BpPayRequest` struct that contains the payment details such as the order ID, amount, and callback URL.

```go
request := bpmellat.BpPayRequest{
    OrderID:         12345,
    SaleOrderID:     54321,
    SaleReferenceID: 0, // Use 0 for a new transaction
    Amount:          100000, // Amount in Rials (e.g., 100,000 Rials)
    LocalDate:       "20230101",
    LocalTime:       "120000",
    AdditionalData:  "Purchase of product ABC", // Optional
    CallBackURL:     "https://your-callback-url.com",
    PayerID:         0, // Optional
}

paymentURL, err := service.BpPayRequest(context.Background(), request)
if err != nil {
    log.Fatalf("Failed to request payment: %v", err)
}

fmt.Println("Redirect the user to:", paymentURL)
```

### 3. Verify a Transaction

Once the user completes the payment, Mellat will redirect them to your callback URL with transaction parameters. You can then verify the transaction using the `BpVerifyRequest` method.

```go
verifyRequest := bpmellat.BpRequest{
    OrderID:         12345,
    SaleOrderID:     54321,
    SaleReferenceID: 1234567890,
}

err := service.BpVerifyRequest(context.Background(), verifyRequest)
if err != nil {
    log.Fatalf("Transaction verification failed: %v", err)
}

fmt.Println("Transaction verified successfully")
```

### 4. Settle a Transaction

After verifying a successful transaction, you should settle the payment using the `BpSettleRequest` method.

```go
err := service.BpSettleRequest(context.Background(), verifyRequest)
if err != nil {
    log.Fatalf("Transaction settlement failed: %v", err)
}

fmt.Println("Transaction settled successfully")
```

### 5. Reverse a Transaction

If you need to reverse a transaction (e.g., if the transaction is not completed), you can use the `BpReversalRequest` method.

```go
err := service.BpReversalRequest(context.Background(), verifyRequest)
if err != nil {
    log.Fatalf("Transaction reversal failed: %v", err)
}

fmt.Println("Transaction reversed successfully")
```

### 6. Inquire About a Transaction

To inquire about the status of a transaction, use the `BpInquiryRequest` method.

```go
err := service.BpInquiryRequest(context.Background(), verifyRequest)
if err != nil {
    log.Fatalf("Transaction inquiry failed: %v", err)
}

fmt.Println("Transaction inquiry successful")
```

## Example

Here is a full example of requesting, verifying, settling, and reversing a payment:

```go
package main

import (
    "fmt"
    "log"
    "context"
    "github.com/hossein-225/Iranian-bank-gateways/gateways/bpmellat"
)

func main() {
    // Initialize Mellat service
    service, err := bpmellat.NewService("your-terminal-id", "your-username", "your-password")
    if err != nil {
        log.Fatalf("Failed to initialize Mellat service: %v", err)
    }

    // Send a payment request
    request := bpmellat.BpPayRequest{
        OrderID:         12345,
        SaleOrderID:     54321,
        SaleReferenceID: 0,
        Amount:         100000,
        LocalDate:      "20230101",
        LocalTime:      "120000",
        AdditionalData: "Purchase of product ABC",
        CallBackURL:    "https://your-callback-url.com",
        PayerID:        0,
    }

    paymentURL, err := service.BpPayRequest(context.Background(), request)
    if err != nil {
        log.Fatalf("Failed to request payment: %v", err)
    }

    fmt.Println("Redirect the user to:", paymentURL)

    // Verify the transaction
    verifyRequest := bpmellat.BpRequest{
        OrderID:         12345,
        SaleOrderID:     54321,
        SaleReferenceID: 1234567890,
    }

    err = service.BpVerifyRequest(context.Background(), verifyRequest)
    if err != nil {
        log.Fatalf("Transaction verification failed: %v", err)
    }

    fmt.Println("Transaction verified successfully")

    // Settle the transaction
    err = service.BpSettleRequest(context.Background(), verifyRequest)
    if err != nil {
        log.Fatalf("Transaction settlement failed: %v", err)
    }

    fmt.Println("Transaction settled successfully")
}
```

## Error Handling

Errors are handled internally by the package. If a request fails, the error message will be returned, and you can handle it accordingly in your application.

## License

This package is open-source and is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Notes

- Ensure that your IP address and callback URL are properly configured with Mellat Bank to allow access to their API.
- Make sure that you handle the transaction reference IDs and response codes securely.
