FROM golang:1.15.3-alpine3.12 as builder
WORKDIR /go/app

COPY go.mod go.sum ./
RUN go mod download -x
COPY . .

ARG VERSION
ARG BUILD_TIME
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X 'main.buildVersion=${VERSION}' -X 'main.buildTime=${BUILD_TIME}'" -o /app/virtual-kubelet ./cmd/provider

# For production
# FROM gcr.io/distroless/base

# For debugging
FROM alpine 

ENV APISERVER_CERT_LOCATION /vkubelet-browser-crt.pem
ENV APISERVER_KEY_LOCATION /vkubelet-browser-key.pem
ENV KUBELET_PORT 10250
ENV API_PORT 80

EXPOSE 10250
EXPOSE 80

# Copy the configuration file for the mock provider.
COPY hack/skaffold/virtual-kubelet/vkubelet-browser-cfg.json /vkubelet-browser-cfg.json
# Copy the certificate for the HTTPS server.
COPY hack/skaffold/virtual-kubelet/vkubelet-browser-crt.pem /vkubelet-browser-crt.pem
# Copy the private key for the HTTPS server.
COPY hack/skaffold/virtual-kubelet/vkubelet-browser-key.pem /vkubelet-browser-key.pem

CMD ["/virtual-kubelet"]
COPY --from=builder /app .
ADD frontend /browser-kube-files
