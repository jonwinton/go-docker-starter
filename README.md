# Go Docker Starts
> A quick Dockerized Go project starter

## Features

- Run a Go server without installing any packages locally
- Uses [`realize`](https://github.com/tockins/realize) to handle in container updates to code
- Local dev specific Dockerfile

## Starting

- To build the image the very first time run `make build`.
- To start the server just run `make`
- The server will spin up on whatever port is defined in your dev env file (`docker/dev.env`)
