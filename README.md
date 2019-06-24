# calc-grpc
Example gRPC server and client written in GO.

## Usage
A Makefile is present to run most commonly used commands. This project
uses docker and docker-compose to ease setup time and keep the
environment consistent.

The following commands are available:
- **make test**  - run all tests
- **make build** - compile .proto files and build executables
- **make run**   - run the gRPC server
- **make push**  - push a docker image with the currently build server
                   binary to docker hub

When running **make run** the port `50051` will be bound by default. To
change the port to bind to a `.env` file can be created in the project
root with the `PORT` variable.

## The Twelve-Factor App
Some care has been taken to adhere to
[The Twelve-Factor App](https://12factor.net/) methodology. By running
all commands in docker and using go mod to handle dependencies, the
project will behave consistently regardless of environment.
