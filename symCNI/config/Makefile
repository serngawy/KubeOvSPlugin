default: build

build:
docker build . -t  sym-cni:v1.0
# This directory /home/ubuntu/.kube  may need to be change based on the deployment env
run: build
sudo docker run --network=host -it  -v /home/ubuntu/.kube:/config/.kube -v /etc/cni/net.d/:/etc/cni/net.d/ -v /opt/cni/bin/:/opt/cni/bin/ ovs-cni:v1.0




