FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates
RUN apt-get -y update; apt-get -y install curl


# Add the binary
ADD ./bin/pokeapi /usr/bin/pokeapi

# Add the entrypoint script and give execute permissions
ADD ./remote/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Use entrypoint script to start service
# The script passes env variables as runtime flags to the program
ENTRYPOINT ["/entrypoint.sh"]
CMD []