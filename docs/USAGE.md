# Usage

This document describes how to use the CProtocol scanner.

## Scanning for Opportunities

To scan for trading opportunities, use the `scan` command:

```
./cgemini scan
```

By default, the scanner will use the Kraken exchange and scan for USD pairs.

### Dry Run

To perform a dry run without executing any trades, use the `--dry-run` flag:

```
./cgemini scan --dry-run
```

### Changing the Exchange

To use a different exchange, use the `--exchange` flag:

```
./cgemini scan --exchange binance
```

### Changing the Quote Currency

To scan for a different quote currency, use the `--pairs` flag:

```
./cgemini scan --pairs USDT
```
