# Stage 1 - Build the binary
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /conapp ./cmd/main.go

# Stage 2 - Deploy the binary
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /conapp /conapp

EXPOSE 4000

USER nonroot:nonroot

ENTRYPOINT [ "/conapp"]
