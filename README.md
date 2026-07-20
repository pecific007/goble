# Goble

It's a wrapper of [`bluetoothctl`](https://man.archlinux.org/man/bluetoothctl.1) (I can't write anything of my own).

### Usage:
```sh
$  ./goble
Usage of ./goble:
  -c	Connect to bluetooth device
  -d	Disconnect from bluetooth device
  -s	Connect to sister's earphones
```

``` sh
$ ./goble -c
Connected and trusted: 99:AA:BB:CC:DD:FF
```

```sh
$ ./goble -d
Disconnected and untrusted: 99:AA:BB:CC:DD:FF
```
