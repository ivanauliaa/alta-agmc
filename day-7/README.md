Image: https://hub.docker.com/repository/docker/ivanauliaa/alta-agmc

Pull:
```
docker pull ivanauliaa/alta-agmc:1.0
```

Run:

```
docker run -d `
-p 5000:5000 `
-e APP_PORT=5000 `
-e APP_ENV=production `
-e DB_USER=root `
-e DB_PASS=root `
-e DB_HOST=localhost `
-e DB_PORT=3306 `
-e DB_NAME=agmc_hexagonal `
-e ACCESS_TOKEN_SECRET=1234 `
-e MONGO_USER=root `
-e MONGO_PASS=example `
-e MONGO_HOST=localhost `
-e MONGO_PORT=27017 `
-e MONGO_NAME=agmc_hexagonal `
--name alta-agmc-app-1 `
--net host `
--add-host host.docker.internal:host-gateway `
ivanauliaa/alta-agmc:1.0
```