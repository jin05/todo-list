.PHONY: test
test:
	$(eval ENV_NAME := test)
	$(eval RICHIGO := $(shell which richgo > /dev/null; echo $$?))
	@if [ $(RICHIGO) = 0 ]; then richgo test ./...; else go test -v ./...; fi

.PHONY: install_richgo
install_richgo:
	# only osx
	brew tap kyoh86/tap
	brew install richgo

.PHONY: install_mockgen
install_mockgen:
	go get github.com/golang/mock/mockgen

.PHONY: gen_mock
gen_mock:
	sh scripts/sh/test/generate_mock.sh
