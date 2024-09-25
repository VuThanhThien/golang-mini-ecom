# Product service
Install project:

```go
sh go_install.sh

make deps
```

create database:
```
make postgres

make migration-up
```

build swagger:
```
make docs
```



run local:

```
make dev
```