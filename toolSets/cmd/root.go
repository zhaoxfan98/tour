package cmd

import "github.com/spf13/cobra"

// 作为根命令

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	//将响应的子命令进行注册
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)

	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", ` 需要计算的时间，有效单位为时间戳或已格式化后的时间 `)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", ` 持续时间，有效时间单位为"ns", "us" (or "µ s"), "ms", "s", "m", "h"`)

	rootCmd.AddCommand(sqlCmd)
}
