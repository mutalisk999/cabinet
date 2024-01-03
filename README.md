# Cabinet

### What is cabinet

Cabinet is a tool set for coder powered by [fyne](https://github.com/fyne-io/fyne).

light theme:
[](https://github.com/mutalisk999/cabinet/blob/main/assets/light.png)

dark theme:
[](https://github.com/mutalisk999/cabinet/blob/main/assets/dark.png)


---

### Features Supported

- [x] conversion
    - [x] base convert (bin/oct/dec/hex)
    - [x] time convert (timestamp/utctime/localtime)
    - [x] case convert (upper/lower/capital/capital(word)/camel/pascal/snake/constant)

- [x] encoder/decoder
    - [x] base64 (base16/base32/base64/base58/base85) encode/decode
    - [x] url escape/unescape
    - [x] html escape/unescape

- [ ] image
    - [ ] image format convert
    - [ ] compress
    - [ ] image to base64 (to/from base64)
    - [ ] qrcode (str to/from qrcode)

- [ ] json
    - [ ] json format (compress/format)
    - [ ] json to yaml (to/from yaml)

- [ ] digest
    - [ ] calc hash (md5/sha1/sha224/sha256/sha384/sha512/keccak256)
    - [ ] file checksum (md5/sha1/sha256/sha384/sha512)

- [ ] crypto (encrypt/decrypt)
    - [ ] aes (cbc/cfb)
    - [ ] des

- [ ] signature (sign/verify)
    - [ ] rsa (256/1024/2048)
    - [ ] ecdsa (secp256k1)

- [x] network
    - [x] get my ip (external ip/local interface ip)
    - [x] ip mask (ip mask calculator)
    - [x] web server (static file web server)

- [ ] others
    - [x] uuid (uuid1/uuid4)
    - [x] random password
    - [ ] rsa key pair
    - [ ] ecdsa key pair
    - [ ] arithmetic expression calculator
    - [ ] markdown

---

### How To Generate bundle.go

```
go get fyne.io/fyne/cmd/fyne
fyne bundle FangZhengHeiTiJianTi.ttf > bundle.go
```

### How To Build

```
git clone https://github.com/mutalisk999/cabinet
cd cabinet
make
```
