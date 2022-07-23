package main

import (
	"github.com/chenhu1001/marketool/cron"
	"github.com/chenhu1001/marketool/goutils"
	"github.com/chenhu1001/marketool/logging"
	"github.com/chenhu1001/marketool/routes"
	"github.com/chenhu1001/marketool/routes/response"
	"github.com/chenhu1001/marketool/webserver"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"path"
	"strings"
)

func main() {
	// 初始化配置文件
	InitWithConfigFile("./config.toml")

	// 创建 gin app
	middlewares := DefaultGinMiddlewares()
	server := webserver.NewGinEngine(middlewares...)

	//// 启动定时任务
	//if viper.GetString("env") == "prod" {
	//	// 需要执行的命令： free -mh
	//	cmd := exec.Command("/bin/bash", "-c", `rm -rf /root/soft/investool/eastmoney_funds_list.json /root/soft/investool/fund_4433_list.json /root/soft/investool/fund_all_list.json /root/soft/investool/fund_managers.json /root/soft/investool/fund_type_list.json /root/soft/investool/industry_list.json && cp /root/temp/eastmoney_funds_list.json /root/temp/fund_4433_list.json /root/temp/fund_all_list.json /root/temp/fund_managers.json /root/temp/fund_type_list.json /root/temp/industry_list.json /root/soft/investool/`)
	//
	//	// 获取管道输入
	//	_, err := cmd.StdoutPipe()
	//	if err != nil {
	//		fmt.Println("无法获取命令的标准输出管道", err.Error())
	//		return
	//	}
	//
	//	// 执行Linux命令
	//	if err := cmd.Start(); err != nil {
	//		fmt.Println("Linux命令执行失败，请检查命令输入是否有误", err.Error())
	//		return
	//	}
	//
	//	_, err1 := http.Get("https://api.day.app/Y3uKSZF6URZQTU7FXuTUUM/爬取成功")
	//	if err1 != nil {
	//		log.Println("err")
	//	}
	//}

	// 启动定时任务
	if viper.GetString("env") == "prod" {
		cron.RunCronJobs(true)
	}

	// 注册路由
	routes.Register(server)

	// 运行服务
	webserver.Run(server)
}

func init() {
	viper.SetDefault("app.chan_size", 50)
}

// InitWithConfigFile 根据 webserver 配置文件初始化 webserver
func InitWithConfigFile(configFile string) {
	// 加载配置文件内容到 viper 中以便使用
	configPath, file := path.Split(configFile)
	if configPath == "" {
		configPath = "."
	}
	ext := path.Ext(file)
	configType := strings.Trim(ext, ".")
	configName := strings.TrimSuffix(file, ext)
	logging.Infof(nil, "load %s type config file %s from %s", configType, configName, configPath)

	if err := goutils.InitViper(configPath, configName, configType, func(e fsnotify.Event) {
		logging.Warn(nil, "Config file changed:"+e.Name)
		logging.SetLevel(viper.GetString("logging.level"))
	}); err != nil {
		// 文件不存在时 1 使用默认配置，其他 err 直接 panic
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(err)
		}
		logging.Error(nil, "Init viper error:"+err.Error())
	}

	// 设置 viper 中 webserver 配置项默认值
	viper.SetDefault("env", "localhost")

	viper.SetDefault("server.addr", ":4869")
	viper.SetDefault("server.mode", gin.ReleaseMode)
	viper.SetDefault("server.pprof", true)

	viper.SetDefault("apidocs.title", "marketool swagger apidocs")
	viper.SetDefault("apidocs.desc", "Using marketool to develop gin app on fly.")
	viper.SetDefault("apidocs.host", "localhost:4869")
	viper.SetDefault("apidocs.basepath", "/")
	viper.SetDefault("apidocs.schemes", []string{"http"})

	viper.SetDefault("basic_auth.username", "admin")
	viper.SetDefault("basic_auth.password", "admin")

	viper.SetDefault("app.chan_size", 50)

	// 初始化 sentry 并创建 sentry 客户端
	sentryDSN := viper.GetString("sentry.dsn")
	if sentryDSN == "" {
		sentryDSN = os.Getenv(logging.SentryDSNEnvKey)
	}
	sentryDebug := viper.GetBool("sentry.debug")
	if viper.GetString("server.mode") == "release" {
		sentryDebug = false
	}
	logging.Debug(nil, "Sentry use dns: "+sentryDSN)
	sentry, err := logging.NewSentryClient(sentryDSN, sentryDebug)
	if err != nil {
		logging.Error(nil, "Sentry client create error:"+err.Error())
	}

	// 根据配置创建 logging 的 logger 并将 logging 的默认 logger 替换为当前创建的 logger
	outputs := viper.GetStringSlice("logging.output_paths")
	var lumberjackSink *logging.LumberjackSink
	for _, output := range outputs {
		if strings.HasPrefix(output, "logrotate://") {
			filename := strings.Split(output, "://")[1]
			maxAge := viper.GetInt("logging.logrotate.max_age")
			maxBackups := viper.GetInt("logging.logrotate.max_backups")
			maxSize := viper.GetInt("logging.logrotate.max_size")
			compress := viper.GetBool("logging.logrotate.compress")
			localtime := viper.GetBool("logging.logrotate.localtime")
			lumberjackSink = logging.NewLumberjackSink("logrotate", filename, maxAge, maxBackups, maxSize, compress, localtime)
		}
	}
	logger, err := logging.NewLogger(logging.Options{
		Level:             viper.GetString("logging.level"),
		Format:            viper.GetString("logging.format"),
		OutputPaths:       outputs,
		DisableCaller:     viper.GetBool("logging.disable_caller"),
		DisableStacktrace: viper.GetBool("logging.disable_stacktrace"),
		AtomicLevelServer: logging.AtomicLevelServerOption{
			Addr:     viper.GetString("logging.atomic_level_server.addr"),
			Path:     viper.GetString("logging.atomic_level_server.path"),
			Username: viper.GetString("basic_auth.username"),
			Password: viper.GetString("basic_auth.password"),
		},
		SentryClient:   sentry,
		LumberjackSink: lumberjackSink,
	})
	if err != nil {
		logging.Error(nil, "Logger create error:"+err.Error())
	} else {
		logging.ReplaceLogger(logger)
	}
}

// DefaultGinMiddlewares 默认的 gin server 使用的中间件列表
func DefaultGinMiddlewares() []gin.HandlerFunc {
	m := []gin.HandlerFunc{
		// 记录请求处理日志，最顶层执行
		webserver.GinLogMiddleware(),
		// 捕获 panic 保存到 context 中由 GinLogger 统一打印， panic 时返回 500 JSON
		webserver.GinRecovery(response.Respond),
	}

	// 配置开启请求限频则添加限频中间件
	if viper.GetBool("ratelimiter.enable") {
		m = append(m, webserver.GinRatelimitMiddleware())
	}
	return m
}
