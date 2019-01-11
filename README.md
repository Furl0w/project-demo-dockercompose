### Simple docker compose for a project

## GO

I use intermediate image for building a single binary for each container and then serve the final app in a scratch container for minimum size and attackn surface. For serverDB I use dep for managing third party package (go driver for mongoDB)

## Python

Deployed in development mode (on host 0.0.0.0) in an alpine container

## Mongo

TODO : add a volume to store the data

Not directly accessible from the host (only reachable in the docker network).
If you want to reach directly the database you can either use :
 - docker exec -it CONTAINER_ID /bin/sh and then run mongo
 - add the lines under the mongo service to the docker compose to do a port mapping :
    - ports:
      - \- "27017:27017"

## Testing

Install docker.
Pull the repository and run $sudo docker-compose up --build

To ensure everything works :

- localhost:5000 from your browser should answer Hello world ! (flask)
- localhost:3030 from your browser should answer Hello world form path / (serverMain)
- localhost:3030/checkDB from your browser should answer with a successfull connection to the DB (serverMain + serverDB + Mongo)
- localhost:3030/testMobile return the json he has received from a post (test for mobile dev)

You can also run $sudo Caddy to enable the reverse proxy Caddy from the Caddyfile, in that case just replace the 3030 by 9090.
