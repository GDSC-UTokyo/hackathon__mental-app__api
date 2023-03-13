FROM golang:1.19 as build
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download

ADD . /cmd
WORKDIR /cmd
RUN go build -o bin ../cmd/app/main.go

FROM gcr.io/distroless/base-debian11
COPY --from=build /cmd/bin /app
EXPOSE 8080
CMD ["/app"]