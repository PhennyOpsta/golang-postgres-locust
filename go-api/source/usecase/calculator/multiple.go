package calculator

import "go-api/entity/message"

func GetValueMultiple(messageData message.ReceiveMessage) int {
	return messageData.FirstNum * messageData.SecondNum
}
