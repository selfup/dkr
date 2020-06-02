# dkr

Docker for Docker containers?

This project was inspired by [Liz Rice](https://www.lizrice.com/)

### deps

1. go ([golang](https://golang.org/))
1. `./scripts/fs.sh` (this pulls minirootfs)

### build

`./scripts/build.sh`

### boot container host

`docker-compose up --build`

### run dkr in container host

`docker-compose exec --privileged dkr ./dkr run /bin/sh`

Here you will be in a shell.

Example of hostname manipulation of dkr container instead of the docker container:

```
dkr $ docker-compose exec --privileged dkr ./dkr run /bin/sh
2020/05/27 23:49:24 run() executing [/bin/sh]
2020/05/27 23:49:24 fork() executing [/bin/sh]
/ # hostname
f60818cfd391
/ # hostname dkr
/ # hostname
dkr
/ # exit
dkr $ docker-compose exec --privileged dkr ./dkr run /bin/sh
2020/05/27 23:51:17 run() executing [/bin/sh]
2020/05/27 23:51:17 fork() executing [/bin/sh]
/ # hostname
f60818cfd391
/ #
```

See how the running `docker-compose up --build` continainer still has the same hostname?

Yea we did it!!

### linux

If you are already on linux just:

```bash
# pulls minitrootfs
./scripts/fs.sh
# builds dkrd and dkr linux bins
./scripts/build.sh

sudo MINIROOTFS_HOME=$(pwd) ./dkr /bin/bash
```

### license

This project is MIT licensed. Please check the LICENSE file.
