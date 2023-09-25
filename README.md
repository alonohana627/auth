# Auth - Golang's Basic Microservice for User Authentication

A simple Protobuf gRPC authentication microservice. The service is very simple, and intuitive. The project was started
in a weekend, so it has a lot of room for improvement.

The microservice uses Postgresql for storage of the information.

## Routes

Currently, there are four routes:

- **Register** - gets a user model and registers to the Database.
- **Login** - gets email and password and verifies them. Generates JWT token accordingly.
- **RefreshToken** - gets a token and refreshes it if the token is valid.
- **ValidToken** - checks whether the JWT token is valid or not.

## Database Schema

A simple schema of a user:

- Email: string, unique
- Nickname: string unique,
- HashedPassword: string

The rest is a Gorm model

## Build and Run

Refer to the docker-compose.yaml. Simply put `docker-compose up` in the closest shell to your home and the microservice
will be up and running. For autocomplete, I have a script for generating the protobuf, you can change the settings for
the db connection so it will take default params for developing outside of Docker.

## What's Next?

Now I have a basic auth microservice that will verify in a blazing fast speed the authentication needed. This is built
as a module that can be integrated into any Docker-like environment.

## This is a Toy Project!

Just remember - you can use snippets from here to understand what is going on, but do not use it for production.