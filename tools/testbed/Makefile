.PHONY: docker-build
docker-build:
	$(MAKE) -C ./A docker-build
	$(MAKE) -C ./B docker-build
	$(MAKE) -C ./C docker-build
	$(MAKE) -C ./D docker-build

.PHONY: docker-push
docker-push:
	$(MAKE) -C ./A docker-push
	$(MAKE) -C ./B docker-push
	$(MAKE) -C ./C docker-push
	$(MAKE) -C ./D docker-push