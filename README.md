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
Attempting to connect to 8C:D3:D9:F7:62:D0
Connection successful
[CHG] Device 8C:D3:D9:F7:62:D0 Trusted: yes
Changing 8C:D3:D9:F7:62:D0 trust succeeded
Connected and trusted: 8C:D3:D9:F7:62:D0%
```

```sh
$ go build && ./goble -d
Attempting to disconnect from 8C:D3:D9:F7:62:D0
[CHG] Device 8C:D3:D9:F7:62:D0 ServicesResolved: no
[SIGNAL] BREDR.Disconnected - org.bluez.Reason.Remote, Connection terminated by remote user
[SIGNAL] Disconnected - org.bluez.Reason.Remote, Connection terminated by remote user
[CHG] Device 8C:D3:D9:F7:62:D0 Connected: no
Disconnection successful
[CHG] Device 8C:D3:D9:F7:62:D0 Trusted: no
Changing 8C:D3:D9:F7:62:D0 untrust succeeded
Disconnected and untrusted: 8C:D3:D9:F7:62:D0%
```
