package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"github.com/mymmsc/gox/api"
	"github.com/mymmsc/gox/logger"
	"github.com/quant1x/quant/category"
	"github.com/quant1x/quant/data/security"
	"io"
	"log"
	"os"
	"os/user"
	"strings"
	"syscall"
)

const (
	tdx_path = "/workspace/data/tdx"
	blk_path = "/T0002/blocknew"
	//blk_filename = "zxg.blk"
	blk_filename = "BKLT.blk"
)

func TdxZxg() {
	//
}

func main() {
	var (
		path string // 数据路径
	)
	flag.StringVar(&path, "path", "", "通达信安装目录")

	if len(path) == 0 {
		u, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Home dir:", u.HomeDir+tdx_path)
		path = u.HomeDir + tdx_path
	}
	filename := path + blk_path + "/" + blk_filename

	f, err := os.Open(filename)
	if err != nil {
		//ENOENT (2)
		if errors.Is(err, syscall.ENOENT) {
			logger.Debugf("自选股[%s]: K线数据文件不存在", filename)
			return
		} else {
			logger.Errorf("自选股[%s]: K线数据文件操作失败:%v", filename, err)
			return
		}
	}

	/*r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)*/

	//var data []byte
	data, err := io.ReadAll(f)
	if err != nil {
		logger.Errorf("自选股[%s]: K线数据文件操作失败:%v", filename, err)
	}

	s := string(data)
	fmt.Println("%s\n", s)
	arr := strings.Split(s, "\r\n")
	// 深圳指数(0, ‘399001’)，上海大盘 (1, ‘999999’)。
	// 数据在’ZXG.blk’中以8个字节来存放。
	fmt.Printf("%v\n", arr)
	fcsv, _ := os.OpenFile("zxg.csv", os.O_RDWR|os.O_CREATE, category.CACHE_FILE_MODE)
	defer api.CloseQuietly(fcsv)
	//fcsv.WriteString("\xEF\xBB\xBF")
	out := csv.NewWriter(fcsv)
	var header = []string{"market", "code", "name"}
	out.Write(header)
	for _, d := range arr {
		d = strings.TrimSpace(d)
		if len(d) != 7 {
			continue
		}
		market := d[:1]
		code := d[1:]
		fmt.Printf("市场编码:%s, 证券代码:%s\n", market, code)
		fullCode := ""
		code = strings.TrimSpace(code)
		if market == "1" {
			market = "上海"
			fullCode = "sh" + code
		} else if market == "0" {
			market = "深圳"
			fullCode = "sz" + code
		} else {
			continue
		}
		info, err := security.GetBasicInfo(fullCode)
		if err != nil {
			fmt.Printf("没有找到 %s\n", fullCode)
			continue
		}
		row := []string{market, fullCode, info.Name}
		out.Write(row)
		out.Flush()
	}
	fcsv.Close()
}
