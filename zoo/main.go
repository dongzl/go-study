package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"os"
	"time"
)

func main() {
	// 创建zk连接地址
	hosts := []string{"127.0.0.1:2181"}
	// 连接zk
	conn, _, err := zk.Connect(hosts, time.Second*5)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	println(conn.Server())

	export(conn, "/governance_ds/metadata/sys/schemas/sys/tables", "sys")
	export(conn, "/governance_ds/metadata/performance_schema/schemas/performance_schema/tables", "performance_schema")
	export(conn, "/governance_ds/metadata/mysql/schemas/mysql/tables", "mysql")
	export(conn, "/governance_ds/metadata/information_schema/schemas/information_schema/tables", "information_schema")
}

func export(conn *zk.Conn, path string, dir string) {
	data, _, err := conn.Children(path)
	if err != nil {
		fmt.Printf("查询%s失败, err: %v\n", path, err)
		return
	}
	fmt.Printf("%s 的值为 %s\n", path, len(data))
	for i := 0; i < len(data); i++ {
		write(conn, path, data[i], dir)
	}
}

func write(conn *zk.Conn, path string, child string, dir string) {
	fmt.Printf("%s 的子节点为 %s\n", path, child)
	leaf, _, _ := conn.Get(path + "/" + child + "/versions/0")
	//fmt.Printf("%s 的值为 %s\n", path, string(leaf))
	f, err := os.OpenFile("zoo/"+dir+"/"+child+".yaml", os.O_RDONLY|os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	defer f.Close()
	if err != nil {
		log.Println("open file error :", err)
		return
	}
	_, err = f.WriteString(string(leaf))
	if err != nil {
		log.Println("write file error :", err)
		return
	}
}
