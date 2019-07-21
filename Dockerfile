FROM golang:1.12.7-alpine3.10 AS builder
COPY webinfo.go .
RUN CGO_ENABLED=0 go build -o /webinfo webinfo.go

FROM scratch
USER 1000
COPY --from=builder /webinfo /webinfo
EXPOSE 8080
ENTRYPOINT ["/webinfo"]