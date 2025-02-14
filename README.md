# ledger-go

[![Actions Status](https://github.com/oasislabs/ledger-go/workflows/CI/badge.svg)](https://github.com/oasislabs/ledger-go/actions)


This project provides a library to connect to ledger devices.

It handles USB (HID) communication and APDU encapsulation.

Linux, OSX and Windows are supported.

# Get source
Apart from cloning, be sure you install dep dependency management tool
https://github.com/golang/dep

## Setup
Install `dep`:
```
go get -v -u github.com/golang/dep/cmd/dep
```

Update dependencies using the following:
```
dep ensure
```

# Building
```
go build
```

# Testing
To run the tests for the APDUWrapper, run the following:
```
go test
```

To run the tests in `ledger_test.go`, connect a ledger device to the computer, enter the passcode, choose an existing app, then run the following:
```
go test -tags ledger_device
```

For tests against the Ethereum Ledger App, make sure that the Ethereum App is installed on the device, connect it, and navigate to the Ethereum app. Then, run the following:
```
go test -tags "ledger_device eth_ledger_app"
```
