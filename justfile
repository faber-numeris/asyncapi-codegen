# justfile for asyncapi-codegen project

# Default recipe to run all checks
check: check-generation lint test
	just check-generation && just lint && just test

check-generation: ## Check files are generated locally
	sh ./scripts/check-generation.sh

clean: local-env/stop ## Clean the project locally
	just local-env/stop
	rm -rf ./tmp/certs

generate: ## Generate files locally
	go generate ./...

lint: ## Lint the code locally
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.0 run ./...

local-env/start: ## Start the local environment
	go run ./tools/generate-certs
	docker compose up -d

local-env/stop: ## Stop the local environment
	docker compose stop

local-env/teardown: ## Kill containers and delete volumes
	docker compose down --volumes

publish TAG=: ## Publish with tag on git, docker hub, etc. locally
	# The original Makefile had a 'dagger/publish' dependency which is not directly translatable without more context.
	# You might need to define a 'dagger/publish' recipe or integrate its functionality here.
	git tag {{TAG}} && git push origin {{TAG}}

test: local-env/start ## Perform tests locally
	just local-env/start
	go test ./...
