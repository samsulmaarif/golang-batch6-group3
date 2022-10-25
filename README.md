# golang-batch6-group3

DB dengan docker

```
docker volume create golang-db

docker run -dti \
	--name postgres \
	-e POSTGRES_PASSWORD=rahasia \
    -e POSTGRES_USER=nest \
    -e POSTGRES_DB=golangnest \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
	-v golang-db:/var/lib/postgresql/data \
    -p 5432:5432 \
	postgres
```