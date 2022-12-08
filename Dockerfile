FROM golang:1.19 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go generate ./...
RUN CGO_ENABLED=0 go build -o bin/service ./
RUN strip bin/service

FROM debian:bullseye

ENV cicapBaseVersion="0.5.10" cicapModuleVersion="0.5.5"
RUN apt-get -y update && apt-get -y install curl tar build-essential
RUN	curl --silent --location --remote-name "https://sourceforge.net/projects/c-icap/files/c-icap/0.5.x/c_icap-${cicapBaseVersion}.tar.gz" && \
	tar -xzf "c_icap-${cicapBaseVersion}.tar.gz" && cd c_icap-${cicapBaseVersion} && \
	./configure --enable-large-files --prefix=/usr/local/c-icap && make && make install

RUN adduser --system --disabled-password --disabled-login --group c-icap --uid 953
RUN chown -R c-icap:c-icap /usr/local/c-icap


RUN mkdir /app
RUN adduser --system --disabled-password --disabled-login --group service
RUN chown -R service:service /app
USER service
WORKDIR /app
EXPOSE 8080
COPY --from=builder /app/bin/service /app/bin/service
ENTRYPOINT ["/app/bin/service"]
