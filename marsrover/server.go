package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type roverRequest struct {
	Coordinates   Coordinates   `json:"coordinates"`
	Direction     string        `json:"direction"`
	MaxGridPlanet MaxGridPlanet `json:"maxGridPlanets"`
}
type roverCommandRequest struct {
	Instructions string `json:"instruction"`
}

var rover Rover

func main() {
	router := gin.Default()
	router.POST("/rover", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			sendErrorResponse(c, err)
		}
		var roverRequest roverRequest
		err = json.Unmarshal(body, &roverRequest)
		if err != nil {
			sendErrorResponse(c, err)
		}

		rover.Coordinates = roverRequest.Coordinates
		rover.Direction = getDirection(roverRequest.Direction)
		rover.MaxGridPlanet = roverRequest.MaxGridPlanet
		sendSuccessResponse(c, "Rover successfully initialised")
	})
	router.POST("/rover/instructions", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			sendErrorResponse(c, err)
		}
		var instructions roverCommandRequest
		err = json.Unmarshal(body, &instructions)
		if err != nil {
			sendErrorResponse(c, err)
		}
		position, errorMsg := rover.start(instructions.Instructions)
		if errorMsg != nil {
			sendErrorResponse(c, errorMsg)
		} else {
			sendSuccessResponse(c, position)
		}
	})

	v1 := router.Group("/rover")
	{
		v1.GET("/left", func(c *gin.Context) {
			if rover.IsEmpty() {
				sendSuccessResponse(c, "Initialize rover")
			}else{
				unitInstructionHandler(c, "L")
			}
		})
		v1.GET("/right", func(c *gin.Context) {
			if rover.IsEmpty() {
				sendSuccessResponse(c, "Initialize rover")
			}else{
				unitInstructionHandler(c, "R")
			}
		})

		v1.GET("/move", func(c *gin.Context) {
			if rover.IsEmpty() {
				sendSuccessResponse(c, "Initialize rover")
			}else{
				unitInstructionHandler(c, "M")
			}
		})
		v1.GET("/", func(c *gin.Context) {
			if rover.IsEmpty() {
				sendSuccessResponse(c, "Initialize rover")
			}else {
				sendSuccessResponse(c, rover.toString())
			}
		})
	}
	router.Run()
}

func unitInstructionHandler(c *gin.Context, instruction string) {
	position, err := rover.start(instruction)
	if err != nil {
		sendErrorResponse(c, err)
	} else {
		sendSuccessResponse(c, position)
	}
}
func sendErrorResponse(c *gin.Context, errorMsg error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"message": errorMsg.Error(),
	})
}
func sendSuccessResponse(c *gin.Context, result string) {
	c.JSON(200, gin.H{
		"roverPosition": result,
	})
}

func getDirection(s string) Direction {
	switch s {
	case "N":
		return Enum.N
	case "S":
		return Enum.S
	case "E":
		return Enum.E
	case "W":
		return Enum.W
	}
	return nil
}
func (s Rover) IsEmpty() bool {
	return s.Equals(Rover{})
}

