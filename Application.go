package go_config

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
	Server   Server   `yaml:"server"`
	Cloud    Cloud    `yaml:"cloud"`
	Database Database `yaml:"datasource"`
	Redis    Redis    `yaml:"redis"`
	Sentinel []Node   `yaml:"sentinel"`
	Rbac     Rbac     `yaml:"admin"`
}

/**
 * @Author yNsLuHan
 * @Description: 管理员权限、角色id配置
 */
type Rbac struct {
	AdminRoleId     uint64 `yaml:"admin-role-id"`
	AdminPermission string `yaml:"admin-permission"`
}

/**
 * @Author yNsLuHan
 * @Description: redis Pool 构造体
 */
type Redis struct {
	Host        string `yaml:"host"`
	Port        uint64 `yaml:"port"`
	Password    string `yaml:"password"`
	Db          uint64 `yaml:"db"`
	AdminPrefix string `yaml:"admin-prefix"`
	prefix      string `yaml:"prefix"`
	MaxIdle     uint64 `yaml:"max-idle"`
	MaxActive   uint64 `yaml:"max-active"`
	Timeout     uint64 `yaml:"timeout"`
	Enable      bool   `yaml:"enable"`
}

/**
 * @Author yNsLuHan
 * @Description: redis Sentinel 构造体  redis Sentinel Node 构造体
 */
type Node struct {
	// 从gin上下文获取的名字
	ContentName     string `yaml:"content-name"`
	Name            string `yaml:"name"`
	AdminPrefix     string `yaml:"admin-prefix"`
	prefix          string `yaml:"prefix"`
	SentinelAddress string `yaml:"sentinel-address"`
	Password        string `yaml:"password"`
	Db              uint64 `yaml:"db"`
	PoolSize        uint64 `yaml:"pool-size"`
	MinIdleConns    uint64 `yaml:"min-idle-conns"`
	Timeout         uint64 `yaml:"timeout"`
	Enable          bool   `yaml:"enable"`
}

/**
 * @Author yNsLuHan
 * @Description: 服务器属性构造体
 */
type Server struct {
	Host               string `yaml:"host"`
	Name               string `yaml:"name"`
	Port               uint64 `yaml:"port"`
	Debug              bool   `yaml:"debug"`
	Token              Token  `yaml:"token"`
	EnableLogger       bool   `yaml:"enable-logger"`
	SaveLogger         bool   `yaml:"save-logger"`
	LoggerDir          string `yaml:"logger-dir"`
	EnableJwt          bool   `yaml:"enable-jwt"`
	EnableRedisJwt     bool   `yaml:"enable-redis-jwt"`
	EnableRbac         bool   `yaml:"enable-rbac"`
	EnableRecover      bool   `yaml:"enable-recover"`
	EnableErrorRecover bool   `yaml:"enable-error-recover"`
}

/**
 * @Author yNsLuHan
 * @Description: 服务器属性构造体
 */
type Token struct {
	SecretKey string `yaml:"secret-key"`
	expired   uint64 `yaml:"expired"`
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
	Host      string `yaml:"host"`
	Port      uint64 `yaml:"port"`
	Namespace string `yaml:"namespace"`
	LogDir    string `yaml:"log-dir"`
	LogLevel  string `yaml:"log-level"`
	Group     string `yaml:"group"`
	Cluster   string `yaml:"cluster"`
}

/**
 * @Author yNsLuHan
 * @Description: Consul
 */
type Consul struct {
	Host      string `yaml:"host"`
	Port      uint64 `yaml:"port"`
	Namespace string `yaml:"namespace"`
	Group     string `yaml:"group"`
	Cluster   string `yaml:"cluster"`
}

/**
 * @Author yNsLuHan
 * @Description: 数据库构造体
 */
type Database struct {
	Host         string `yaml:"host"`
	Port         uint64 `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Db           uint64 `yaml:"db"`
	MaxIdle      uint64 `yaml:"max-idle"`
	MaxOpenConns uint64 `yaml:"max-open-conns"`
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
		log.Fatalln("error 配置文件打开异常： ", err.Error())
	}
	//
	bytes, err := ioutil.ReadAll(file)
	//
	if err != nil {
		log.Fatalln("error 配置文件读取异常： ", err.Error())
	}
	//
	err = yaml.Unmarshal(bytes, &c)
	//
	if err != nil {
		log.Fatalln("error 配置文件绑定结构体异常： ", err.Error())
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
