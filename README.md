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
Connected and trusted: AA:BB:CC:DD:FF:GG%
```

```sh
$ ./goble -d
Attempting to disconnect from AA:BB:CC:DD:FF:GG
[CHG] Device AA:BB:CC:DD:FF:GG ServicesResolved: no
[SIGNAL] BREDR.Disconnected - org.bluez.Reason.Remote, Connection terminated by remote user
[SIGNAL] Disconnected - org.bluez.Reason.Remote, Connection terminated by remote user
[CHG] Device AA:BB:CC:DD:FF:GG Connected: no
Disconnection successful
[CHG] Device AA:BB:CC:DD:FF:GG Trusted: no
Changing AA:BB:CC:DD:FF:GG untrust succeeded
Disconnected and untrusted: AA:BB:CC:DD:FF:GG%
```
