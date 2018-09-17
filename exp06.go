package main

import (
	"encoding/json"
	"log"
	"os"
)

var (
	OutputFile = "E:\\yinzhengjie.json"
)

type TenScenicSpots struct { //定义10个景区的名称。
	FirstScenic   string
	SecondScenic  string
	ThirdScenic   string
	FourthScenic  string
	FifrhScenic   string
	SixthScenic   string
	SeventhScenic string
	EigthtScenic  string
	NinthScenic   string
	TenthScenic   string
}

type TouristInformation struct { //定义游客信息
	VisitorName string            //游客姓名
	Nationality string            //游客国籍
	City        string            //想要去的城市
	ScenicSpot  []*TenScenicSpots //想要去看的景区

}

func main() {
	ChaoyangDistrict := &TenScenicSpots{"中华名族园", "北京奥林匹克公园", "国家体育馆", "中国科学技术官", "奥林匹克公园网球场", "蟹岛绿色生态农庄", "国家游泳中心（水立方）", "中国紫檀博物馆", "北京欢乐谷", "元大都城"}
	DaxingDistrict := &TenScenicSpots{"北京野生动物园", "男孩子麋鹿苑", "中华文化园", "留民营生态农场", "中国印刷博物馆", "北普陀影视城", "大兴滨河森林公园", "呀路古热带植物园", "庞各庄万亩梨园", "西黄垈村"}
	District := TouristInformation{"尹正杰", "中国", "北京", []*TenScenicSpots{ChaoyangDistrict, DaxingDistrict}}

	/* GolangJson, err := json.Marshal(District) //这个步骤就是序列化的过程。json.Marshal方法会返回一个字节数组，即GolangJson,与此同时，District已经是JSON格式的啦。
	   if err != nil {
	       log.Fatal("序列化报错是：%s", err)
	   }
	   fmt.Printf("JSON format: %s", GolangJson)
	*/
	file, _ := os.OpenFile(OutputFile, os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	Write := json.NewEncoder(file) //创建一个编码器。
	err := Write.Encode(District)  //由于District已经被json.Marshal方法处理过了，所以我们直接把JSON格式的District传给Write写入器,调用该写入器的Encode方法可以对JSON格式的数据进行编码。如果顺利的话，我们会得到一个nil参数，否则我们会得到编码的错误信息。
	if err != nil {
		log.Println("Error in encoding json")
	}
}
