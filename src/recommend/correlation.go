package recommend

import (
	"os"
	"fmt"
	"strconv"
	"strings"
	"math"
	"bufio"
)

type MapSlices map[int64][]int64

func CreateFileForMapSlice(file string, tmps map[int64][]int64) {


	hrFile, err := os.Create(getFilePath(file))
	defer hrFile.Close()
	if err != nil {
		fmt.Println(file, err)
		return
	}
	for k, v := range tmps {
		hrFile.WriteString(strconv.FormatInt(k, 10) + "\t" + JoinForInt64(v, "\t") + "\n")
	}

}

func CreateFileForMapMap(file string, tmps map[int64]map[int64]float64) {




	hrFile, err := os.Create(getFilePath(file))
	defer hrFile.Close()
	if err != nil {
		fmt.Println(file, err)
		return
	}
	for k, v := range tmps {
		var line string = ""
		for resume,score:=range v{
			line+="\t"
			line+=strconv.FormatInt(resume,10)
			line+=","
			line+=strconv.FormatFloat(score,'f',3,64)
		}
		hrFile.WriteString(strconv.FormatInt(k, 10) + line + "\n")
	}

}

func CreateFileForMapMap2(file string, tmps map[int64]map[int64]float64) {

	hrFile, err := os.Create(getFilePath(file))
	w := bufio.NewWriter(hrFile)
	defer hrFile.Close()
	if err != nil {
		fmt.Println(file, err)
		return
	}
	for k, v := range tmps {
		var line string = ""
		for resume,score:=range v{
			line+="\t"
			line+=strconv.FormatInt(resume,10)
			line+=","
			line+=strconv.FormatFloat(score,'f',3,64)
		}
		w.WriteString(strconv.FormatInt(k, 10) + line + "\n")
	}
	w.Flush()
}


func BuildStructure(file string,sep string,aIndex int,bIndex int) (MapSlices, MapSlices) {
	//"material/20140822110003_195910.txt"
	aToBMap := make(MapSlices)
	bToAMap := make(MapSlices)

	start := 0
	FileOpen(file, func(position int, line string) () {
			tmps := strings.Split(line, sep)
			if len(tmps) < 2 || position == 0 {
				return
			}
			aId, _ := strconv.ParseInt(tmps[aIndex], 10, 64)
			bId, _ := strconv.ParseInt(strings.TrimSpace(tmps[bIndex]), 10, 64)
			if _, ok := aToBMap[aId]; ok {
				if !ContainForInt64(aToBMap[aId], bId) {
					aToBMap[aId] = append(aToBMap[aId], bId)
				}
			}else {
				aToBMap[aId] = []int64{bId}
			}
			if _, ok := bToAMap[bId]; ok {
				if !ContainForInt64(bToAMap[bId], aId) {
					bToAMap[bId] = append(bToAMap[bId], aId)
				}
			}else {
				bToAMap[bId] = []int64{aId}
			}
			start++
		})
	return aToBMap, bToAMap
}

func Calculate(bToaMap MapSlices,aToBMap MapSlices) map[int64]map[int64]float64{
	scoreMap:=make(map[int64]map[int64]float64)
	for k,v:=range bToaMap{
		if len(v) == 1{
			continue
		}

		correlateMap:=make(map[int64]int)
		for _,hr:=range v{

			for _,resume:=range aToBMap[hr]{
				if resume == k{
					continue
				}
				if _,ok:=correlateMap[resume];ok{
					correlateMap[resume]+=1
				}else{
					correlateMap[resume]=1
				}
			}
		}
		curr:=len(bToaMap[k])

		for resume,inter:=range correlateMap{
			if inter>1{
				if _,ok:=scoreMap[k];!ok{
					scoreMap[k]=make(map[int64]float64)
				}
				count:=len(bToaMap[resume])
				scoreMap[k][resume]=float64(inter)/(math.Pow(float64(curr),0.5)+math.Pow(float64(count),1-0.5))
			}
		}
	}
	return scoreMap
}

func ( mainMap MapSlices) Merge(tmp MapSlices) {
	for k,_ := range tmp{
		if _,ok:= mainMap[k];ok{
			mainMap[k]=append(mainMap[k],tmp[k]...)
		}else{
			mainMap[k]=tmp[k]
		}
	}
}


