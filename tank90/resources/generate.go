package resources

// 相关资料看 : https://mojotv.cn/2018/12/26/golang-generate
//go:generate file2byteslice -package=resources -input=./tile.png -output=./tile.go -var=Tile_png
//go:generate file2byteslice -package=resources -input=./bird.png -output=./bird.go -var=Bird_png
//go:generate file2byteslice -package=resources -input=./enemy_1.png -output=./enemy_1.go -var=Enemy_1_png
//go:generate file2byteslice -package=resources -input=./enemy_2.png -output=./enemy_2.go -var=Enemy_2_png
//go:generate file2byteslice -package=resources -input=./enemy_3.png -output=./enemy_3.go -var=Enemy_3_png
//go:generate file2byteslice -package=resources -input=./enemy_4.png -output=./enemy_4.go -var=Enemy_4_png
//go:generate file2byteslice -package=resources -input=./enemy_5.png -output=./enemy_5.go -var=Enemy_5_png
//go:generate file2byteslice -package=resources -input=./tank_user_1.png -output=./tank_user_1.go -var=User_1_png
//go:generate file2byteslice -package=resources -input=./tank_user_2.png -output=./tank_user_2.go -var=User_2_png
//go:generate file2byteslice -package=resources -input=./move.ogg -output=./move_ogg.go -var=Move_ogg
//go:generate file2byteslice -package=resources -input=./setting.ttf -output=./setting_ttf.go -var=Setting_ttf
//go:generate gofmt -s -w .
