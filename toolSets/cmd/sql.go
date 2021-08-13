package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zhaoxfan98/Tour/internal/sql2struct"
)

//声明7个全局变量，用于接收外部的命令行参数
var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

//定义对应的子命令
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql 转化和处理",
	Long:  "sql 转化和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

//完成对数据库的查询 模板对象的组装 渲染等动作
var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql 转化",
	Long:  "sql 转化",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err:%v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	//初始化和命令行参数的绑定
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "root", "请输入数据库账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "301421", "请输入数据库密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库编码")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名")
}
