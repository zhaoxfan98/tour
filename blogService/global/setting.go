package global

import (
	"github.com/zhaoxfan98/blog/pkg/logger"
	"github.com/zhaoxfan98/blog/pkg/setting"
)

//读取了文件的配置信息后，还需要将配置信息和应用程序关联起来
//对最初预估的三个区段进行了配置并声明了全局变量，以便接下来关联起来，提供给应用程序内部调用
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	//完成日志库的编写后，我们需要定义一个 Logger 对象便于我们的应用程序使用
	Logger     *logger.Logger //用于日志文件的初始化
	JWTSetting *setting.JWTSettingS
)
