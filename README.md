# nuts-registry-admin-demo
Demo application which shows how to integrate with the Nuts registry.

## Building and running
### Production
To build for production:

```shell
$ npm install
$ npm run build
$ go run main.go
```

This will serve the front end from the embedded filesystem.
### Development

During front-end development, you probably want to use the real filesystem and webpack in watch mode:

```shell
$ npm install
$ npm run watch
$ go run main.go live
```

The API is generated from the `api/api.yaml`.
```shell
$ oapi-codegen -generate types,server -package api api/api.yaml > api/generated.go
```

### Docker
```shell
$ docker run -p 1303:1303 nutsfoundation/nuts-registry-admin-demo
```

## Configuration
When running in Docker without a config file mounted at `/app/server.config.yaml` it will use the default configuration.
In this case the default username will be `demo@nuts.nl`. The password is generated and printed in the log on startup.