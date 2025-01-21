package chat

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type Message struct {
	User    string `json:"user"`
	Content string `json:"content"`
}

var (
	messages      []Message
	users         = make(map[string]chan Message)
	messagesMutex sync.Mutex
	usersMutex    sync.Mutex
)

func JoinChat(c *gin.Context) {
	username := c.Param("username")

	usersMutex.Lock()
	if _, exist := users[username]; !exist {
		users[username] = make(chan Message, 10)
	}
	usersMutex.Unlock()

	c.JSON(200, gin.H{"message": "joined", "user": username})
}

func SendMessage(c *gin.Context) {
	var newMsg Message
	if err := c.ShouldBindJSON(&newMsg); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	messagesMutex.Lock()
	messages = append(messages, newMsg)
	messagesMutex.Unlock()

	usersMutex.Lock()
	for _, userChan := range users {
		userChan <- newMsg
	}
	usersMutex.Unlock()

	c.JSON(200, gin.H{"message": "message sent successfully"})

}

func GetMessages(c *gin.Context) {
	username := c.Query("username")
	usersMutex.Lock()
	userChan, exists := users[username]
	usersMutex.Unlock()

	if !exists {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	select {
	case msg := <-userChan:
		c.JSON(200, msg)
	default:
		c.JSON(204, nil)
	}
}
