.PHONY: deploy
deploy:
	go run ./cmd/deploy/main.go -build="hugo" -output="public" -branch="live"
