package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var WindPlantCodeToId = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	//"YNKMLB":   1,
	//"YNHDXMSA": 2,
	//"YNKMSSP":  3,
	//"YNKMYN":   4,
}

var LightPlantCodeToId = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	//"GFDC1": 1,
	//"GFDC2": 2,
	//"GFDC3": 3,
	//"GFDC4": 4,
}

func PacketJson(mData DataBase) ([]byte, error) {
	b, err := json.Marshal(mData)
	if err != nil {
		fmt.Println("json.Marshal :", mData, " error:", err)
	}
	return b, err
}

func (d WindRealData) Random(plantid int, t uint64) DataBase {
	var data WindRealData
	data.Time = t
	data.GridPower = uint8(rand.Uint32())
	data.OperatingCapacity = uint8(rand.Uint32())
	data.AccumulativePower = rand.Float64() * 1000
	data.Restricted = false
	return data
}

func (d WindTurbineStatus) Random(plantid int, t uint64) DataBase {
	var data WindTurbineStatus
	//每一个风电场对应的风机id，这里每个风电场(1,2,3,4)只有一个风机(1,2,3,4)
	mapWindTurbinePlantId := map[int][]int{1: {1}, 2: {2}, 3: {3}, 4: {4}}
	TurId, ok := mapWindTurbinePlantId[plantid]
	if !ok {
		return data
	}
	data.Time = t

	fmt.Println("TurId", TurId)
	data.Status = make([]WindTurbineStatusItem, len(TurId), len(TurId))
	for i, id := range TurId {
		fmt.Println("i", i)
		data.Status[i].ID = id
		data.Status[i].ActivePower = rand.Float64() * 1000
		data.Status[i].ReactivePower = rand.Float64() * 1000
		data.Status[i].Windspeed = uint16(rand.Uint32())
		data.Status[i].WindDirection = uint16(rand.Uint32() % 360)
		data.Status[i].Status = uint8(rand.Uint32() % 9)
	}
	return data
}

func (d WindMeterTowerData) Random(plantid int, t uint64) DataBase {
	var data WindMeterTowerData
	//每一个风电场对应的测风塔id，这里每个风电场(1,2)有六个测风塔(1,2,3,4,5,6)
	//每个测风塔有6层
	mapWindMeterTowerPlantId := map[int][]int{1: {1, 2, 3, 4, 5, 6}, 2: {7, 8, 9, 10, 11, 12}}
	Towid, ok := mapWindMeterTowerPlantId[plantid]
	if !ok {
		fmt.Println("not found meter tower id , maybe this wind plant have none meter tower")
		return data
	}

	data.Time = t

	data.WindTower = make([][][2]int, len(Towid), len(Towid))
	for i, _ := range Towid {
		data.WindTower[i] = make([][2]int, 6, 6)
		for j := 0; j < 6; j++ {
			data.WindTower[i][j][0] = rand.Intn(100)
			data.WindTower[i][j][1] = rand.Int() % 360
		}
	}
	return data
}

func (d LightRealData) Random(plantid int, t uint64) DataBase {
	var data LightRealData
	data.Time = t
	data.GridPower = uint8(rand.Uint32())
	data.OperatingCapacity = uint8(rand.Uint32())
	data.AccumulativePower = rand.Float64() * 1000
	data.Restricted = false
	return data
}

func (d LightObservationData) Random(plantid int, t uint64) DataBase {
	var data LightObservationData
	data.Time = t
	data.Temperature = uint16(rand.Uint32())
	data.Humidity = uint16(rand.Uint32())
	data.Pressure = rand.Float64()
	data.Windspeed = uint16(rand.Uint32())
	data.WindDirection = uint16(rand.Uint32())
	data.TotalRadiance = rand.Float64()
	data.DirectRadiation = rand.Float64()
	data.DiffuseRadiation = rand.Float64()
	data.Componenttemperature = uint16(rand.Uint32())
	return data
}

