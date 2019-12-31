package handler

import "net/http"

// Student represent information of a student
type Student struct {
	Name       string // 姓名
	ID         string // 性别
	Department string // 院系
	Class      string // 班级
	Level      string // 学历层次
	Mod        string // 学习形式
	Graduation string // 毕业时间
}

// ManageHandler handles manage request of student information
var ManageHandler = makeHandler(manageHandler)

func manageHandler(w http.ResponseWriter, req *http.Request) {

}

var students = []*Student{
	&Student{"张三", "311698555977", "计算机学院", "1703", "本科", "四年制", "2021-7-6"},
	&Student{"李四", "311698555977", "计算机学院", "1703", "本科", "四年制", "2021-7-6"},
	&Student{"王五", "311698555977", "计算机学院", "1703", "本科", "四年制", "2021-7-6"},
	&Student{"赵田", "311698555977", "计算机学院", "1703", "本科", "四年制", "2021-7-6"},
	&Student{"范德彪", "311698555977", "计算机学院", "1703", "本科", "四年制", "2021-7-6"},
}
