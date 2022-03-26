# Dapp Template

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)

## About <a name = "about"></a>

This project aim to create a minimal template to develop a golang dapp

## Getting Started <a name = "getting_started"></a>

This template use [ganache-cli](https://github.com/trufflesuite/ganache#readme)

### Prerequisites

Golang 1.18 and node 14.18.2

### Installing

To install ganache-cli

```
npm i --save
```

To install go dependencies

```
go get ./...
```

## Usage <a name = "usage"></a>

To launch local blockchain

```
npx ganache-cli
```

Now you should see something like this

```
Ganache CLI v6.12.2 (ganache-core: 2.13.2)

Available Accounts
==================
(0) 0xbD53Ef4e1d736776F4458E9B43f6781CA43596cc (100 ETH)
(1) 0x1a773d57A3A32a722C31568c5BaB024cd8f793d8 (100 ETH)
(2) 0xb7b14ddaF8256Cd1E0DA2Dcc8A662AE2A35FDFA3 (100 ETH)
```

Change the account in the src/client to view the balance of the
account then

```
go run src/client.go
```
