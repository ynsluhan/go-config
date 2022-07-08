package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
)
// 配置文件
var conf Config

/**
 * @Author yNsLuHan
 * @Description: 初始化配置文件
 */
func init() {
	ReadConf(&conf)
}

/**
 * @Author yNsLuHan
 * @Description: 获取配置文件实例对象
 * @return Config
 */
func GetConf() *Config {
	return &conf
}

/**
 * @Author yNsLuHan
 * @Description: config配文件构造体
 */
type Config struct {
	Server       Server   `yaml:"server"`
	Cloud        Cloud    `yaml:"cloud"`
	Database     Database `yaml:"datasource"`
	Redis        Redis    `yaml:"redis"`
	Sentinel     []Node   `yaml:"sentinel"`
	Rbac         Rbac     `yaml:"admin"`
}

/**
 * @Author yNsLuHan
 * @Description: 管理员权限、角色id配置
 */
type Rbac struct {
	AdminRoleId     int    `yaml:"admin-role-id"`
	AdminPermission string `yaml:"admin-permission"`
}

/**
 * @Author yNsLuHan
 * @Description: redis Sentinel 构造体  redis Sentinel Node 构造体
 */
type Node struct {
	// 从gin上下文获取的名字
	ContentName     string `yaml:"content-name"`
	Name            string `yaml:"name"`
	SentinelAddress string `yaml:"sentinel-address"`
	Password        string `yaml:"password"`
	Db              int    `yaml:"db"`
	PoolSize        int    `yaml:"pool-size"`
	MinIdleConns    int    `yaml:"min-idle-conns"`
}

/**
 * @Author yNsLuHan
 * @Description: 服务器属性构造体
 */
type Server struct {
	Host  string `yaml:"host"`
	Name  string `yaml:"name"`
	Port  uint64 `yaml:"port"`
	Debug bool   `yaml:"debug"`
	Token Token  `yaml:"token"`
}

/**
 * @Author yNsLuHan
 * @Description: 服务器属性构造体
 */
type Token struct {
	SecretKey string `yaml:"secret-key"`
	expired   string `yaml:"expired"`
}

/**
 * @Author yNsLuHan
 * @Description: Cloud
 */
type Cloud struct {
	Nacos  Nacos  `yaml:"nacos"`
	Consul Consul `yaml:"consul"`
}

/**
 * @Author yNsLuHan
 * @Description: Nacos
 */
type Nacos struct {
	Host string `yaml:"host"`
	Port uint64 `yaml:"port"`
}

/**
 * @Author yNsLuHan
 * @Description: Consul
 */
type Consul struct {
	Host string `yaml:"host"`
	Port uint64 `yaml:"port"`
}

/**
 * @Author yNsLuHan
 * @Description: 数据库构造体
 */
type Database struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Db           string `yaml:"db"`
	MaxIdle      string `yaml:"max-idle"`
	MaxOpenConns string `yaml:"max-open-conns"`
}

/**
 * @Author yNsLuHan
 * @Description: redis Pool 构造体
 */
type Redis struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	Db        string `yaml:"db"`
	MaxIdle   string `yaml:"max-idle"`
	MaxActive string `yaml:"max-active"`
}

/**
 * @Author yNsLuHan
 * @Time 2021-06-08 15:05:18
 * @Description: 获取yaml配置文件
 * @param c config
 */
func ReadConf(c *Config) {
	// 获取项目路径
	GetWd, _ := os.Getwd()
	// 拼接config路径
	var FilePath = path.Join(GetWd, "src/com/yniii/config")
	// 获取配置文件路径
	var FileName = path.Join(FilePath, "application.yaml")
	// 读取配置
	file, err := os.Open(FileName)
	//
	if err != nil {
		log.Println("error 配置文件打开异常： ", err.Error())
	}
	//
	bytes, err := ioutil.ReadAll(file)
	//
	if err != nil {
		log.Println("error 配置文件读取异常： ", err.Error())
	}
	//
	err = yaml.Unmarshal(bytes, &c)
	//
	if err != nil {
		log.Println("error 配置文件绑定结构体异常： ", err.Error())
	}
}

/**
 * @Author yNsLuHan
 * @Time 2021-06-08 15:05:38
 * @Description: 获取Env
 * @param host host
 * @return string env
 */
func GetEnv(host string) string {
	//
	var env = os.Getenv("Gin")
	//
	if env == "Products" {
		return "127.0.0.1"
	} else {
		return host
	}
}
