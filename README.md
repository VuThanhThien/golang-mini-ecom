Generate SSL certificates (Optional)

> If you don't SSL now, change `SSL=TRUE` to `SSL=FALSE` in the `.env` file

```
$ mkdir cert/
```

```
$ sh generate-certificate.sh
```

if you have issue with $PATH of zsh, try :

```
export PATH=$(go env GOPATH)/bin:$PATH
```

build swagger:
```
make docs
```
