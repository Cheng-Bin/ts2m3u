package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type TSFile struct {
	Format Format `json:format`
}

type Format struct {
	Filename string `json:filename`
	Duration string `json:duration`
}

func ExecMul(datas []string) {
	//files = make([]TSFile, 0, 10)
	sig := make(chan TSFile, 100)
	done := make(chan bool)

	length := 25
	for i := 0; i < length; i++ {
		go Execute(datas[i], sig, done, i, length)
	}

	var flag bool
	for {
		select {
		case result := <-sig:
			fmt.Println(result.Format.Filename)
		case flag = <-done:
			close(sig)
			close(done)
			break
		}

		if flag {
			break
		}
	}

}

func Execute(file string, result chan TSFile, done chan bool, cur int, length int) {
	cmd := exec.Command("ffprobe", "-print_format", "json", "-show_format", "-show_streams", "-i", file)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	checkError(err)
	outStr := out.String()

	tsFile := TSFile{}
	err = json.Unmarshal([]byte(outStr), &tsFile)

	checkError(err)
	tsFile.Format.Filename = file

	result <- tsFile

	if cur == length-1 {
		done <- true
	}
}
