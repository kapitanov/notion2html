FROM golang:1.18-alpine As build
WORKDIR /app
ENV CGO_ENABLED=0
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
RUN go mod download
COPY . /app/
RUN go build -buildvcs=false -o=/out/notion2html ./cmd/notion2html

FROM alpine:latest
ENV LISTEN_ADDR=0.0.0.0:80
WORKDIR /opt/notion2html/
COPY docker-entrypoint.sh /docker-entrypoint.sh
COPY --from=build /out/notion2html /opt/notion2html/notion2html
COPY --from=build /app/templates /opt/notion2html/templates
ENTRYPOINT [ "/docker-entrypoint.sh" ]
