package calculator

import "go-api/entity/message"

func GetValueMinus(messageData message.ReceiveMessage) int {
	return messageData.FirstNum - messageData.SecondNum
}
