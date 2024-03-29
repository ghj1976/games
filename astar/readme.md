A*（A-Star)算法是一种静态路网中求解最短路径最有效的直接搜索方法。
[百度百科A*算法](https://baike.baidu.com/item/A%2A%E7%AE%97%E6%B3%95)

## A*算法关键点：

### 维护两个列表

* 一个记录下所有被考虑来寻找最短路径的方块（open 列表）
    * 需要在其中找到最小F值
* 一个记录下不会再被考虑的方块（closed列表）
    * 需要找在不在close列表

### 计算每个节点的综合评估距离

评估公式： f(n) = g(n) + h(n) 

#### f(n) 节点n的综合评估距离 
走到终点消耗的代价值 ，要知道F值，需要计算G和H的值 。
估计是 evaluation **function** 的简写。

#### h(n)  启发式估算距离
任意顶点n到目标顶点的估算距离
估计是 **heuristic** function 的缩写
这个常被称为探视，因为我们不确定移动量是多少 – 仅仅是一个估算值。
* 移动量估算值离真实值越接近，最终的路径会更加精确。如果估算值停止作用，很可能生成出来的路径不会是最短的（但是它可能是接近的）。

有关这个的详细分析可以看 [A*算法不同启发算法对寻路影响](./heuristic/readme.md)
后面有专门对这个的分析。

#### g(n)  实际消耗
从起点到任意顶点n的实际距离 
估计是 **generated** 的缩写
G值是从起点到该点，是一段已经走过的路程，代价是准确可知的

* 如果你允许对角线移动，你可以针对对角线移动把移动量调得大一点。
* 如果你有不同的地形，你可以将相应的移动量调整得大一点 – 例如针对一块沼泽，水，或者猫女海报:-)



### 伪代码
来自百度百科

````
算起点的h(s);
将起点放入OPEN表;
while(OPEN!=NULL)
{
    从OPEN表中取f(n)最小的节点n;
    if(n节点==目标节点)
        break;
    for(当前节点n的每个临近节点X)
    {
        计算f(X);
        if(X in OPEN)
            if(新的f(X)< OPEN中的f(X))
            {
                把n设置为X的父亲;
                更新OPEN表中的f(n);
            }
            
        if(X in CLOSE)
            continue;
            
        if(X not in both)
        {
            把n设置为X的父亲;
            求f(X);
            并将X插入OPEN表中;//还没有排序
        }
    }//endfor
    将n节点插入CLOSE表中;
    按照f(n)将OPEN表中的节点排序; //实际上是比较OPEN表内节点f的大小，从最小路径的节点向下进行。
}//endwhile(OPEN!=NULL)
````
保存路径，即从终点开始，每个节点沿着父节点移动直至起点，这就是你的路径；

#### 路径回溯
通过回溯各个前后节点的关系得到最佳路径
每个节点都记录前一个节点，按照前节点倒退的路径就是最佳路径

#### 加入open列表的时候，要保证的3点

* 这个点不能在close表中
* 这个点不能在open表中
* 这个点是合法的，什么叫合法？这里简单的就是可以通行的，不能是障碍物

### 演示代码

* [JavaScript](./html/js.md)
* [Golang](./go.md)

### 参考：
* [A*寻路算法](https://www.jianshu.com/p/65282bd32391)
* [不同寻路算法的可视化比较](http://qiao.github.io/PathFinding.js/visual/)


## 八卦

### 为啥叫 A Star ？ 
在AStar之前就有A1 A2算法，Peter Hart改进A2算法后统称这类算法A*
估计是觉得这是A算法的升级最终版所以就叫A*
<https://www.zhihu.com/question/29528928>


##  go编译

`
cd /Users/guohongjun/Documents/MyCodes/mygocodes/src/github.com/ghj1976/games/astar
GOOS=js GOARCH=wasm go build -o astar.wasm 

`

可选内容
地图类型： 坦克大战、随机地图、U型地图
A*寻路算法H值算法： x + y 、 2 * (x + y)、 x*x + y*y