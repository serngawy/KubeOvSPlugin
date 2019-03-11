FROM alpine:latest
RUN mkdir -p /root/kubeproxy/
RUN mkdir -p /root/.kube/
RUN mkdir -p /var/run/openvswitch/
RUN mkdir -p /etc/cni/net.d/
WORKDIR /root/kubeproxy/
COPY ./symkubeproxy /root/kubeproxy/
ENTRYPOINT ["sh","-c","/root/kubeproxy/symkubeproxy","$confile"]