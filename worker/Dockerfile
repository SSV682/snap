FROM golang:1.19.2-alpine as base
WORKDIR /build
COPY go.* ./
RUN go mod download
COPY . .

FROM base as build
ENV OS linux
ENV ARCH amd64
COPY --from=base /go/pkg /go/pkg
COPY . /app
RUN CGO_ENABLED=0 GOOS=$OS GOARCH=$ARCH go build -o /bin/app ./cmd

FROM scratch
COPY --from=build /app/config /config
COPY --from=build /bin/app /app

CMD ["/app"]
