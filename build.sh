hub=${HUB-idefav}
image=${IMAGE-httpserver:latest}
docker build . -t "${hub}/${image}"