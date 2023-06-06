package models

//Pump holds pump repair form data
type Pump struct {
	EquipmentID         string
	Person              string
	CheckValveCleaned   string
	SealsReplaced       string
	PumpHeadPiston      string
	CheckValvesReplaced string
	TimeRequired        string
	MeasuredEfficiency  string
	MaintanenceDate     string
}
