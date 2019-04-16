# kakaotalk binctr

This idea is from `binctr`. Anyway this application run without rootless some reasons.e.g. Non root users to non root user mode in rootless is currently not acceptable.

## What is [binctr](https://github.com/genuinetools/binctr)?

Fully static, unprivileged, self-contained, containers as executable binaries.

## Build

```
sudo make kakaotalk
```

## Run

```
sudo ./kakaotalk
```

## Note

* You can only run this image in the linux host. Also you need to use UIM for
    user IM module.
* If you want to check base image which is `docker.io/leoh0/kakaotalk-wine-root`, then check this [dockerfile](https://github.com/leoh0/dockerfiles/tree/master/kakaotalk-wine-root) first.
