package ikun

import (
	"errors"
	"io/ioutil"
	"strings"
)

// Config 变量用于保存配置信息
var Config *ServerConfig

// ServerConfig 结构体用于表示配置
type ServerConfig struct {
	Server struct {
		Port string
	}
}

// ReadConfig 方法用于读取配置文件并将配置应用到服务器中
func (s *srv) ReadConfig(path string) (*ServerConfig, error) {
	// 读取配置文件
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// 解析配置文件到 Config 结构体
	config, err := parseYAML(string(data))
	if err != nil {
		return nil, err
	}

	// 将读取到的配置信息保存在 Config 变量中
	Config = config

	// 返回配置结构体
	return config, nil
}

// parseYAML 解析简单的 YAML 字符串
func parseYAML(data string) (*ServerConfig, error) {
	config := &ServerConfig{}
	lines := strings.Split(data, "\n")
	var currentSection string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			// 跳过空行和注释行
			continue
		}
		if strings.HasSuffix(line, ":") {
			// 解析部分标题 (section)
			currentSection = strings.TrimSuffix(line, ":")
		} else {
			// 解析键值对
			parts := strings.SplitN(line, ":", 2)
			if len(parts) < 2 {
				return nil, errors.New("invalid YAML format")
			}
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if currentSection == "server" && key == "port" {
				config.Server.Port = value
			}
		}
	}
	return config, nil
}
