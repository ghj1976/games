
**常用命令**

```shell
cd /Users/guohongjun/Documents/MyCodes/mygocodes/src/github.com/ghj1976/games/tank90/tank
go build

GOOS=js GOARCH=wasm go build -o tank90.wasm 


cd /Users/guohongjun/Documents/MyCodes/mygocodes/src/github.com/ghj1976/games/tank90/images
go generate




GO111MODULE=off go get github.com/blizzy78/ebitenui
GO111MODULE=off go get github.com/gabstv/ebiten-imgui

# go get 下载的包不在src目录下生成,而全部到了$GOPATH$/pkg/mod 目录下



```
