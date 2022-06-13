RUN_HUGO_CONTAINER_ARGS := --rm \
  --mount type=bind,source="$(shell pwd)",target=/w \
  --workdir /w \
  jojomi/hugo:0.76.5
POSTS_ROOT := ./content/posts
YEAR := $(shell date '+%Y')

hugo_build:
	docker run $(RUN_HUGO_CONTAINER_ARGS) hugo

hugo_new:
	$(eval FILE_NAME := $(shell read -p "File Name (Up to 64 characters a-z, 0-9 or -): " input && if test "$$input" = ""; then input=unknown; fi && echo $$input))
	mkdir -p '$(POSTS_ROOT)/$(YEAR)'
	docker run $(RUN_HUGO_CONTAINER_ARGS) hugo new '$(POSTS_ROOT)/$(YEAR)/$(FILE_NAME).md'

hugo_server:
	docker run --publish 1313:1313 $(RUN_HUGO_CONTAINER_ARGS) hugo server --bind 0.0.0.0
