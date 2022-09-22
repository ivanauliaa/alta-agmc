## Docker Image

Image: https://hub.docker.com/repository/docker/ivanauliaa/alta-agmc

Pull:
```
docker pull ivanauliaa/alta-agmc:atlas-1.0
```

Run:

```
docker run -d \
-p 5000:5000 \
-e APP_PORT=5000 \
-e APP_ENV=production \
-e DB_USER=admin \
-e DB_PASS=supersecret \
-e DB_HOST=<> \
-e DB_PORT=3306 \
-e DB_NAME=agmc_hexagonal \
-e ACCESS_TOKEN_SECRET=1234 \
-e MONGO_USER=root \
-e MONGO_PASS=root \
-e MONGO_HOST=<> \
-e MONGO_PORT=27017 \
-e MONGO_NAME=agmc_hexagonal \
--name alta-agmc-app-1 \
ivanauliaa/alta-agmc:atlas-1.0
```

## API

1. Create user
![](screenshots/Screenshot%20(154).png)

2. Login user
![](screenshots/Screenshot%20(155).png)

3. Collection variables
![](screenshots/Screenshot%20(156).png)

4. Get users
![](screenshots/Screenshot%20(157).png)

5. Get user by id
![](screenshots/Screenshot%20(158).png)

6. Update user by id
![](screenshots/Screenshot%20(159).png)

![](screenshots/Screenshot%20(160).png)

7. Create book
![](screenshots/Screenshot%20(161).png)

8. Get books
![](screenshots/Screenshot%20(162).png)

9. Get book by id
![](screenshots/Screenshot%20(163).png)

10. Update book by id
![](screenshots/Screenshot%20(164).png)

![](screenshots/Screenshot%20(165).png)

11. Delete book by id
![](screenshots/Screenshot%20(166).png)

12. Delete user by id
![](screenshots/Screenshot%20(167).png)
