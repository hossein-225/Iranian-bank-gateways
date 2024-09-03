# Iranian Bank Gateways

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/hossein-225/Iranian-bank-gateways/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/hossein-225/iranian-bank-gateways)](https://goreportcard.com/report/github.com/hossein-225/iranian-bank-gateways)
[![Go Reference](https://pkg.go.dev/badge/github.com/hossein-225/iranian-bank-gateways.svg)](https://pkg.go.dev/github.com/hossein-225/iranian-bank-gateways)
[![GitHub contributors](https://img.shields.io/github/contributors/hossein-225/iranian-bank-gateways)](https://github.com/hossein-225/Iranian-bank-gateways/graphs/contributors)

A free and open-source library written in Go (Golang) designed to streamline the integration process with Iranian bank payment gateways.

## Introduction

When developing Go-based projects that require connections to Iranian payment gateways, developers traditionally had to write custom code for each integration. This repository eliminates that need by providing reusable code for common payment gateway interactions.

## Features

- **Request and Retrieve Token**: Simplify the process of obtaining tokens for transactions.
- **Payment Verification**: Easily verify payments across supported gateways.
- **Supported Gateways**:
  - **Mellat Bank**
  - **Bitpay**
  - **Zarinpal**

## Installation

To install the library, you can use either of the following commands:

```sh
go get github.com/hossein-225/iranian-bank-gateways
```

or

```sh
go install github.com/hossein-225/iranian-bank-gateways
```

## Usage

To use the library, create a new instance of the gateway service and utilize the available functions as per the documentation. Here’s a quick example:

```go
import (
    "github.com/hossein-225/iranian-bank-gateways/bpMellat"
)

func main() {
    gateway := bpMellat.NewService("your-terminal-id", "your-username", "your-password")

    // Add your logic to handle payment request
}

```

## Configuration

Configuration depends on the gateway you're integrating with:

- **Mellat Bank**: Requires a username and password, which are provided by the bank.
- **Bitpay**: Requires an API key, obtainable from your Bitpay account.

Ensure you have the necessary credentials before attempting integration.

## Testing

Testing support is not provided yet. Contributions are welcome to add test coverage.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

## Support

For support, please contact [@hossein-225](https://github.com/hossein-225).

## Contributors

- [@mostafaparaste](https://github.com/mostafaparaste)

- [@mhap75](https://github.com/mhap75)

## Changelog

A changelog has not been provided yet. Major updates will be documented in future releases.
