package pkg

import "github.com/joewilson27/go-vue-project/models"

func process(numsdata models.NumsRequest) models.NumsResponseData {
	var numsres models.NumsResponseData
	numsres.Add = numsdata.Num1 + numsdata.Num2
	numsres.Mul = numsdata.Num1 * numsdata.Num2
	numsres.Sub = numsdata.Num1 - numsdata.Num2
	numsres.Div = numsdata.Num1 / numsdata.Num2

	return numsres
}
