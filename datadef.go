package main

//struct字段命名首字母大写，json里的别名可以小写
type WindRealData struct {
	Time              uint64  `json:"time"`
	GridPower         uint8   `json:"gridPower"`
	OperatingCapacity uint8   `json:"operatingCapacity"`
	AccumulativePower float64 `json:"accumulativePower"`
	Restricted        bool    `json:"restricted"`
}

type WindTurbineStatusItem struct {
	ID            int     `json:"ID"`
	ActivePower   float64 `json:"activePower"`
	ReactivePower float64 `json:"reactivePower"`
	Windspeed     uint16  `json:"windspeed"`
	WindDirection uint16  `json:"windDirection"`
	Status        uint8   `json:"status"`
}
type WindTurbineStatus struct {
	Time   uint64                  `json:"time"`
	Status []WindTurbineStatusItem `json:"status"`
}

type WindMeterTowerData struct {
	Time      uint64     `json:"time"`
	WindTower [][][2]int `json:"windTower"`
}

type LightRealData WindRealData

type LightObservationData struct {
	Time                 uint64  `json:"time"`
	Temperature          uint16  `json:"temperature"`
	Humidity             uint16  `json:"humidity"`
	Pressure             float64 `json:"pressure"`
	Windspeed            uint16  `json:"windspeed"`
	WindDirection        uint16  `json:"windDirection"`
	TotalRadiance        float64 `json:"totalRadiance"`
	DirectRadiation      float64 `json:"directRadiation"`
	DiffuseRadiation     float64 `json:"diffuseRadiation"`
	Componenttemperature uint16  `json:"componenttemperature"`
}

type InverterStatusItem struct {
	ID            uint16  `json:"ID"`
	ActivePower   float64 `json:"activePower"`
	ReactivePower float64 `json:"reactivePower"`
	Status        uint8   `json:"status"`
}

type InverterStatus struct {
	Time   uint64               `json:"time"`
	Status []InverterStatusItem `json:"status"`
}

////////////////////////////////////////////////////////////////////////////////////////
//超短期
type ShortpredictionPowerdata struct {
	Time                 uint64    `json:"time"`
	ShortPowerPrediction []float64 `json:"shortPowerPrediction"`
}

//超短期
type UltraShortpredictionPowerdata struct {
	Time                      uint64    `json:"time"`
	UltraShortPowerPrediction []float64 `json:"Ultra-shortPowerPrediction"`
}

//5.7	检修计划，结构体有问题，文档写的不清晰
type CheckPlanStatus struct {
	ID           int     `json:"ID"`
	IsTemporary  bool    `json:"isTemporary"`
	Capability   float64 `json:"capability"`
	CheckingDate string  `json:"checkingDate"`
	Proposer     string  `json:"proposer"`
	Approver     string  `json:"approver"`
	StartTime    uint64  `json:"startTime"`
	EndTime      uint64  `json:"endTime"`
}
type WindCheckPlan struct {
	Time   uint64            `json:"time"`
	Status []CheckPlanStatus `json:"status"`
}

//5.8	限电记录，结构体有问题，文档写的不清晰
type RestrictedPlanStatus struct {
	ID                   int     `json:"ID"`
	BeforeRestriction    bool    `json:"beforeRestriction"`
	AfterRestriction     bool    `json:"afterRestriction"`
	RestrictionPower     float64 `json:"restrictionPower"`
	RestrictionPoweVolum float64 `json:"restrictionPoweVolum"`
	StartTime            uint64  `json:"startTime"`
	EndTime              uint64  `json:"endTime"`
	Reason               string  `json:"reason"`
	Memo                 string  `json:"memo"`
}
type WindRestrictedPlan struct {
	Time   uint64            `json:"time"`
	Status []RestrictedPlanStatus `json:"status"`
}

//5.9	发电计划，结构体有问题，文档写的不清晰
type GenerationPlanStatus struct {
	ID                   int     `json:"ID"`
	BeforeRestriction    bool    `json:"beforeRestriction"`
	AfterRestriction     bool    `json:"afterRestriction"`
	RestrictionPower     float64 `json:"restrictionPower"`
	RestrictionPoweVolum float64 `json:"restrictionPoweVolum"`
	StartTime            uint64  `json:"startTime"`
	EndTime              uint64  `json:"endTime"`
	Reason               string  `json:"reason"`
	Memo                 string  `json:"memo"`
}
type WindGenerationPlan struct {
	Time   uint64                 `json:"time"`
	Status []GenerationPlanStatus `json:"status"`
}

type LightCheckPlan WindCheckPlan
type LightRestrictedPlan WindRestrictedPlan
type LightGenerationPlan WindGenerationPlan
////////////////////////////////////////////////////////////////////////////////////////
//接口，类似于基类
type DataBase interface {
	//Packet() ([]byte, error)
	Random(plantid int, t uint64) DataBase
}
