# C-icap exporter

Prometheus exporter for c-icap


## Run

    docker build -t olfeo/trustlane:c-icap-exporter-latest --platform linux/amd64 .
    docker run --rm -p 8080:8080 olfeo/trustlane:c-icap-exporter-latest
    docker push olfeo/trustlane:c-icap-exporter-latest

