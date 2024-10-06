FROM --platform=$BUILDPLATFORM golang:alpine AS builder

RUN mkdir /app
WORKDIR /app

# RUN --mount=type=cache,target=/var/cache/apt \
#     apt-get update && apt-get install -y build-essential

# ENV CGO_ENABLED=0 \
#     GOPATH=/go \
#     GOCACHE=/go-build

COPY go.* .

RUN go mod tidy
RUN go mod verify
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download

COPY . .

# RUN --mount=type=cache,target=/go/pkg/mod/cache \
#     --mount=type=cache,target=/go-build \
#     go build -o bin/exe main.go

# CMD ["/app"]

# FROM builder as dev-envs

# RUN <<EOF
# apk update
# apk add git
# EOF

# RUN <<EOF
# addgroup -S docker
# adduser -S --shell /bin/bash --ingroup docker vscode
# EOF

# install Docker tools (cli, buildx, compose)
# COPY --from=gloursdocker/docker / /

CMD ["go", "run", "main.go"]

# FROM scratch
# COPY --from=builder /app/bin/exe /usr/local/bin/exe

# CMD ["/usr/local/bin/exe"]