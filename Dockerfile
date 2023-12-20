FROM golang:1.21 as build-go
WORKDIR /src
COPY . /src
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux
RUN go build ./cmd/server/*.go

FROM scratch
COPY --from=build-go /src/server /
CMD ["/server"]
EXPOSE 80
