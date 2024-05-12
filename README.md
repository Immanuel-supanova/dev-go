# dev-go


```
// .env

DB="host=host user=user password=pass dbname=dbname port=port sslmode=disable"
DEBUG=<Either true or false>
```

```
// main.go

func init() {
	devgo.Config()
	devgo.ConnectToDb()
	devgo.SyncDatabase()
}

...
func main() {
	devgo.DevGoServer.ListenAndServe()
}
```