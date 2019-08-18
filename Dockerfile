FROM golang:1.10 AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code dependencies and install it
WORKDIR $GOPATH/src/hades/
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

# Copy the source and build it
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./

EXPOSE 8080
ENTRYPOINT ["./app"]