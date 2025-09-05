# USD-Only (No Stablecoins)

This document explains why the CProtocol scanner only supports USD pairs and does not support stablecoin pairs.

## Simplicity

The primary reason for this decision is simplicity. By only supporting USD pairs, we can avoid the complexity of dealing with different quote currencies and the need to convert between them.

## Redundancy

Stablecoin pairs such as USDT, USDC, and BUSD are redundant with USD pairs. They are all pegged to the US dollar, so there is no need to support them separately.
