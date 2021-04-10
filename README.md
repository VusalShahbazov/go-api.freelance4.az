

###migration tool 

```
migrate -path ./migrations  -database "mysql://root:@tcp(localhost:3306)/freelance4?multiStatements=true" up
```


````
migrate create -ext sql -dir ./database/migrations <migration_name>
````

```
go build cmd/api/main.go && main.exe
```