func (d InverterStatus) Random(plantid int, t uint64) DataBase {
	var data InverterStatus
	//每一个光伏电场对应的逆变器id，这里每个风电场(1,2,3,4)只有一个逆变器(1,2,3,4)
	mapLightInverterPlantId := map[int][]uint16{1: {1}, 2: {2}, 3: {3}, 4: {4}}
	InvId, ok := mapLightInverterPlantId[plantid]
	if !ok {
		return data
	}
	data.Time = t
	data.Status = make([]InverterStatusItem, len(InvId), len(InvId))
	for i, id := range InvId {
		data.Status[i].ID = id
		data.Status[i].ActivePower = rand.Float64() * 1000
		data.Status[i].ReactivePower = rand.Float64() * 1000
		data.Status[i].Status = uint8(rand.Uint32()%9 + 1)
	}
	return data
}

func (d ShortpredictionPowerdata) Random(plantid int, t uint64) DataBase {
	var data ShortpredictionPowerdata
	//未来72小时预测功率
	data.Time = t
	data.ShortPowerPrediction = make([]float64, 72, 72)
	for i := 0; i < 72; i++ {
		data.ShortPowerPrediction[i] = rand.Float64() * 100
	}
	return data
}

func (d UltraShortpredictionPowerdata) Random(plantid int, t uint64) DataBase {
	var data UltraShortpredictionPowerdata
	//超短期功率预测,未来4小时预测功率
	data.Time = t
	data.UltraShortPowerPrediction = make([]float64, 72, 72)
	for i := 0; i < 72; i++ {
		data.UltraShortPowerPrediction[i] = rand.Float64() * 100
	}
	return data
}

func (d WindCheckPlan) Random(plantid int, t uint64) DataBase {
	var data WindCheckPlan
	//检修计划
	data.Time = t
	data.Status = make([]CheckPlanStatus, 4, 4)

	i := 0
	for _, v := range WindPlantCodeToId {
		data.Status[i].ID = v
		data.Status[i].IsTemporary = false
		data.Status[i].Capability = rand.Float64() * 100
		data.Status[i].CheckingDate = time.Now().Format("20060102150405")
		data.Status[i].Proposer = GetRandomSalt()
		data.Status[i].Approver = GetRandomSalt()
		data.Status[i].StartTime = uint64(time.Now().Unix() * 1000)
		data.Status[i].EndTime = uint64(time.Now().Unix() * 1000)

		i++
	}
	return data
}

func (d WindRestrictedPlan) Random(plantid int, t uint64) DataBase {
	var data WindRestrictedPlan
	//检修计划
	data.Time = t
	data.Status = make([]RestrictedPlanStatus, 4, 4)

	i := 0
	for _, v := range WindPlantCodeToId {
		data.Status[i].ID = v
		data.Status[i].BeforeRestriction = false
		data.Status[i].AfterRestriction = false
		data.Status[i].RestrictionPower = rand.Float64() * 100
		data.Status[i].RestrictionPoweVolum = rand.Float64() * 100
		data.Status[i].Reason = GetRandomSalt()
		data.Status[i].Memo = GetRandomSalt()
		data.Status[i].StartTime = uint64(time.Now().Unix() * 1000)
		data.Status[i].EndTime = uint64(time.Now().Unix() * 1000)

		i++
	}
	return data
}



func (d WindGenerationPlan) Random(plantid int, t uint64) DataBase {
	var data WindGenerationPlan
	//检修计划
	data.Time = t
	data.Status = make([]GenerationPlanStatus, 4, 4)

	i := 0
	for _, v := range WindPlantCodeToId {
		data.Status[i].ID = v
		data.Status[i].BeforeRestriction = false
		data.Status[i].AfterRestriction = false
		data.Status[i].RestrictionPower = rand.Float64() * 100
		data.Status[i].RestrictionPoweVolum = rand.Float64() * 100
		data.Status[i].Reason = GetRandomSalt()
		data.Status[i].Memo = GetRandomSalt()
		data.Status[i].StartTime = uint64(time.Now().Unix() * 1000)
		data.Status[i].EndTime = uint64(time.Now().Unix() * 1000)

		i++
	}
	return data
}


func (d LightCheckPlan) Random(plantid int, t uint64) DataBase {
	var data LightCheckPlan
	//检修计划
	data.Time = t
	data.Status = make([]CheckPlanStatus, 4, 4)

	i := 0
	for _, v := range LightPlantCodeToId {
		data.Status[i].ID = v
		data.Status[i].IsTemporary = false
		data.Status[i].Capability = rand.Float64() * 100
		data.Status[i].CheckingDate = time.Now().Format("20060102150405")
		data.Status[i].Proposer = GetRandomSalt()
		data.Status[i].Approver = GetRandomSalt()
		data.Status[i].StartTime = uint64(time.Now().Unix() * 1000)
		data.Status[i].EndTime = uint64(time.Now().Unix() * 1000)

		i++
	}
	return data
}

