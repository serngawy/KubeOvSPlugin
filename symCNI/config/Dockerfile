FROM alpine:latest
# Install kubectl
ADD https://storage.googleapis.com/kubernetes-release/release/v1.12.0/bin/linux/amd64/kubectl /usr/local/bin/kubectl
RUN set -x && \
    apk add --no-cache curl ca-certificates && \
    chmod +x /usr/local/bin/kubectl
RUN mkdir -p /root/symcni/
RUN mkdir -p /root/.kube/
RUN mkdir -p /opt/cni/bin/
RUN mkdir -p /etc/cni/net.d/
WORKDIR /root/symcni/
COPY ./symOvSCNI /root/symcni/
COPY ./symcni.sh /root/symcni/
RUN chmod 775 /root/symcni/*
ENTRYPOINT ["./symcni.sh"]
