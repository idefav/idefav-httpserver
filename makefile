All:
HUB ?=idefav
IMAGE ?=httpserver:latest
build: All
	./build.sh
push: build
	docker push "${HUB}/${IMAGE}"
run: All
	docker run -p 8080:8080 ${HUB}/${IMAGE}