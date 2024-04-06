package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/samuel/go-zookeeper/zk"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

type Dataset struct {
	XMLName  xml.Name `xml:"dataset"` //将元素名写入该字段
	Metadata Metadata `xml:"metadata"`
}

type Metadata struct { //后面的内容是struct tag，标签，是用来辅助反射的
	Columns []Col `xml:"column"`
}

type Col struct {
	Name string `xml:"name,attr"`
}

type Table struct {
	Name    string             `yaml:"name"`
	Type    string             `yaml:"type"`
	Columns map[string]*Column `yaml:"columns"`
}

type Column struct {
	caseSensitive bool   `yaml:"caseSensitive"`
	dataType      int    `yaml:"dataType"`
	generated     bool   `yaml:"generated"`
	name          string `yaml:"name"`
	nullable      bool   `yaml:"nullable"`
	primaryKey    bool   `yaml:"primaryKey"`
	unsigned      bool   `yaml:"unsigned"`
	visible       bool   `yaml:"visible"`
}

// select * from information_schema.columns where lower(table_name) = 'views';

func main1() {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/information_schema"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面

	fp := "/Users/dongzonglei/source_code/Github/go-study/zoo/information_schema/"

	files, _ := os.ReadDir(fp)

	for _, file := range files {
		fullName := path.Base(fp + file.Name())
		fileExt := path.Ext(fp + file.Name())
		filePrefix := strings.TrimSuffix(fullName, fileExt)
		//fmt.Printf("file name is %s\n", filePrefix)

		rows, err := db.Query("select column_name from information_schema.columns where lower(table_name) = ?", filePrefix)

		v := &Dataset{}

		for rows.Next() {
			var name string
			if err := rows.Scan(&name); err != nil {
				log.Fatal(err)
			}
			v.Metadata.Columns = append(v.Metadata.Columns, Col{Name: strings.ToLower(name)})
			//fmt.Printf("column name is %s\n", name)
		}

		output, err := xml.MarshalIndent(v, "  ", "    ")
		if err != nil {
			fmt.Printf("error : %v\n", err)
		}
		fmt.Println()
		f, err := os.OpenFile(fp+"dataset/"+filePrefix+".yaml", os.O_RDONLY|os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		_, err = f.WriteString(string(output))
		if err != nil {
			log.Println("write file error :", err)
			return
		}
	}
}

func main2() {
	dataBytes, err := os.ReadFile("/Users/dongzonglei/source_code/Github/go-study/zoo/information_schema/administrable_role_authorizations.yaml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("yaml 文件的内容: \n", string(dataBytes))

	table := Table{}
	err = yaml.Unmarshal(dataBytes, &table)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
	fmt.Printf("config → %+v\n", table)

	v := &Dataset{}

	for key, _ := range table.Columns {
		v.Metadata.Columns = append(v.Metadata.Columns, Col{Name: key})
	}

	//v.Columns = append(v.Columns, Column1{Name: "user"})
	//v.Columns = append(v.Columns, Column1{Name: "host"})
	//v.Columns = append(v.Columns, Column1{Name: "ip"})
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	//os.Stdout.Write([]byte(xml.Header)) //输出预定义的xml头  <?xml version="1.0" encoding="UTF-8"?>
	os.Stdout.Write(output)
}

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

	export(conn, "/governance_ds/metadata/sys_db/schemas/sys_db/tables", "sys")
	export(conn, "/governance_ds/metadata/performance_schema_db/schemas/performance_schema_db/tables", "performance_schema")
	export(conn, "/governance_ds/metadata/mysql_db/schemas/mysql_db/tables", "mysql")
	export(conn, "/governance_ds/metadata/information_schema_db/schemas/information_schema_db/tables", "information_schema")
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
