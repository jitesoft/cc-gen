# syntax = docker/dockerfile:experimental
ARG BASE_IMAGE="scratch"
FROM $BASE_IMAGE

ARG TARGETARCH
ADD /bin/cc-gen-${TARGETARCH} /usr/bin/cc-gen

VOLUME [ "/data" ]
WORKDIR /data

ENTRYPOINT [ "cc-gen" ]
CMD ["help"]
