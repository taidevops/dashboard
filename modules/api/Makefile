ROOT_DIRECTORY = $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))/../..

.PHONY: codegen-update
codegen-update: --codegen-update-run --codegen-update-post

.PHONY: --codegen-update-run
--codegen-update-run:
  @${GOPATH}/bin/client-gen \
  --input