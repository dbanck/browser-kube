FROM gcr.io/distroless/base

ENV APISERVER_CERT_LOCATION /vkubelet-browser-crt.pem
ENV APISERVER_KEY_LOCATION /vkubelet-browser-key.pem
ENV KUBELET_PORT 10250

# Use the pre-built binary in "bin/virtual-kubelet".
COPY bin/e2e/virtual-kubelet /virtual-kubelet
# Copy the configuration file for the mock provider.
COPY hack/skaffold/virtual-kubelet/vkubelet-browser-cfg.json /vkubelet-browser-cfg.json
# Copy the certificate for the HTTPS server.
COPY hack/skaffold/virtual-kubelet/vkubelet-browser-crt.pem /vkubelet-browser-crt.pem
# Copy the private key for the HTTPS server.
COPY hack/skaffold/virtual-kubelet/vkubelet-browser-key.pem /vkubelet-browser-key.pem

CMD ["/virtual-kubelet"]
