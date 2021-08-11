package main

import (
	"flag"
	"log"
)

//flag基本使用
// func main() {
// 	var name string
// 	flag.StringVar(&name, "name", "Go语言编程之旅", "帮助信息")
// 	flag.StringVar(&name, "n", "Go语言编程之旅", "帮助信息")
// 	flag.Parse()

// 	log.Printf("name: %s", name)
// }

var name string

//子命令的实现
func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		//该方法会返回带有指定名称和错误处理属性的空命令集给我们去使用，相当于创建一个新的命令集去支持子命令
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "Go 语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "n", "PHP 语言", "帮助信息")
		_ = phpCmd.Parse(args[1:])
	}

	log.Printf("name: %s", name)
}
