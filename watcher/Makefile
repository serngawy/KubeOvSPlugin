default: build

build:
	docker build -t watcher .

run: build
        docker run -it -e odl_ip_adress=http://127.0.0.1:8181 -e user_value=admin -e password_value=admin --network=host -v /home/ubuntu/.kube/:/root/.kube/  watcher
