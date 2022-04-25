# Cabinet


### What is cabinet
Cabinet is a tool set for coder powered by [fyne](https://github.com/fyne-io/fyne).

### Features of cabinet
* conversion
    - base convert (bin/oct/dec/hex)
    - time convert (timestamp/utctime/localtime)
    - case convert (upper/lower/capital/capital(word)/camel/pascal/snake/constant)

* encode/decode
    - base64 (base16/base64/base58)
    - url
    - html

* image
    - image format convert
    - compress
    - image to json (to/from json)
    - qrcode (str to/from qrcode)

* json
    - json format (compress/format)
    - json to yaml (to/from yaml)

* digest
    - calc hash (md5/sha1/sha224/sha256/sha384/sha512/keccak256)
    - file checksum (md5/sha1/sha256/sha384/sha512)
    
* crypto (encrypt/decrypt)
    - aes (cbc/cfb)
    - des

* signature (sign/verify)
    - rsa (256/1024/2048)
    - ecdsa (secp256k1)

* network
    - get my ip
    - ip mask (ip mask calculator)
    - web server (static file web server)

* other
    - uuid (v1/v2/v3/v4)
    - random password
    - rsa key pair
    - ecdsa key pair
    - arithmetic expression calculator
    - markdown


### How To generate bundle.go
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