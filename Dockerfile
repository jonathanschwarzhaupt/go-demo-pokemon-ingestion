FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates
RUN apt-get -y update; apt-get -y install curl


ADD pokeapi /usr/bin/pokeapi

CMD ["pokeapi"]
