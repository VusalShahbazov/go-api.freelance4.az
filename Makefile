run:
	go build cmd/api/main.go
	main

build:
	go build cmd/api/main.go

migration:
	migrate create -ext sql -dir ./database/migrations $(name)

migrate_up:
	migrate -path ./database/migrations  -database "mysql://root:@tcp(localhost:3306)/freelance4?multiStatements=true" up

migrate_down:
	migrate -path ./database/migrations  -database "mysql://root:@tcp(localhost:3306)/freelance4?multiStatements=true" down

migrate_fresh:
	migrate -path ./database/migrations  -database "mysql://root:@tcp(localhost:3306)/freelance4?multiStatements=true" down
	migrate -path ./database/migrations  -database "mysql://root:@tcp(localhost:3306)/freelance4?multiStatements=true" up

