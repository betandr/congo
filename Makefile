all: build
.PHONY: build clean test run

PLATFORM=local

build:
	@docker build --target bin \
	--output bin/ \
	--platform ${PLATFORM} .

test:
	@docker build . --target unit-test

clean:
	rm ./bin/*
	rmdir bin
