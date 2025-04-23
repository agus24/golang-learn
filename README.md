install these:
```sh
brew install golang-migrate

```

create migration

```sh
migrate create -ext sql -dir db/migrations -seq create_users_table
```
