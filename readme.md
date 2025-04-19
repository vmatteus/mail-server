# Simple SMTP Server (WIP)

A complete SMTP server implementation in Go that receives emails and delivers them directly to recipient mail servers.

## Features

- Receives emails on configurable port (default: 1025)
- Resolves MX records for recipient domains
- Implements mail queue for reliable delivery
- Handles multiple recipients
- Graceful shutdown with active connection handling

## Requirements

- Go 1.18 or higher

## Installation

## Usage

Run the server:

The server will start listening on port `1025` (by default).

## Configuration

Edit the settings at the beginning of `main.go` to customize:

- Listening port
- Domain name
- Message size limits
- Timeout values
- Number of recipients

## Testing

You can test the server using standard mail clients by configuring them to use your server (`localhost:1025`) as the outgoing SMTP server.

Alternatively, use the `telnet` command:

Then use SMTP commands.

## Limitations

- No TLS support yet
- Limited spam protection
- No authentication
- No DKIM/SPF implementation

## License

MIT