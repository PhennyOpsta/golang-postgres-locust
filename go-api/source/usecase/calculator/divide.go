package calculator

import "go-api/entity/message"

func GetValueDivide(messageData message.ReceiveMessage) int {
	return messageData.FirstNum / messageData.SecondNum
}
