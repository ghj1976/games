**说明：**

* psd 文件有两个红框是为了避免 psd 文件生成 png时，空白区域全部都变没了，导致位置不准了。
* 每个外形的图片只有一个，控制每个外形展示多长是在代码里面控制的，这样可以减少尺寸。

把图片文件编译到代码中的技术看： https://mojotv.cn/2018/12/26/golang-generate

日常生成命令
```bash
#!/bin/bash
cd /Users/guohongjun/Documents/project/mygocodes/src/github.com/ghj1976/games/nerveincats/resources
go generate

```
