FROM golangci/golangci-lint:v1.27-alpine AS build
RUN apk add git
ENV CGO_ENABLED=0
RUN apk --no-cache add ca-certificates

WORKDIR /src
COPY go.* .
RUN go mod download

COPY . .
RUN golangci-lint run --timeout 10m0s ./... \
  && go test -v ./... \
  && go build -o /out/api .

FROM scratch
# copy the ca-certificate.crt from the build stage
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/api /

EXPOSE 3001
STOPSIGNAL SIGINT

ENTRYPOINT [ "/api" ]