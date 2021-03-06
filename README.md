# Go bindings and CLI tool for the Skycoin hardware wallet

[![Build Status](https://travis-ci.com/skycoin/hardware-wallet-go.svg?branch=master)](https://travis-ci.com/skycoin/hardware-wallet-go)

## Installation

### Install golang

    https://github.com/golang/go/wiki/Ubuntu

## Usage

### Download source code
    
    go get github.com/skycoin/hardware-wallet-go

### Dependancies management

This project uses dep [dependancy manager](https://github.com/golang/dep).

Don't modify anything under vendor/ directory without using [dep commands](https://github.com/golang/dep/blob/master/docs/Gopkg.toml.md).

Download dependencies using command:

    dep ensure

### Generate protobuf files

If you need to generate google protobuf files yourself (if you are creating a new message for instance). You need to [install protoc](http://google.github.io/proto-lens/installing-protoc.html) first, then run the following:

```bash
cd device-wallet
protoc -I ./messages  --go_out=./messages messages/messages.proto messages/types.proto messages/descriptor.proto 
```

### Run

    go run cli.go

See also [CLI.md](https://github.com/skycoin/hardware-wallet-go/blob/master/CLI.md) for information about the Command Line Interface.

## Wiki

More information in [the wiki](https://github.com/skycoin/hardware-wallet-go/wiki)
