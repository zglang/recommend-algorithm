package recommend

import (
	"io"
	"os"
	"bufio"
	"fmt"
	"bytes"
	"encoding/binary"
	"strconv"
)

func Substring(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl-1+start
	}
	end = start+length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func Read(path string) string {

	f , err := os.Open(getFilePath(path))
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	br := bufio.NewReader(f)
	content := ""
	for {
		line , err := br.ReadString(byte('\n'))
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		content+=line
	}
	return content
}

func binSearch(items []rune, item rune) bool {
	var low, mid int = 0, 0
	hight := len(items) - 1
	for low <= hight {
		mid = (low+hight)/2
		if items[mid] == item {
			return true
		} else if items[mid] > item {
			hight = mid-1
		} else {
			low = mid+1
		}
	}
	return false
}

func ContainForInt64(items []int64, item int64) bool {
	for _,v := range items{
		if v == item{
			return true
		}
	}
	return false
}

func JoinForInt64(a []int64, sep string) string {
	var tmp string = ""
	for i:=0;i<len(a);i++{
		if i>0{
			tmp+=sep
		}
		tmp+=strconv.FormatInt(a[i],10)
	}
	return tmp
}

func clearTagSuffix(tag []rune) []rune {
	end := len(tag) - 1
	for end >= 0 {
		if tag[end] == 10 || tag[end] == 13 || tag[end] == 0 {
			end--
		}else {
			break
		}
	}
	if end >= 0 {
		return tag[0:end+1]
	}else {
		return tag[0:0]
	}
}

func ClearValueSuffix(val []rune) []rune {
	end := len(val) - 1
	for end >= 0 {
		if val[end] == 124 || val[end] == 10 || val[end] == 13 || val[end] == 0 || val[end] == 32 {
			end--
		}else {
			break
		}
	}
	if end >= 0 {
		return val[0:end+1]
	}else {
		return val[0:0]
	}
}

func buildNewTag(tag []rune) []rune {
	newTag := make([]rune, len(tag)*2)
	for i, j := 0, 0; i < len(tag); i, j = i+1, j+1 {
		if tag[i] != 58 {
			newTag[j] = 32
			j++
		}
		newTag[j] = tag[i]
	}
	return clearTagSuffix(newTag)
}

var replaceMap map[rune]rune = map[rune]rune{12288:32, 65306:58, 65372:124, 12290:46}

func formatContent(content []rune) []rune {
	newContent := make([]rune, len(content))
	j := 0
	for i := 0; i < len(content); i++ {
		if item, ok := replaceMap[content[i]]; ok {
			newContent[j] = item
		}else {
			newContent[j] = content[i]
		}
		if newContent[j] == 32 && j > 0 {
			if newContent[j-1] == 32 {
				j--
			}
		}
		j++
	}
	return newContent
}

func SplitBytes(content []byte, sep byte) [3][]byte {
	var tmps [3][]byte
	start, position := 0, 0
	for i := 0; i < len(tmps); i++ {
		for position < len(content) && content[position] != 9 {
			position++
		}
		if content[position-1] == 10 {
			tmps[i] = content[start:position-1]
		}else {
			tmps[i] = content[start:position]
		}

		position++
		start = position
	}
	return tmps
}


func ContainBytes(content []byte, sub []byte) bool {
	math := true
	position := 0
	for position < len(content) {

		for i := 0; i < len(sub); i++ {
			if sub[i] == content[position] {
				position++
				math = true
			}else {
				math = false
				break
			}
		}
		for !math && position < len(content) {
			if content[position] == 9 {
				position++
				break
			}
			position++
		}

	}
	return math
}

func BytesToint64(tmps []byte) int64{
	b_buf := bytes.NewBuffer(tmps)
	var x int64
	binary.Read(b_buf, binary.BigEndian, &x)
	return x
}

func BytesToInt(tmps []byte) int{
	fmt.Println(tmps)
	b_buf := bytes.NewBuffer(tmps)
	var x int
	binary.Read(b_buf, binary.BigEndian, &x)
	fmt.Println(x)
	return x
}

func Int64ToBytes(tmp *int64) []byte{
	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, tmp)
	return b_buf.Bytes()
}

//func clearContent(content []rune) []rune{
//
//}
