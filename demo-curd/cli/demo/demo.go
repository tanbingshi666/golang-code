package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	confFile string
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A simple command-line application",
	Long:  "This is a simple command-line application built with Cobra.",
	Run: func(cmd *cobra.Command, args []string) {
		// 这里是主命令的执行逻辑
		fmt.Println("Hello from the main command!")
	},
}

var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "A subcommand",
	Long:  "This is a subcommand of the main application.",
	Run: func(cmd *cobra.Command, args []string) {
		// 这里是子命令的执行逻辑
		fmt.Println(args)
		fmt.Println("Hello from the subcommand!")
	},
}

func init() {
	// 将子命令添加到主命令中
	subCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	rootCmd.AddCommand(subCmd)
}

func main() {
	// 执行命令行应用
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// go run .\demo.go sub -f hello.conf
	// go run .\demo.go -f hello.conf sub
	fmt.Println(confFile)
}
