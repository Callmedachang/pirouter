# Pirouter
Pirouter(π-router) a toy-level http router

## Usage
```shell
go get github.com/willxm/pirouter
```

```golang
func main() {
	r := pirouter.NewRouter()
	r.Register("POST", "/user/login", greet)
	r.Register("GET", "/user/info", hix)
	r.Register("GET", "/user/photo", hiy)
	r.Run(":8080")
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! %s", time.Now())
}

func hix(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi x! %s", time.Now())
}

func hiy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi y! %s", time.Now())
}

```

## TODO
- error handler
- middleware