package images

// 相关资料看 : https://mojotv.cn/2018/12/26/golang-generate
//go:generate file2byteslice -package=images -input=./tiles.png -output=./tiles.go -var=Tiles_png
//go:generate file2byteslice -package=images -input=./tanks.png -output=./tanks.go -var=Tanks_png
//go:generate file2byteslice -package=images -input=./items.png -output=./items.go -var=Items_png
//go:generate gofmt -s -w .
