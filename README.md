# Migrations

## To generate a new migration, run the following command:

```bash
migrate create -ext sql -dir ./config/db/migrations -seq <migration_name>
```

## To execute the migrations, run the following command:

```bash
migrate -database "mysql://root:password@tcp(localhost:3306)/xikiturl" -path ./config/db/migrations up
```