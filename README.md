# qrencode

![Version](https://img.shields.io/badge/version-0.2.0-blue)

Fast QR code encoder written in Go.

> **Fork notice:** This project is a fork of [d3vilh/qrencode](https://github.com/d3vilh/qrencode).

## changelog

See [CHANGELOG.md](CHANGELOG.md) for a full list of changes.

## build

```shell
git clone https://github.com/OZON08/qrencode
cd qrencode
go build -o qrencode main.go
```

## usage

```shell
qrencode "some text" > qr.png
```