package models

type NumsRequest struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

type NumsResponseData struct {
	Add float64 `json:"add"`
	Mul float64 `json:"mul"`
	Sub float64 `json:"sub"`
	Div float64 `json:"div"`
}

type RespCalc struct {
	Data interface{} `json:"data"`
}
