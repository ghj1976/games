


```bash

cd /Users/guohongjun/Documents/MyCodes/mygocodes/src/github.com/ghj1976/games/nerveincats/nic



## 编译web版  https://ebiten.org/documents/webassembly.html

GOOS=js GOARCH=wasm go build -o nic.wasm 

## 编译Android版本
gomobile build -target=android github.com/ghj1976/games/nerveincats/nic

```