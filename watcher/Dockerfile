FROM alpine:latest
RUN mkdir -p /watcher
COPY ./watcher  /watcher
WORKDIR  /watcher
ENTRYPOINT ./watcher odl --host $odl_ip_adress  --username $user_value --password $password_value

