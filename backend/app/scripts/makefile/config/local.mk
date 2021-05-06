.PHONY: local
local:
	$(eval ENV_NAME := local)
	$(eval LISTEN_PORT := 8080)
	$(eval ALLOW_ORIGIN := http://localhost:3000)

	@# DB
	$(eval DB_USER := root)
	$(eval DB_HOST := 127.0.0.1)
	$(eval DB_PORT := 3306)
	$(eval DB_PASSWORD := mysql)
	$(eval DB_NAME := todo)

	@# AWS
	$(eval AWS_USER_POOL_ID := ap-northeast-1_axGAoF8m8)
	$(eval AWS_USER_POOL_CLIENT_ID := 6u3mjd9gmvo638193q9ci4lmi6)
