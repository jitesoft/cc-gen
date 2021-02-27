# syntax=docker/dockerfile:experimental
FROM registry.gitlab.com/jitesoft/dockerfiles/alpine:latest
ARG VERSION
LABEL maintainer="Johannes Tegnér <johannes@jitesoft.com>" \
      maintainer.org="Jitesoft" \
      maintainer.org.uri="https://jitesoft.com" \
      com.jitesoft.project.repo.type="git" \
      com.jitesoft.project.repo.uri="https://gitlab.com/jitesoft/go/cc-gen" \
      com.jitesoft.project.repo.issues="https://gitlab.com/jitesoft/go/cc-gen/issues" \
      com.jitesoft.project.registry.uri="registry.gitlab.com/jitesoft/go/cc-gen" \
      com.jitesoft.project.app.cc-gen.version="${VERSION}"

ARG VERSION
ENV CC_GEN_VERSION=${VERSION}

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
