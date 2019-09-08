package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type roverRequest struct {
	XPosition int `json:"xPosition"`
	YPosition int `json:"yPosition"`
	Direction string `json:"direction"`
	MaxGridPlanet MaxGridPlanet `json:"maxGridPlanets"`
}
type roverCommandRequest struct {
	Instructions string `json:"instruction"`
}
var rover Rover

func main(){
	router :=gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
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
		    rover.MaxGridPlanet = roverRequest.MaxGridPlanet
			rover.XPosition =roverRequest.XPosition
		    rover.YPosition =roverRequest.YPosition
		    rover.Direction =roverRequest.Direction
		 sendSuccessResponse(c, "Rover successfully initialised")
	})
	router.POST("/rover/instructions", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			sendErrorResponse(c,err)
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
	router.GET("/rover/left", func(c *gin.Context) {
		unitInstructionHandler(c, "L")
	})

	router.GET("/rover/right", func(c *gin.Context) {
		unitInstructionHandler(c, "R")
	})

	router.GET("/rover/move", func(c *gin.Context) {
		unitInstructionHandler(c, "M")
	})

	router.GET("/", func(c *gin.Context) {
		sendSuccessResponse(c, rover.toString())
	})

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