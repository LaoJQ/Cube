package main

import (
    "fmt"
    "os"
)

type RubikCube struct {
    cubeList []Cube
}

type Cube struct {
    colorList []byte
}

func newRubikCube() *RubikCube {
    rc := new(RubikCube)
    for i:=0; i<20; i++ {
        c := Cube{
            colorList : []byte{'Y', 'B', 'O', 'G', 'R', 'W'},
        }
        rc.cubeList = append(rc.cubeList, c)
    }
    return rc
}

func (rc *RubikCube) print() {
    for i:=0; i<6; i++ {
        lci := LayerCubeIndex[i]
        colorArray := make([]byte, 9)
        for _, cubeId := range lci {
            cube := rc.cubeList[cubeId]
            colorArray = append(colorArray, cube.colorList[i])
        }
        fmt.Println(string(colorArray))
    }
}

func (rc *RubikCube) rotate(layer int, clockWise bool) {
    layerMove := 2
    cubeMove := 3
    if !clockWise {
        layerMove = 6
        cubeMove = 1
    }
    lci := LayerCubeIndex[layer]
    rc.cubeList[lci[0]],
    rc.cubeList[lci[1]],
    rc.cubeList[lci[2]],
    rc.cubeList[lci[3]],
    rc.cubeList[lci[4]],
    rc.cubeList[lci[5]],
    rc.cubeList[lci[6]],
    rc.cubeList[lci[7]] =
        rc.cubeList[lci[(0+layerMove)%8]],
    rc.cubeList[lci[(1+layerMove)%8]],
    rc.cubeList[lci[(2+layerMove)%8]],
    rc.cubeList[lci[(3+layerMove)%8]],
    rc.cubeList[lci[(4+layerMove)%8]],
    rc.cubeList[lci[(5+layerMove)%8]],
    rc.cubeList[lci[(6+layerMove)%8]],
    rc.cubeList[lci[(7+layerMove)%8]]

    crs := CubeRotateSeq[layer]
    for i:=0; i<8; i++ {
        c := rc.cubeList[lci[i]].colorList
        c[crs[0]],
        c[crs[1]],
        c[crs[2]],
        c[crs[3]] =
            c[crs[(0+cubeMove)%4]],
        c[crs[(1+cubeMove)%4]],
        c[crs[(2+cubeMove)%4]],
        c[crs[(3+cubeMove)%4]]
    }
}

var LayerCubeIndex [][]int = [][]int{
    []int{0, 1, 2, 3, 4, 5, 6, 7},
    []int{12, 13, 14, 9, 2, 1, 0, 8},
    []int{18, 19, 12, 8, 0, 7, 6, 11},
    []int{6, 5, 4, 10, 16, 17, 18, 11},
    []int{14, 15, 16, 10, 4, 3, 2, 9},
    []int{18, 17, 16, 15, 14, 13, 12, 19},
}

var CubeRotateSeq [][]int = [][]int{
    []int{1, 2, 3, 4},
    []int{2, 0, 4, 5},
    []int{0, 1, 5, 3},
    []int{0, 2, 5, 4},
    []int{1, 0, 3, 5},
    []int{2, 1, 4, 3},
}


type Actions struct {
    layer int
    clockWise bool
}

var ActionsMap map[byte]Actions = map[byte]Actions{
    'U' : Actions{0, true},
    'u' : Actions{0, false},

    'F' : Actions{1, true},
    'f' : Actions{1, false},

    'L' : Actions{2, true},
    'l' : Actions{2, false},

    'B' : Actions{3, true},
    'b' : Actions{3, false},

    'R' : Actions{4, true},
    'r' : Actions{4, false},

    'D' : Actions{5, true},
    'd' : Actions{5, false},
}

func main() {
    rc := newRubikCube()
    if len(os.Args) >= 2 {
        for _, op := range []byte(os.Args[1]) {
            if act, ok := ActionsMap[op]; ok {
                rc.rotate(act.layer, act.clockWise)
            }
        }
    }
    rc.print()
}
