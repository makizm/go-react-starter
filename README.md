# Monorepo

A monorepo containing a Go backend API server and React web UI.

## Project Structure

```
monorepo/
├── docs/             # Swagger (swag) generated spec (json/yaml)
├── server/           # Go API server
│   ├── cmd/          # Application entrypoints
│   │   └── api/      # API server main package
│   ├── internal/     # Private application code
│   │   ├── handlers/ # HTTP handlers
│   │   ├── models/   # Data models
│   │   └── services/ # Business logic
│   └── pkg/          # Public packages
├── ui/               # React web application
│   ├── public/       # Static assets
│   └── src/          # React source code
```

## Prerequisites

- Go 1.25+
- [swag](https://github.com/swaggo/swag) - OpenAPI (Swagger) generator
- Node.js 22+ (LTS) - use [fnm](https://github.com/Schniz/fnm) or [nvm](https://github.com/nvm-sh/nvm) to manage versions
- Docker or Podman for running full stack app with live reload
- GNU make

### Generating Swagger docs

Swagger docs are generated from Go annotations using `swag`.

```sh
make docs
```

## Getting Started

To start a full stack app for development simply run `docker compose up`.

By default, the app runs on `http://localhost:5173`.

- http://localhost::5173/ (ui)
- http://localhost:5173/api (api)
- http://localhost:5173/api/docs (swagger api docs)

The port can be changed suing compose env `UI_PORT` https://docs.docker.com/compose/how-tos/environment-variables/variable-interpolation/#env-file

## License

MIT
