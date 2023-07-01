# Qwik-Gin-DApp

Fast and performance-efficient DApp with Qwik and TypeScript as frontend and Gin and Go as backend.

## 🛠 Built With

<div align="left">
<a href="https://nodejs.org/en/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/nodejs.svg" width="36" height="36" alt="NodeJS" /></a>
<a href="https://qwik.builder.io/docs/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/qwik.svg" width="36" height="36" alt="Qwik" /></a>
<a href="https://go.dev/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/go.svg" width="36" height="36" alt="Go" /></a>
<a href="https://gin-gonic.com/docs/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/gin.svg" width="36" height="36" alt="Gin" /></a>
<a href="https://soliditylang.org/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/solidity.svg" width="36" height="36" alt="Solidity" /></a>
<a href="https://metamask.io/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/metamask.svg" width="36" height="36" alt="MetaMask" /></a>
<a href="https://tailwindcss.com/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/tailwindcss.svg" width="36" height="36" alt="TailwindCSS" /></a>
</div>

## ⚙️ Run Locally

Clone the project:

```bash
git clone https://github.com/DEMYSTIF/qwik-gin-dapp.git
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

Add the account address to the '**.env**' file inside '**client**'.

Start the client:

```bash
pnpm dev
```

Click **Connect** to issue certificate.

## 📜 License

Click [here](./LICENSE.md).

## 🎗️ Contributing

Click [here](./CONTRIBUTING.md).

## ⚖️ Code of Conduct

Click [here](./CODE_OF_CONDUCT.md).
