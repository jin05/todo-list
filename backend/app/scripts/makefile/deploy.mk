.PHONY: deploy
deploy:
	@if [ "$(ENV_NAME)" != "prod" ]; then echo "should ENV (prod)"; exit 1; fi
	sls deploy --verbose
