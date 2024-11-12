package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"vid-con/entities"
)

type CallController struct {
	rooms map[string]*entities.Room
	mu    sync.Mutex
}

func InitVideo() *CallController {
	return &CallController{}
}

func (cc *CallController) Show(c *gin.Context) {
	roomId := c.Param("room_id")

	if roomId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id required"})
		return
	}
	room, exists := cc.rooms[roomId]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room you re trying to join does not exist"})
		return
	}

	attendantId := c.Query("attendantID")
	if attendantId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot join without attendant id"})
		return
	}
	attendant, exist := room.Attendants[attendantId]
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you re not authorized to join this room"})
	}
	fmt.Println("this is attendant id", attendant.ID)
	c.File("view/call/show.html")
}

func (cc *CallController) ConfirmJoin(c *gin.Context) {
	c.File("views/call/confirm_join.html")
}

func (cc *CallController) Join(c *gin.Context) {
	roomId := c.PostForm("id")
	if roomId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room id required"})
		return
	}

	attendantName := c.PostForm("attendantName")
	if attendantName == "" {
		c.JSON(http.StatusBadGateway, gin.H{"error": "attendant name cannot be blank"})
		return
	}
	attendant := entities.NewAttendant(attendantName)

	room, exists := cc.rooms[roomId]
	if !exists {
		fmt.Println("room must exist")
		return
	}

	room.AddAttendant(attendant)

	redirectURL := fmt.Sprintf("/call/%s?attendantId=%s", roomId, attendant.ID)
	c.Redirect(http.StatusFound, redirectURL)
}

func (cc *CallController) ConfirmStart(c *gin.Context) {
	c.File("views/call/confirm_start.html")
}

func (cc *CallController) Start(c *gin.Context) {
	attendantName := c.PostForm("attendantName")
	if attendantName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "attendant name cannot be blank"})
		return
	}
	room := entities.NewRoom()
	cc.rooms[room.ID] = room
	attendant := entities.NewAttendant(attendantName)
	room.AddAttendant(attendant)
	redirectURL := fmt.Sprintf("/call/%s?attendantId=%s", room.ID, attendant.ID)
	c.Redirect(http.StatusFound, redirectURL)
}

func (cc *CallController) Leave(c *gin.Context) {
	roomID := c.Param("room_id")
	attendantID := c.Param("attendant_id")

	if roomID == "" || attendantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomID or attendantID cannot be blank"})
		return
	}

	room, exist := cc.rooms[roomID]
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room does not exist"})
		return
	}

	attendant, exist := room.Attendants[attendantID]
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "attendant does not exist"})
	}
	room.RemoveAttendant(attendant)

	if len(room.Attendants) == 0 {
		delete(cc.rooms, roomID)
		fmt.Println("room deleted as no attendants remains")
	}
	c.Redirect(http.StatusOK, "/")
}
