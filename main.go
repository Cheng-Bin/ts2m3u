package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chapin/ts2m3u/utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("请输入视频文件夹名称")
	}

	dirPath := os.Args[1]

	files, err := utils.GetFiles(dirPath)
	check(err)

	tsFiles := utils.ExecMul(files)

	m3u8Writer := utils.NewM3U8(true)
	for _, item := range tsFiles {
		m3u8Writer.WriteTs(item.Format.Duration, item.Format.Filename)
	}
	m3u8Writer.WriterEnd()
	m3u8Data := m3u8Writer.GetM3U8()

	fmt.Println(m3u8Data)
	utils.SaveM3U8("vod", m3u8Data)
}
