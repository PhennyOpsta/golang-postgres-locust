package calculator

import "go-api/entity/message"

func GetValuePlus(messageData message.ReceiveMessage) int {
	return messageData.FirstNum + messageData.SecondNum
}
