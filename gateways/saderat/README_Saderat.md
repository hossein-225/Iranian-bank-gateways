# Saderat Gateway Integration

This Go package provides a simple way to integrate the Saderat payment gateway into your Golang application. It allows you to send payment requests and advise transactions, and more using the Saderat Bank Payment Gateway API.

## Requirements

- **Terminal ID**: Provided by Saderat Bank.

## Installation

To install the package, use `go get`:

```bash
go get github.com/hossein-225/Iranian-bank-gateways/gateways/saderat
```

## Usage

### 1. Initialize the Saderat Payment Service

You need to initialize the `PaymentService` service by passing your API key.

```go
import "github.com/hossein-225/Iranian-bank-gateways/gateways/saderat"

func main() {
	service, err := saderat.NewPaymentService("your-terminal-id")
	if err != nil {
		log.Fatalf("Failed to initialize Saderat payment service: %v", err)
	}
}
```

### 2. Send a Payment Request

To send a payment request, use the `SendRequest` method. The method requires the amount,redirect URL, InvoiceID and Additional data under the Payload field.

```go
    paymentToken, err := service.SendRequest(context.Background(), 100000, "https://your-callback-url.com", "123456", "payload")
    if err != nil {
    	log.Fatalf("Failed to request paymentToken: %v", err)
    }

    fmt.Printf("Your payment token: %v\n", paymentToken)
```

After a successful request, you will receive a payment token that the user needs to be redirected to in order to complete the transaction. Once the paymentToken is received, you need to construct an HTML form to redirect the user to the payment gateway. The form should be sent using the POST method, and it must include both the TerminalID and the paymentToken.

```html
<form
  id="paymentTokenForm"
  action="https://mabna.shaparak.ir:8080/Pay"
  method="post"
>
  <input type="hidden" name="TerminalID" value="your-Terminal-id" />
  <input type="hidden" name="token" value="your-Payment-Token" />
</form>
<script type="text/javascript">
  document.getElementById("paymentTokenForm").submit();
</script>
```

TerminalID: This is the Terminal ID provided by the bank, which identifies your merchant account.

Token: This is the payment token returned from the SendRequest method.

The form is automatically submitted by the JavaScript code, redirecting the user to the payment gateway page where they can finalize the payment. You need to replace your-Terminal-id and your-Payment-Token with actual values in your application.

### 3. advise a Transaction

Once the user completes the payment, Saderat will redirect them to your callback URL with transaction parameters. You can then advise the transaction using the `ConfirmTransaction` method.

```go
    adviseResponse, err := service.ConfirmTransaction(context.Background(), "digitalReceipt", 100000)
    if err != nil {
    	log.Fatalf("Transaction verification failed: %v", err)
    }

    fmt.Printf("Verification result: %v", adviseResponse)
```

### 4. rollBack a Transaction

If you need to rollBack a transaction (e.g., if the transaction is not completed), you can use the `RollbackTransaction` method.

```go
    rollBackResponse, err := service.RollbackTransaction(context.Background(), "digitalReceipt", 100000)
    if err != nil {
    	log.Fatalf("Transaction reversal failed: %v\n", err)
    }

    fmt.Printf("Reversal result: %v\n", rollBackResponse)
```

## Example

Here is a full example of requesting, verifying, and reversing a payment:

```go
package main

import (
"context"
"fmt"
"log"

    "github.com/hossein-225/Iranian-bank-gateways/gateways/saderat"

)

func main() {
	// Initialize Saderat payment service
	service, err := saderat.NewPaymentService("your-terminal-id")
	if err != nil {
		log.Fatalf("Failed to initialize Saderat payment service: %v", err)
	}


    // Send a payment request and get the payment token
    paymentToken, err := service.SendRequest(context.Background(), 100000, "https://your-callback-url.com", "123456", "payload")
    if err != nil {
    	log.Fatalf("Failed to request paymentToken: %v", err)
    }

    fmt.Printf("Your payment token: %v\n", paymentToken)

	/* After successfully receiving the paymentToken from the SendRequest method,
	 you should generate and inject an HTML form into the response that includes
	 the TerminalID and paymentToken, and use JavaScript to automatically submit
	 the form. This will redirect the user to the payment gateway page to complete the transaction 
	*/

	fmt.Printf(`
	<form id="paymentTokenForm" action="https://mabna.shaparak.ir:8080/Pay" method="post">
    	<input type="hidden" name="TerminalID" value="your-terminal-id"/>
   		<input type="hidden" name="token" value="%s"/>
	</form>
	<script type="text/javascript">
    	document.getElementById('paymentTokenForm').submit();
	</script>
	`, paymentToken)

    // Advise the transaction
    adviseResponse, err := service.ConfirmTransaction(context.Background(), "digitalReceipt", 100000)
    if err != nil {
    	log.Fatalf("Transaction verification failed: %v", err)
    }

    fmt.Printf("Verification result: %v", adviseResponse)

    // RollBack the transaction if necessary
    rollBackResponse, err := service.RollbackTransaction(context.Background(), "digitalReceipt", 100000)
    if err != nil {
    	log.Fatalf("Transaction reversal failed: %v\n", err)
    }

    fmt.Printf("Reversal result: %v\n", rollBackResponse)

}
```

## Error Handling

Errors are handled internally by the package. If a request fails, the error message will be returned, and you can handle it accordingly in your application.

## License

This package is open-source and is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Notes

- Ensure that your IP address and callback URL are properly configured with Saderat Bank to allow access to their API.
- Make sure that you handle the transaction digital receipts and response codes securely.
