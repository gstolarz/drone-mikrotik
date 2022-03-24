FROM golang:1-alpine AS build
RUN apk add --no-cache upx
WORKDIR $GOPATH/src/github.com/gstolarz/drone-mikrotik
COPY . .
RUN CGO_ENABLED=0 go install -ldflags="-s -w" ./... \
  && upx --ultra-brute $GOPATH/bin/drone-mikrotik

FROM plugins/base:multiarch
LABEL maintainer="grzegorz.stolarz@gmail.com" \
  org.label-schema.name="Drone MikroTik Plugin" \
  org.label-schema.vendor="Grzegorz Stolarz" \
  org.label-schema.schema-version="1.0"
COPY --from=build /go/bin/drone-mikrotik /bin/drone-mikrotik
ENTRYPOINT ["/bin/drone-mikrotik"]
