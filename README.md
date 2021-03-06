# 魔方旋转demo

golang运行: go run cube.go UuDdFfBbLlRr

erlang运行: escript cube_erl UuDdFfBbLlRr


# 数据结构

以cube(小正方块)为基本旋转单位来实现, 旋转过程抽象成旋转层(layer)上8个cube间隔替换, 并沿同一方向自转.

```
3层Cube分布:
06 05 04
07 -- 03 
00 01 02
----------
11 -- 10
-- -- --
08 -- 09
----------
18 17 16
19 -- 15
12 13 14

层分布:
   G3
O2 Y0 R4
   B1
   W5

face分布:
         6 5 4
         7 G 3
         0 1 2

0 7 6    6 5 4    4 3 2
1 O 5    7 Y 3    5 R 1
2 3 4    0 1 2    6 7 0

         6 5 4
         7 B 3
         0 1 2

         6 5 4
         7 W 3
         0 1 2

旋转layer:
顺: 01234567 <- 23456701:  x = (x+2)%8
逆: 01234567 <- 67012345:  x = (x+6)%8
cube自转:
顺: 0123 <- 3012:  x = (x+3)%4
逆: 0123 <- 1230:  x = (x+1)%4
```


# 其他

另一种解决方案[RubiksCube](https://github.com/LaoJQ/RubiksCube), 以face(小正方块上的一个面)为旋转基本单位, 旋转过程抽象成face颜色的替换.
