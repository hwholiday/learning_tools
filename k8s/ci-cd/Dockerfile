FROM golang:1.15.6 as build
COPY . /build/
WORKDIR /build
# RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o app .

FROM scratch
COPY --from=build /build/app /
ENTRYPOINT ["/app"]