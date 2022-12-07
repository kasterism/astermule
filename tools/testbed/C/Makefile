IMG ?= kasterism/test_c
docker-build:
	docker buildx build -t ${IMG} . --load

docker-push:
	docker buildx build --platform linux/amd64,linux/arm64 -t ${IMG} . --push