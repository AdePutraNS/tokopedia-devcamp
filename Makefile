# go build command
gobuild:
	@go build -v -o generated_apps cmd/generated_apps/*.go

# go run command
gorun:
	make gobuild
	@./generated_apps --config_path='./config/tkp-app.{TKPENV}.yaml'

# deploy command
deploy:
	@echo "deploying"
