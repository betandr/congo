FROM --platform=${BUILDPLATFORM} golang:1.16.4-alpine3.13 AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* .
RUN go mod download
COPY . .

FROM base as build
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/server .

FROM base as unit-test
RUN go test -v .

FROM scratch AS bin-unix
COPY --from=base /out/server /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=base /out/server /server.exe

FROM bin-${TARGETOS} AS bin