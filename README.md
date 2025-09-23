# Qwik-Gin-DApp

Fast and performance-efficient DApp with Qwik and TypeScript as frontend and Gin and Go as backend.

## 🛠 Built With

[![Node.js Badge](https://img.shields.io/badge/Node.js-393?logo=nodedotjs&logoColor=fff&style=for-the-badge)](https://nodejs.org/en)
[![Qwik Badge](https://img.shields.io/badge/Qwik-3178C6?logo=typescript&logoColor=fff&style=for-the-badge)](https://qwik.builder.io)
[![Go Badge](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=fff&style=for-the-badge)](https://go.dev/)
[![Gin Badge](https://img.shields.io/badge/Gin-008ECF?logo=gin&logoColor=fff&style=for-the-badge)](https://gin-gonic.com/)
[![Solidity Badge](https://img.shields.io/badge/Solidity-363636?logo=solidity&logoColor=fff&style=for-the-badge)](https://soliditylang.org/)
[![Geth Badge](https://img.shields.io/badge/Geth-3C3C3D?logo=ethereum&logoColor=fff&style=for-the-badge)](https://geth.ethereum.org/)
[![MetaMask Badge](https://img.shields.io/badge/MetaMask-3C3C3D?logo=ethereum&logoColor=fff&style=for-the-badge)](https://metamask.io/)
[![Tailwind CSS Badge](https://img.shields.io/badge/Tailwind%20CSS-06B6D4?logo=tailwindcss&logoColor=fff&style=for-the-badge)](https://tailwindcss.com/)

## ⚙️ Run Locally

Clone the project:

```bash
git clone https://github.com/tr1sm0s1n/qwik-gin-dapp.git
cd qwik-gin-dapp
```

### Server

Go to the '**server**' directory:

```bash
cd server
```

Add/update abi and bytecode inside the '**lib**' directory:

```bash
touch lib/Cert.json
touch lib/Cert.bin
```

Install abigen:

```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

Generate Go binding for contract:

```bash
abigen --v2 --abi lib/Cert.json --bin lib/Cert.bin --pkg lib --type Cert --out lib/Cert.go
```

Run a simulated blockchain and add its JSON-RPC URL and a private key (with sufficient balance) to the '**.env**' file inside '**server**'.

Deploy contract:

```bash
go run scripts/deploy.go
```

Start the server:

```bash
go run .
```

Use air for live reload (Optional):

```bash
# install air
go install github.com/air-verse/air@latest
# start air
air
```

### Client

Go to the '**client**' directory:

```bash
cd client
```

Install pnpm (Optional):

```bash
npm i -g pnpm
```

Install dependencies:

```bash
pnpm install
```

Import the deployer account to the **MetaMask** extension using the private key.

Start the client:

```bash
pnpm dev
```

Click **Connect** to create a sign. Copy the sign to the '**.env**' file inside '**client**' and connect again.

## 🐋 Run via Docker

Deploy contract:

```bash
just deploy-d
```

Build and run (detached) the application containers:

```bash
just start
```

Stop the application containers:

```bash
just stop
```

**_Note_**: Docker uses '**.env**' file in the root. Update the file before the build.
