.PHONY: prod
prod:
	$(eval ENV_NAME := prod)
	$(eval LISTEN_PORT := 8080)
	$(eval ALLOW_ORIGIN := https://d2semoivot4v0t.cloudfront.net)

	@# DB
	$(eval DB_SECRET_NAME := rds)

	@# AWS
	$(eval AWS_USER_POOL_ID := ap-northeast-1_3yXLWYIMZ)
	$(eval AWS_USER_POOL_CLIENT_ID := 79ruu0siuha5neotiggbfvaj3l)
