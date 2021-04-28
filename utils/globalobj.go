// Package utils 提供zinx相关工具类函数
// 包括:
//		全局配置
//		配置文件加载
//
// 当前文件描述:
// @Title  globalobj.go
// @Description  相关配置文件定义及加载方式
// @Author  Aceld - Thu Mar 11 10:32:29 CST 2019
package utils

import (
	"os"
)

/*
	存储一切有关Zinx框架的全局参数，供其他模块使用
	一些参数也可以通过 用户根据 zinx.json来配置
*/
type GlobalObj struct {
	/*
		Server
	*/
	// TCPServer ziface.IServer //当前Zinx的全局Server对象
	ServerId   string `json:"server_id"`   //当前服务器ID
	ServerType string `json:"server_type"` //当前服务器类型
	ServerName string `json:"server_name"` //当前服务器名称
	ServerIp   string `json:"server_ip"`   //当前服务器主机IP
	ClientPort int    `json:"client_port"` //当前服务器主机监听端口号

	/*
		Zinx
	*/
	Version          string `json:"version"`             //当前Zinx版本号
	MaxPacketSize    uint32 `json:"max_packet_size"`     //都需数据包的最大值
	MaxConn          int    `json:"max_conn"`            //当前服务器主机允许的最大链接个数
	WorkerPoolSize   uint32 `json:"worker_pool_size"`    //业务工作Worker池的数量
	MaxWorkerTaskLen uint32 `json:"max_worker_task_len"` //业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    uint32 `json:"max_msg_chan_len"`    //SendBuffMsg发送消息的缓冲最大长度

	/*
		config file path
	*/
	//ConfFilePath string `json:"conf_file_path"`

	/*
		logger
	*/
	LogDir        string `json:"log_dir"`         //日志所在文件夹 默认"./log"
	LogFile       string `json:"log_file"`        //日志文件名称   默认""  --如果没有设置日志文件，打印信息将打印至stderr
	LogDebugClose bool   `json:"log_debug_close"` //是否关闭Debug日志级别调试信息 默认false  -- 默认打开debug信息
}

/*
	定义一个全局的对象
*/
var GlobalObject *GlobalObj

//PathExists 判断一个文件是否存在
//func PathExists(path string) (bool, error) {
//	_, err := os.Stat(path)
//	if err == nil {
//		return true, nil
//	}
//	if os.IsNotExist(err) {
//		return false, nil
//	}
//	return false, err
//}
//
////Reload 读取用户的配置文件
//func (g *GlobalObj) Reload() {
//
//	if confFileExists, _ := PathExists(g.ConfFilePath); confFileExists != true {
//		//fmt.Println("Config File ", g.ConfFilePath , " is not exist!!")
//		return
//	}
//
//	data, err := ioutil.ReadFile(g.ConfFilePath)
//	if err != nil {
//		panic(err)
//	}
//	//将json数据解析到struct中
//	err = json.Unmarshal(data, g)
//	if err != nil {
//		panic(err)
//	}
//
//	//Logger 设置
//	if g.LogFile != "" {
//		zlog.SetLogFile(g.LogDir, g.LogFile)
//	}
//	if g.LogDebugClose == true {
//		zlog.CloseDebug()
//	}
//}

/*
	提供init方法，默认加载
*/
func InitGlobal(gb GlobalObj) {
	pwd, err := os.Getwd()
	if err != nil {
		pwd = "."
	}
	//初始化GlobalObject变量，设置一些默认值
	GlobalObject = &GlobalObj{
		ServerName:       "ZinxServerApp",
		Version:          "V0.11",
		ClientPort:       8999,
		ServerIp:         "0.0.0.0",
		MaxConn:          12000,
		MaxPacketSize:    4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
		MaxMsgChanLen:    1024,
		LogDir:           pwd + "/log",
		LogFile:          "",
		LogDebugClose:    false,
	}
	if len(gb.ServerName) > 0 {
		GlobalObject.ServerName = gb.ServerName
	}
	if gb.ClientPort > 0 {
		GlobalObject.ClientPort = gb.ClientPort
	}
	if len(gb.ServerIp) > 0 {
		GlobalObject.ServerIp = gb.ServerIp
	}
	if gb.MaxConn > 0 {
		GlobalObject.MaxConn = gb.MaxConn
	}
	if gb.MaxPacketSize > 0 {
		GlobalObject.MaxPacketSize = gb.MaxPacketSize
	}
	if gb.WorkerPoolSize > 0 {
		GlobalObject.WorkerPoolSize = gb.WorkerPoolSize
	}
	if gb.MaxWorkerTaskLen > 0 {
		GlobalObject.MaxWorkerTaskLen = gb.MaxWorkerTaskLen
	}
	if gb.MaxMsgChanLen > 0 {
		GlobalObject.MaxMsgChanLen = gb.MaxMsgChanLen
	}
	if len(gb.LogDir) > 0 {
		GlobalObject.LogDir = gb.LogDir
	}
	if len(gb.LogFile) > 0 {
		GlobalObject.LogFile = gb.LogFile
	}
	GlobalObject.LogDebugClose = gb.LogDebugClose
	//
	////NOTE: 从配置文件中加载一些用户配置的参数
	//GlobalObject.Reload()
}
