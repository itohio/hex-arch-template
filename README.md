# hex-arch-template
This is a full-stack Go template for Hexagonal Architecture. 
Includes backend with different kinds of API, a few flavors of DB, Docker and docker-compose.
Front end is React with TypeScript using Apollo client for GraphQL consumption as well as Auth0 for authentication.
There is a simple frontend with Vecty.
The structure is inspired by [hexArchGoGRPC](https://github.com/selikapro/hexArchGoGRPC) repo as well as the [video](https://t.co/QaN1cAzDmu?amp=1) itself.

## Basic idea
The basic idea is to create business/application logic in the Application/Core layers.
Declare ports interfaces and arrange them to the left and right side.
The right side is the controlled systems by the Application while the left side is the controller of the Application.
Such things as APIs go to the left and DB, FS and external APIs/systems go to the right.
One important rule is that the core must not depend on Application, while Application must not depend on APIs or DBs.
The communication with external layers should go through ports.

# How to use
You can use this repo as a template for your project.
Just make sure to rename the app and Go module together with imports.

## Backend
1. Replace `hexarch` in imports and in `go.mod`
2. Groom left and right side adaptors
3. Modify application and core code to your liking along with respective ports
4. Update `main.go` to import ant instantiate only what is needed

## Frontend
### React
1. Replace App name, title, etc
2. Add items to the router and header
3. Add components and GraphQL queries/mutations
3.1. Add custom React hooks `use*` as shown in the example
3.2. run `npm run generate` to update query/mutation types

### Vecty
Vecty is a React-like Web frontend framework that is written completely in Go.

1. In the main.go update Auth instantiation variables (Audience, ClientID, Domain).
2. Update sitemap, pages, etc.
3. update `*.graphql` files and run `go generate` to update GraphQL client.

NOTE: It is important to have `AuthPre` and `AuthPost` incapsulating `spa` so that redirects are working properly.
Components cannot return `mdc` components directly - weird errors happen randomly.

IMPORTANT NOTE: Authentication component and `Auth0` was not battle tested. Use on your own risk.

### Fyne
Fyne is a cross-platform GUI framework based on Material Design.

## Authentication
1. Register on Auth0
2. Setup your App and API
3. Configure respective environment variables in `.env`
