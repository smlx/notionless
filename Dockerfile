FROM alpine:3.24@sha256:a2d49ea686c2adfe3c992e47dc3b5e7fa6e6b5055609400dc2acaeb241c829f4
ARG BINARY=binary-build-arg-not-defined
ENV BINARY=${BINARY}
ENTRYPOINT ["sh", "-c"]
CMD ["exec /${BINARY}"]
# TARGETPLATFORM is defined by goreleaser during the build
ARG TARGETPLATFORM
COPY ${TARGETPLATFORM}/${BINARY} /
