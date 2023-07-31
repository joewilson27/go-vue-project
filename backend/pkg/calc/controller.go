package pkg

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/go-vue-project/models"
)

func MainCalc(c *fiber.Ctx) error {
	// prepare model request to contain data
	var numsRequest models.NumsRequest
	// prepare response model
	var resData models.RespCalc

	if err := c.BodyParser(&numsRequest); err != nil {
		fmt.Println("Errrr", err.Error())
		resData.Data = nil

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
			"data":    resData,
		})
	}

	// process data
	returnData := process(numsRequest)
	resData.Data = returnData
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Ok",
		"data":    resData,
	})

}

// func calc(w http.ResponseWriter, request *http.Request) {
// 	// we are declaring a JSON decoder to decode the incoming JSON data from
// 	// therequest's body.
// 	decoder := json.NewDecoder(request.Body)

// 	// declared structures
// 	var numsData numbers
// 	var numsResData numsResponseData

// 	decoder.Decode(&numsData)

// 	numsResData = process(numsData)

// 	fmt.Println(numsResData)

// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(http.StatusOK)
// 	if err := json.NewEncoder(w).Encode(numsResData); err != nil {
// 		panic(err)
// 	}
// }

// func CreateFact(c *fiber.Ctx) error {
// 	var fact models.Fact
// 	if err := c.BodyParser(&fact); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	result := db.DB.Create(&fact)
// 	if result.Error != nil {
// 		return NewFactView(c)
// 	}
// 	//return c.Status(http.StatusOK).JSON(fact)

// 	// return to confirmation page view
// 	return ListFacts(c)
// }
