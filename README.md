# QR code generator

Simple app for QR code generation. If you need testing, saved or packet generation QR codes. (Tested on windows 10/11 and Qt 5.12.11)

### Notes: 

For build app with GUI use Qt (install first and check ENV qmake and mingw32-make). For correct build Not use '-j16'

```sh
$ mingw32-make clean -j16 & qmake & mingw32-make
```

For building DLL ('go_src' folder):

```sh
go build -o qrcode.dll -buildmode=c-shared -ldflags "-s -w"
```

For using DLL on windows use

- Save file:

```sh
RUNDLL32 qrcode.dll, Rundll32save "test example" 1 testName.png
```

- Copy to clipboard:

```sh
RUNDLL32 qrcode.dll, Rundll32ToClipboard "test example" 3
```

Screenshots:


<div style="width:50%">
<img src="/pictures/001.png" >
</div>