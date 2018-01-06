# manpowersupply-backend

## Prerequisite

- Golang
- PostgreSQL

## Setup Database

- Create a databse named `civimech`
- Run following commands
    - Status: `goose -dir ./db/migrations/ postgres "user=postgres dbname=civimech sslmode=disable password=C0mplexPwd" status`
    - Upgrade: `goose -dir ./db/migrations/ postgres "user=postgres dbname=civimech sslmode=disable password=C0mplexPwd" up`
    - Downgrade: `goose -dir ./db/migrations/ postgres "user=postgres dbname=civimech sslmode=disable password=C0mplexPwd" up`