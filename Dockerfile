# syntax=docker/dockerfile:experimental
FROM registry.gitlab.com/jitesoft/dockerfiles/alpine:latest

ARG TARGETARCH
RUN --mount=type=bind,source=./bin,target=/tmp \
    addgroup -g 1000 ccgen \
 && adduser -u 1000 -G ccgen -s /bin/ash -D ccgen \
 && tar -xzf /tmp/cc-gen-${TARGETARCH}-linux.tar.gz -C /usr/local/bin \
 && chmod +x /usr/local/bin/cc-gen \
 && mkdir /workdir \
 && chown ccgen:ccgen /workdir

WORKDIR "/workdir"
USER ccgen
ENTRYPOINT ["cc-gen"]
