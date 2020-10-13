package main

import "fmt"

type SysMenu struct {
	SysBaseMenu
	Children   []SysMenu              `json:"children"`
	Parameters []SysBaseMenuParameter `json:"parameters"`
}

type SysBaseMenu struct {
	Children   []SysBaseMenu          `json:"children"`
	Parameters []SysBaseMenuParameter `json:"parameters"`
}

type SysBaseMenuParameter struct {
	SysBaseMenuID uint
	Type          string `json:"type"`
	Key           string `json:"key"`
	Value         string `json:"value"`
}

func main() {
	sbmp := []SysBaseMenuParameter{{
		1234,
		"aaa",
		"bbb",
		"ccc",
	}, {
		5678,
		"sss",
		"sss",
		"sss",
	},
	}
	msbm := SysBaseMenu{
		[]SysBaseMenu{
			{
				[]SysBaseMenu{},
				sbmp,
			},
		},
		sbmp,
	}

	sm := SysMenu{
		msbm,
		[]SysMenu{},
		sbmp,
	}

	// 可以共同存在,链式而不是直接嵌入
	fmt.Printf("%v\n%v\n%v\n", sm.SysBaseMenu, sm.Children, sm.Parameters)
}
