All:
HUB ?=idefav
IMAGE ?=httpserver:0.0.1
build: All
	./build.sh
push: build
	docker push "${HUB}/${IMAGE}"
run: All
	docker run -p 8080:8080 ${HUB}/${IMAGE}