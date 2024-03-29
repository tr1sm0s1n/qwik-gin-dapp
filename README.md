# Qwik-Gin-DApp

Fast and performance-efficient DApp with Qwik and TypeScript as frontend and Gin and Go as backend.

## 🛠 Built With

[![Node.js](https://img.shields.io/badge/node.js-olivedrab?style=for-the-badge&logo=node.js&logoColor=white)](https://nodejs.org/en)
[![Qwik](https://img.shields.io/badge/qwik-olivedrab?style=for-the-badge&logo=node.js&logoColor=white)](https://qwik.builder.io)
[![Go](https://img.shields.io/badge/go-dodgerblue?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Gin](https://img.shields.io/badge/gin-dodgerblue?style=for-the-badge&logo=go&logoColor=white)](https://gin-gonic.com/)
[![Solidity](https://img.shields.io/badge/solidity-sienna?style=for-the-badge&logo=solidity&logoColor=white)](https://soliditylang.org/)
[![Geth](https://img.shields.io/badge/geth-darkslategray?style=for-the-badge&logo=ethereum&logoColor=white)](https://geth.ethereum.org/)
[![MetaMask](https://img.shields.io/badge/metamask-darkslategray?style=for-the-badge&logo=ethereum&logoColor=white)](https://metamask.io/)
[![TailwindCSS](https://img.shields.io/badge/tailwindcss-indigo?style=for-the-badge&logo=tailwindcss&logoColor=white)](https://tailwindcss.com/)

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

Create a '**build**' directory and add abi and bytecode inside:

```bash
mkdir build
touch build/Cert.abi
touch build/Cert.bin
```

Install abigen:

```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

Generate Go binding for contract:

```bash
abigen --abi build/Cert.abi --bin build/Cert.bin --pkg lib --type Cert --out lib/Cert.go
```

Run a simulated blockchain on port **8545**, and add a private key to the '**.env**' file inside '**server**'.

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
go install github.com/cosmtrek/air@latest
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

## 📜 License

Click [here](./LICENSE.md).

## 🎗️ Contributing

Click [here](./CONTRIBUTING.md).

## ⚖️ Code of Conduct

Click [here](./CODE_OF_CONDUCT.md).
