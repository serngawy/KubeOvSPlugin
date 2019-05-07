FROM alpine:latest
# Install kubectl
ADD https://storage.googleapis.com/kubernetes-release/release/v1.12.0/bin/linux/amd64/kubectl /usr/local/bin/kubectl
RUN set -x && \
    apk add --no-cache curl ca-certificates && \
    chmod +x /usr/local/bin/kubectl
RUN mkdir -p /root/ovscni/
RUN mkdir -p /root/.kube/
RUN mkdir -p /opt/cni/bin/
RUN mkdir -p /etc/cni/net.d/
WORKDIR /root/ovscni/
COPY ./OvSCNI /root/ovscni/
COPY ./ovscni.sh /root/ovscni/
RUN chmod 775 /root/ovscni/*
ENTRYPOINT ["./ovscni.sh"]
