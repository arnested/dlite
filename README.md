# DLite

this branch represents the latest *beta* version. the stable version can be found in the legacy branch.

## Building

install dependencies

```sh
brew install golang libev
go get -u github.com/jteeuwen/go-bindata/...
git submodule init
git submodule update
```

update dependencies (use this if you've already built the project before)

```sh
git submodule foreach git pull origin master
```

build the binary

```sh
go generate
go build
```
