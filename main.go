package main

import (
	"fmt"
	"net/http"
	"reflect"
)

//自定义字符串型错误
type StringErr string

func (s StringErr) Error() string {
	return string(s)
}

var httppath = []string{
	"/API/realTimePowerdata4wind.jsp",//5.1	实时功率数据
	"/API/shortpredictionPowerdata4wind.jsp",//5.2	短期预测功率接口
	"/API/ultrashortpredictionPowerdata4wind.jsp",//5.3	超短期预测功率接口
	"/API/windPower.jsp",//5.4	测风塔实时数据
	"/API/preditctionData4wind.jsp",//5.5	气象预测数据（暂时不开发）
	"/API/WindPowerGeneratorStatus.jsp",//5.6	风机运行状态
	"/API/WindPoweCheckPlan.jsp",//5.7	检修计划
	"/API/WindPowerRestrictedPlan.jsp",//5.8	限电记录
	"/API/WindPowerGenerationPlan.jsp",//5.9	发电计划
	"/API/realTimePowerdata4photovoltaic.jsp",//6.1	实时功率数据
	"/API/shortpredictionPowerdata4photovoltaic.jsp",//6.2	短期预测功率接口
	"/API/ultrashortpredictionPowerdata4photovoltaic.jsp",//6.3	超短期预测功率接口
	"/API/observationData4photovoltaic.jsp",//6.4	实时观测数据数据
	"/API/preditctionData4photovoltaic.jsp",//6.5	气象预测数据（暂时不开发）
	"/API/InverterStatus.jsp",//6.6	逆变器运行状态
	"/API/PhotovoltaicCheckPlan.jsp",//6.7	检修计划
	"/API/PhotovoltaicRestrictedPlan.jsp",//6.8	限电记录
	"/API/PhotovoltaicGenerationPlan.jsp",//6.9	发电计划
}

var params = [2]string{
	"SiteID",
	"Time",
}

func getPathIndex(path string) (int, error) {
	var inx int = -1
//	for i := 0; i < len(httppath); i++ {
	for i,url := range httppath{
		if path == url {
			inx = i
			break
		}
	}
	if inx == -1 {
		fmt.Println("path not found ", path)
		return -1, StringErr("error path")
	}
	return inx, nil
}

func parseParam(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数, 默认是不会解析的
	//fmt.Println(r.Form) //这些是服务器端的打印信息
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])

	//fmt.Println("method", r.Method)
	//fmt.Println("uri", r.RequestURI)
	fmt.Println("path", r.URL.Path)
	//for k, v := range r.Form {
	//	fmt.Println("key:", k)
	//	fmt.Println("val:", strings.Join(v, ""))
	//}
	inx, err := getPathIndex(r.URL.Path)
	if err != nil {
		return
	}
	fmt.Println("inx:", inx)

	randat := GetRandomData(inx, r.Form.Get("SiteID"), r.Form.Get("Time"))
	fmt.Println(string(randat))
	fmt.Fprintf(w, string(randat))
	/*
		dat,err:=GetJsonFromFile(uint(inx))
		if err == nil{
			fmt.Fprintf(w, string(dat))
		}
	*/
}

type ABC struct {
	Str   string
	Id    int
	Count uint32
	Data  float64
}

//给数据结构的成员赋值，记得大写。。。
func RandomStruct(p interface{}) {
	t := reflect.TypeOf(p)
	if t.Kind() != reflect.Ptr {
		return
	}
	t = t.Elem()
	if t.Kind() != reflect.Struct {
		return
	}

	fmt.Println("kind ", t.Kind())
	v := reflect.ValueOf(p).Elem()
	fmt.Println(v.CanSet(), v.Field(0))
	fmt.Println(v.Kind(), v.Type(), v.NumField(), v.Interface())
	v.Field(0).SetString("haha")

	n := v.NumField()
	for i := 0; i < n; i++ {
		if !v.Field(i).CanSet() {
			fmt.Println("field ", t.Field(i).Name, t.Field(i).Type, "can not be set")
			continue
		}
		switch v.Field(i).Kind() {
		case reflect.String:
			v.Field(i).SetString("haha")
		case reflect.Int:
			v.Field(i).SetInt(1234)
		case reflect.Float64:
			v.Field(i).SetFloat(0.12345)
		case reflect.Uint32:
			v.Field(i).SetUint(123)
		}
	}

	fmt.Println(v.Interface())
}

func main() {
	//test()
	//fmt.Println("test...")
	//return
	fmt.Println("go http running")
	//return
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/API/", parseParam)
	http.ListenAndServe(":8080", nil)
}
