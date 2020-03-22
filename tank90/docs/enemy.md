**敌人坦克说明**

|  类型   | level 1 <br/> 灰色  | level 2 <br/> 绿色 | level 3 <br/> 橙色 | level 4 <br/> 红色 | 分数 | 炮弹 | 速度 | 是否能击毁石墙 |
|  :---- | ----  | ----  | ----  | ----  | ----  | ----  | ----  | ----  |
|  类型 1   | ![](enemy11.png)  | ![](enemy12.png) | ![](enemy13.png) | ![](enemy14.png) | 100 | 1 | 2.6/8.0 | 否 |
|  类型 2   | ![](enemy21.png)  | ![](enemy22.png) | ![](enemy23.png) | ![](enemy24.png) | 200 | 1 | 4.6/10.0 | 否 |
|  类型 3   | ![](enemy31.png)  | ![](enemy32.png) | ![](enemy33.png) | ![](enemy34.png) | 300 | 2 | 3.2/11.2 | 否 |
|  类型 4   | ![](enemy41.png)  | ![](enemy42.png) | ![](enemy43.png) | ![](enemy44.png) | 400 | 1 | 3.6/11.6 | 否 |
|  类型 5   | ![](enemy51.png)  | ![](enemy52.png) | ![](enemy53.png) | ![](enemy54.png) | 500 | 1 | 3.0/9.6 | 是 |

击毁 level 4 敌人，将会出现奖励


**难度级别说明**
* 普通级别 所有敌人的速度 + 0.1 ， 
* 难度级别 所有敌人的速度 +0.2 ，同时更聪明。




**参考坦克逻辑一**

* 随机一定时间发射一枚炮弹。   比如  
  r.nextInt(40)>36  // 随机数取40内的数字，如果数字大与36就发射。 
  参考<https://blog.csdn.net/a1275302036/article/details/54232751>
* 敌方的坦克根据随机数来控制随机的方向和路径的，
  <https://www.write-bug.com/article/2548.html>
* 当敌方坦克撞到阻碍物时，会转回到前一步的位置，从而解决了坦克撞到阻碍物不回头的问题。 
  



**复杂算法**

* [A*搜索算法原理与控制台坦克大战AI实现](https://blog.csdn.net/JILVAN21/article/details/82863983)
* [搜索算法之坦克大战（bfs+优先队列）](https://blog.csdn.net/lee371042/article/details/79114718)
* 通过使用一个存储了地图内各个元素的二维数组Game.map，使用广度优先算法遍历出一条路线，将结果存放于栈之中。 <https://blog.csdn.net/madonghyu/article/details/78960228>
* [广度优先算法实现的一款“最短路”游戏](https://github.com/MummyDing/Tank-Battle-Data-Structure-Design)
* [最短路径——Dijkstra算法以及二叉堆优化（含证明）](https://www.bbsmax.com/A/Vx5Mre3gdN/)
* [单源最短路——Dijkstara算法](http://bbs.eeworld.com.cn/thread-1107262-1-1.html) 


【Unity3D】制作自己进化的AI-神经网络控制坦克大战
<https://www.bilibili.com/video/av14819865/>
<https://github.com/trulyspinach/Unity-Neural-Network-Tanks-AI>


[双人回合制坦克大战游戏](https://wiki.botzone.org.cn/index.php?title=Tank)


**扩展思路**

* [百度之星坦克大战寻矿并在视野内有敌坦克时攻击的AI代码](http://www.aiseminar.com/bbs/forum.php?mod=viewthread&tid=1205)
* [AI坦克对战（实现人机）](https://blog.csdn.net/lossatsea/article/details/80737476)
* [游戏中的AI算法总结与改进](https://blog.csdn.net/cordova/article/details/51607407) 