func (d LightRestrictedPlan) Random(plantid int, t uint64) DataBase {
	var data LightRestrictedPlan
	//检修计划
	data.Time = t
	data.Status = make([]RestrictedPlanStatus, 4, 4)

	i := 0
	for _, v := range LightPlantCodeToId {
		data.Status[i].ID = v
		data.Status[i].BeforeRestriction = false
		data.Status[i].AfterRestriction = false
		data.Status[i].RestrictionPower = rand.Float64() * 100
		data.Status[i].RestrictionPoweVolum = rand.Float64() * 100
		data.Status[i].Reason = GetRandomSalt()
		data.Status[i].Memo = GetRandomSalt()
		data.Status[i].StartTime = uint64(time.Now().Unix() * 1000)
		data.Status[i].EndTime = uint64(time.Now().Unix() * 1000)

		i++
	}
	return data
}



func (d LightGenerationPlan) Random(plantid int, t uint64) DataBase {
	var data LightGenerationPlan
	//检修计划
	data.Time = t
	data.Status = make([]GenerationPlanStatus, 4, 4)

	i := 0
	for _, v := range LightPlantCodeToId {
		data.Status[i].ID = v
		data.Status[i].BeforeRestriction = false
		data.Status[i].AfterRestriction = false
		data.Status[i].RestrictionPower = rand.Float64() * 100
		data.Status[i].RestrictionPoweVolum = rand.Float64() * 100
		data.Status[i].Reason = GetRandomSalt()
		data.Status[i].Memo = GetRandomSalt()
		data.Status[i].StartTime = uint64(time.Now().Unix() * 1000)
		data.Status[i].EndTime = uint64(time.Now().Unix() * 1000)

		i++
	}
	return data
}




func GetRandomData(index int, siteid string, t string) []byte {
	var b []byte
	var tmp []byte
	var plantid int
	var ok bool
	var time uint64
	time, err := strconv.ParseUint(t, 10, 64)
	if err != nil {
		fmt.Println("time string error :", index, siteid, t, err.Error())
		return tmp
	}

	switch index {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
		plantid, ok = WindPlantCodeToId[siteid]
	case 9, 10, 11, 12, 13, 14, 15, 16, 17:
		plantid, ok = LightPlantCodeToId[siteid]
	}

	if !ok {
		fmt.Println("plantid error ", siteid)
		return tmp
	}

	fmt.Println("GetRandomData:", index, siteid, time, plantid)
	var database DataBase
	switch index {
	case 0:
		var data WindRealData
		database = data.Random(plantid, time)
	case 1, 10:
		var data ShortpredictionPowerdata
		database = data.Random(plantid, time)
	case 2, 11:
		var data UltraShortpredictionPowerdata
		database = data.Random(plantid, time)
	case 3:
		var data WindMeterTowerData
		database = data.Random(plantid, time)
	case 5:
		var data WindTurbineStatus
		database = data.Random(plantid, time)
	case 6:
		var data WindCheckPlan
		database = data.Random(plantid, time)
	case 7:
		var data WindRestrictedPlan
		database = data.Random(plantid, time)
	case 8:
		var data WindGenerationPlan
		database = data.Random(plantid, time)

	case 9:
		var data LightRealData
		database = data.Random(plantid, time)
	case 12:
		var data LightObservationData
		database = data.Random(plantid, time)
	case 14:
		var data InverterStatus
		database = data.Random(plantid, time)
	case 15:
		var data LightCheckPlan
		database = data.Random(plantid, time)
	case 16:
		var data LightRestrictedPlan
		database = data.Random(plantid, time)
	case 17:
		var data LightGenerationPlan
		database = data.Random(plantid, time)
	}
	b, err = PacketJson(database)
	if err != nil {
		fmt.Println("GetRandomData err:", err.Error())
		return tmp
	}
	return b
}

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// return len=8  salt
func GetRandomSalt() string {
	return GetRandomString(8)
}

//生成随机字符串
func GetRandomString(length int) string {
	var bytes []byte
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes = []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
