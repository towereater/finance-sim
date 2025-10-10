package service

import (
	"mainframe-lib/checking-account/model"
	qcom "mainframe-lib/common/queue"
	scom "mainframe-lib/common/service"
	"mainframe/checking-account/config"
)

func QueuePayment(queue config.Queue, abi string, payment model.Payment) error {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(queue.Timeout)
	defer cancel()

	// Queue document id
	err := qcom.QueueContent(ctx, queue.Queue, queue.Topics.Payments, abi, payment.Id)

	return err
}
