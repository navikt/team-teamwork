FROM golang:1.22 as builder

WORKDIR /src

COPY go.sum go.sum
COPY go.mod go.mod
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /src/app

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /src/app /
CMD ["/app"]
