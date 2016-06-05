# memcat

Copy raw kernel memory to stdout.

## example usage

Peak in your BIOS:

`sudo ./memcat -base 0x000F0000 -len 0x00010000 | strings -n 20`

```
f=MSDOt-f=MSWIt%f=NTFSt
OIEMAG  IZ PTAPA IlFpoyp
@0080 a  IHATHC IDC-R7703                        iPnoee rDCR-MOA ATIP
Reboot and Select proper Boot device
or Insert Boot Media in selected Boot device and press a key
PCI ROM Setup, B00 D00 F0
?l=\?*@Q?Q?Q?Q?Q?Q?~?X?
P1: INTEL SSDSC2BW240A4
Booting from CDROM with Multiple Boot Image
Select Boot Image  :
No Emulation Image :
1.2MB  Floppy Image:
1.44MB Floppy Image:
2.88MB Floppy Image:
Hard Disk Image    :
Unknown Image      :
```

## cross-compiling

If you want to run this on ARM:

`GOOS=linux GOARCH=arm go build`

If you didn't `go get` this package, you will need to:

`go get golang.org/x/sys/unix`
