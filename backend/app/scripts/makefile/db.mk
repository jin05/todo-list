####################################
# migration
####################################
db_install_schemalex:
	$(eval UNAME := $(shell uname -s | tr '[A-Z]' '[a-z]'))
	$(eval SCHEMALEX_FILE := schemalex_$(UNAME)_amd64)
	$(eval SCHEMALEX_EXT := $(shell if [ "$(UNAME)" = "darwin" ]; then echo 'zip'; else echo 'tar.gz'; fi))
	curl -L -O https://github.com/schemalex/schemalex/releases/download/v0.1.1/$(SCHEMALEX_FILE).$(SCHEMALEX_EXT)
	if [ "$(SCHEMALEX_EXT)" = "zip" ]; then \
		unzip $(SCHEMALEX_FILE).$(SCHEMALEX_EXT); \
	else \
		tar xvzf $(SCHEMALEX_FILE).$(SCHEMALEX_EXT); \
	fi
	mv $(SCHEMALEX_FILE)/schemalex /usr/local/bin/schemalex
	rm -f $(SCHEMALEX_FILE).$(SCHEMALEX_EXT)
	rm -rf $(SCHEMALEX_FILE)

db_create_sql:
	go run db/create_ddl/create_ddl.go

# should ENV
db_schemalex:
	schemalex "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" db/sql/master.sql > db/sql/diff.sql

# should ENV
db_migrate: db_schemalex
	mysql -h$(DB_HOST) -P$(DB_PORT) -u$(DB_USER) -p$(DB_PASSWORD) $(DB_NAME) < db/sql/diff.sql
