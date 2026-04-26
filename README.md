# Saltpack Decryptor (Go)

A lightweight CLI tool written in Go for decrypting saltpack signcrypted messages using a hex-encoded X25519 key.

## Features

- Decrypts saltpack signcrypted data
- Uses 32-byte hex-encoded X25519 private keys
- Supports file input and output
- Cross-platform (Linux, Windows, macOS)
- Minimal dependencies

## Usage

```bash
go run main.go -key <hex-key> -in <input-file> [-out <output-file>]
