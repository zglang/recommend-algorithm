package recommend

import (
	"fmt"
)

func ExecuteContainFile() {
	//	aToBMap,bToAMap:=BuildStructure("material/2014-08-28.txt","^",0,2)
	//	result:=Calculate(bToAMap,aToBMap)
	//	CreateFileForMapSlice("material/aToBMap.txt",aToBMap)
	//	CreateFileForMapSlice("material/bToAMap.txt",bToAMap)
	//	CreateFileForMapMap("material/score2.txt",result)
}

var pathFiles []string = []string{
//	"material/2014-08-22.txt",
//	"material/2014-08-23.txt",
//	"material/2014-08-24.txt",
//	"material/2014-08-25.txt",
	"material/2014-08-26.txt",
	"material/2014-08-27.txt",
	"material/2014-08-28.txt"}

func Execute() {

	aToBMap, bToAMap := make(MapSlices), make(MapSlices)

	//index:=0
	for _, path := range pathFiles {
		tmpAToBMap, tmpBToAMap := BuildStructure(path, "^", 0, 2)
		//		if index == 0 {
		//			aToBMap=tmpAToBMap
		//			bToAMap=tmpBToAMap
		//		}else{
		aToBMap.Merge(tmpAToBMap)
		bToAMap.Merge(tmpBToAMap)
		//		}
		//index ++
	}
	fmt.Println("aToBMap,bToAMap=", len(aToBMap), ",", len(bToAMap))



	//	aToBMap1,bToAMap1:=BuildStructure("material/2014-08-26.txt","^",0,2)
	//	aToBMap2,bToAMap2:=BuildStructure("material/2014-08-27.txt","^",0,2)
	//	aToBMap3,bToAMap3:=BuildStructure("material/2014-08-28.txt","^",0,2)
	//
	//	aToBMap1.Merge(aToBMap2)
	//	bToAMap1.Merge(bToAMap2)
	//	aToBMap1.Merge(aToBMap3)
	//	bToAMap1.Merge(bToAMap3)
	//
	result := Calculate(bToAMap, aToBMap)
	fmt.Println("result.len=",len(result))
	//CreateFileForMapMap2("material/score2.txt", result)
}
