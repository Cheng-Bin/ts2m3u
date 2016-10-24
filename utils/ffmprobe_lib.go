package utils

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

type TSFile struct {
	Format Format `json:format`
}

type Format struct {
	Filename string `json:filename`
	Duration string `json:duration`
}

func ExecMul(datas []string) (tss []TSFile) {

	length := len(datas)
	for i := 0; i < length; i++ {
		result := Execute(datas[i])
		tss = append(tss, result)
	}

	return
}

func Execute(file string) TSFile {
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

	return tsFile
}
