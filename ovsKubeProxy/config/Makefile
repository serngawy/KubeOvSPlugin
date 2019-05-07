default: build

build:
	docker build . -t symkubeproxy:v1.1

# This directory /home/vagrant/.kube/ may need to be change based on the deployment env
run: build
	docker run --network=host -v /home/vagrant/.kube/:/root/.kube/ -v /var/run/openvswitch/:/var/run/openvswitch/ symkubeproxy:v1.1
