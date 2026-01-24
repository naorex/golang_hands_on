# Go 言語　ハンズオン

## Docker

```shell
docker build ./ -t golang-hands-on
```

```shell
docker container run -itd --name my-golang golang-hands-on
```

## PATH edit

```shell
echo 'export PATH=$PATH:/usr/local/go/bin' >> .bash_profile
source .bash_profile
```
