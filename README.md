if you have issue with $PATH of zsh, try :

```
export PATH=$(go env GOPATH)/bin:$PATH
```


build swagger:
```
make docs
```

run local:

```
make dev
```