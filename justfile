server_dir := `pwd` + '/server'
client_dir := `pwd` + '/client'
pkg_m := `if command -v pnpm > /dev/null; then echo pnpm; else echo npm; fi`
contract := 'Cert'

# Display recipes.
default:
    @just --list

# Install abigen.
abigen:
    @mkdir -p {{ server_dir }}/tmp
    @GOBIN={{ server_dir }}/tmp go install github.com/ethereum/go-ethereum/cmd/abigen@latest

# Create binding.
binding:
    @{{ server_dir }}/tmp/abigen --v2 --bin {{ server_dir }}/lib/{{ contract }}.bin --abi {{ server_dir }}/lib/{{ contract }}.json --pkg lib --type {{ contract }} --out {{ server_dir }}/lib/{{ contract }}.go

# Deploy contract.
deploy:
    @cd {{ server_dir }} && go run ./scripts/deploy.go

# Run the server with live reload.
dev-server:
	@if [ ! -f {{ server_dir }}/tmp/air ]; then \
		echo "Installing air..."; \
		mkdir -p {{ server_dir }}/tmp; \
		GOBIN={{ server_dir }}/tmp go install github.com/air-verse/air@latest; \
	fi
	@cd {{ server_dir }} && {{ server_dir }}/tmp/air

# Install client dependencies
install:
    @cd {{ client_dir }} && {{ pkg_m }} install

# Run the client in dev mode.
dev-client:
    @cd {{ client_dir }} && {{ pkg_m }} run dev

# Deploy using Docker.
deploy-d:
    @docker compose up deployer --build
    @sleep 1
    @docker compose down deployer

# Start applications via Docker.
start:
    @docker compose up client server --build -d

# Stop Docker containers.
stop:
    @docker compose down
