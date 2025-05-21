# Renovate Go Demo

This is a simple Go application that demonstrates the use of Renovate to automatically update vulnerable dependencies.

## What this program does

This program performs a simple text transformation using the `golang.org/x/text` package. It takes a string with non-ASCII characters and processes it through an encoder.

## How to run

```bash
# Run the program
go run main.go
```

## Vulnerable Dependency

This project deliberately uses a vulnerable version of `golang.org/x/text` (v0.3.6), which is affected by:

- **CVE-2021-38561**: Improper handling of UTF-8 encodings in `golang.org/x/text` before 0.3.7 can lead to security issues.
- **Link**: https://nvd.nist.gov/vuln/detail/CVE-2021-38561

The vulnerability allows for bypassing of security mechanisms or causing resource exhaustion when handling improperly formatted UTF-8 text.

## Renovate Configuration

This repository uses Renovate to automatically detect and fix vulnerable dependencies.