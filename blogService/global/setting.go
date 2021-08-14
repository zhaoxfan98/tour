package global

//读取了文件的配置信息后，还需要将配置信息和应用程序关联起来
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)
