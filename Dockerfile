FROM golang:1.20 as build
WORKDIR /go/src/app
ARG BUILD_TARGET
COPY go.mod go.sum ./
RUN go mod download
COPY db db
COPY lib lib
COPY $BUILD_TARGET $BUILD_TARGET
RUN CGO_ENABLED=0 go build -o /go/bin/app github.com/Goldziher/go-monorepo/$BUILD_TARGET
RUN chmod +x /go/bin/app

FROM gcr.io/distroless/static-debian11
COPY --from=build /go/bin/app /
CMD ["/app"]
