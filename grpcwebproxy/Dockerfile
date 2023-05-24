FROM alpine:3.12.0

# Install `grpcwebproxy` and `openssl`.
COPY grpcwebproxy /go/bin/grpcwebproxy

# Copy the SSL certificates.
COPY ./fullchain.pem /etc/ssl/fullchain.pem
COPY ./privkey.pem /etc/ssl/privkey.pem

# Start the `grpcwebproxy` binary as the main process.
ENTRYPOINT ["/go/bin/grpcwebproxy"]

CMD ["--backend_addr=qrpay-be:50051", \
    "--server_tls_cert_file=/etc/ssl/fullchain.pem", \
    "--server_tls_key_file=/etc/ssl/privkey.pem", \
    "--backend_tls=true", \
    "--allow_all_origins", \
    "--enable_request_debug", \
    "--server_http_max_read_timeout=0s", \
    "--server_http_max_write_timeout=0s", \
    "--backend_tls_noverify", \
    "--server_http_tls_port=443"]