.ONESHELL:
.PHONY: help env
ENV_FILE=.env
help:

# only for the subshell
env:
	@export $(egrep -v '^#' $(ENV_FILE) | xargs)
