package resources

// 相关资料看 : https://mojotv.cn/2018/12/26/golang-generate
//go:generate file2byteslice -package=resources -input=./cat.png -output=./cat.go -var=Cat_png
//go:generate file2byteslice -package=resources -input=./victory.png -output=./victory.go -var=Victory_png
//go:generate file2byteslice -package=resources -input=./failed.png -output=./failed.go -var=Failed_png
//go:generate file2byteslice -package=resources -input=./replay.png -output=./replay.go -var=Replay_png
//go:generate file2byteslice -package=resources -input=./HuaKangWaWaTi.ttc -output=./HuaKangWaWaTi.go -var=FontHuaKangWaWaTi_ttc
//go:generate gofmt -s -w .
