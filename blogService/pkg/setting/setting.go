package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//针对读取配置的行为进行封装，便于应用程序的使用
type Setting struct {
	vp *viper.Viper
}

//初始化本项目配置的基础属性
func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	//新增文件热更新的监听和变更处理
	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			//处理热更新的文件监听事件回调
			_ = s.ReloadAllSection()
		})
	}()
}
