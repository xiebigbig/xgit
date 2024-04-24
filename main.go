package main

import (
	"fmt"
	"os"
	"time"
	"xgit/utils"
	"xgit/core/version"
)
// 入口文件
func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法", "xgit add|list|goto")
		os.Exit(1)
	}

	if os.Args[1] != "goto" && os.Args[1] != "add" && os.Args[1] != "list" {
		fmt.Println("用法", "xgit add|list|goto")
		os.Exit(1)
	}

	if os.Args[1] == "goto" && len(os.Args) < 3 {
		fmt.Println("必须输入要goto到的版本")
		os.Exit(1)
	}

	action := os.Args[1]
	versionNum := ""
	if os.Args[1] == "goto" {
		versionNum = os.Args[2]
	}

	start := time.Now()
	if action == "add" {
		version.Commit(utils.GetWd(), utils.Tmd5())
		elapsed := time.Since(start)
		fmt.Printf("执行时长 %s", elapsed)
		os.Exit(0)
	}

	if action == "list" {
		if len(os.Args) < 3 {
			version.ShowVersions()
			elapsed := time.Since(start)
			fmt.Printf("执行时长 %s", elapsed)
			os.Exit(0)
		} else {
			version.ShowFileVersion(os.Args[2])
			elapsed := time.Since(start)
			fmt.Printf("执行时长 %s", elapsed)
			os.Exit(0)
		}
	}

	if action == "goto" {
		version.Checkout(versionNum)
		elapsed := time.Since(start)
		fmt.Printf("执行时长 %s", elapsed)
		os.Exit(0)
	}
}
