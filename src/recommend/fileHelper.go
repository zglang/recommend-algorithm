package recommend

import (
	"os"
	"bufio"
	"fmt"
	"io"
)

func FileOpen(file string,process func(int,string)) {
	//"material/download.txt"
	index:=0
	path:=getFilePath(file)
	fmt.Println("打开文件按:"+path)

	f , err := os.Open(path)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	br := bufio.NewReader(f)
	for {
		line , err := br.ReadString(byte('\n'))
		if err == io.EOF {
			break
		}else {
			process(index,line)
		}
		index++
	}
}

func getFilePath(fileName string) string {
	basePath, err := os.Getwd()
	if err != nil {
		return fileName
	}
	filePath := basePath + "/" + fileName
	return filePath
}
