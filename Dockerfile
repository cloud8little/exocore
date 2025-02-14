# This is the published docker image for exocore.

FROM golang:1.21.12-alpine3.19 AS build-env

WORKDIR /go/src/github.com/ExocoreNetwork/exocore

COPY go.mod go.sum ./

RUN apk add --no-cache ca-certificates=20241121-r1 build-base~=0.5 git~=2.43 linux-headers~=6.5

RUN --mount=type=bind,target=. --mount=type=secret,id=GITHUB_TOKEN \
    git config --global url."https://$(cat /run/secrets/GITHUB_TOKEN)@github.com/".insteadOf "https://github.com/"; \
    go mod download

COPY . .

RUN make build && go install github.com/MinseokOh/toml-cli@latest

FROM alpine:3.19

WORKDIR /root

COPY --from=build-env /go/src/github.com/ExocoreNetwork/exocore/build/exocored /usr/bin/exocored
COPY --from=build-env /go/bin/toml-cli /usr/bin/toml-cli

RUN apk add --no-cache \
	ca-certificates=20241121-r1 \
	libstdc++~=13.2 \
	jq~=1.7 \
	curl~=8.12 \
	bash~=5.2 \
    && addgroup -g 1000 exocore \
    && adduser -S -h /home/exocore -D exocore -u 1000 -G exocore

USER 1000
WORKDIR /home/exocore

EXPOSE 26656 26657 1317 9090 8545 8546

# Every 30s, allow 3 retries before failing, timeout after 30s.
HEALTHCHECK --interval=30s --timeout=30s --retries=3 CMD curl -f http://localhost:26657/health || exit 1

CMD ["exocored"]
