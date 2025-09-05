# Credential Management

This document describes how to manage credentials for the CProtocol scanner.

## Kraken API Credentials

The Kraken API key and secret should be stored as environment variables:

- `KRAKEN_API_KEY`
- `KRAKEN_API_SECRET`

For production deployments, it is recommended to use a secret management tool like HashiCorp Vault or AWS Secrets Manager to store these credentials securely.
