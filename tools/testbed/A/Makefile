IMG ?= kasterism/test_a
docker-build:
	docker buildx build -t ${IMG} . --load

docker-push:
	docker buildx build --platform linux/amd64,linux/arm64 -t ${IMG} . --push