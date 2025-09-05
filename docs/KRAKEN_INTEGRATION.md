# Kraken Integration

This document describes the specifics of the Kraken integration.

## API Client

The Kraken API client is implemented in the `api/kraken.go` file. It uses the official Kraken API and supports the following features:

- Fetching ticker, OHLC, and order book data
- WebSocket support for real-time data
- Rate limiting
- Circuit breaker

## Primary Exchange

Kraken is the primary exchange for the CProtocol scanner. All microstructure data is fetched from Kraken, and all trades are executed on Kraken.
