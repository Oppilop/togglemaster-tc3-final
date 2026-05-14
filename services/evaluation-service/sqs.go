package main

import (
<<<<<<< HEAD
	"context"
=======
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	"encoding/json"
	"log"
	"time"

<<<<<<< HEAD
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
=======
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
)

// Evento que será enviado para a fila
type EvaluationEvent struct {
	UserID    string    `json:"user_id"`
	FlagName  string    `json:"flag_name"`
	Result    bool      `json:"result"`
	Timestamp time.Time `json:"timestamp"`
}

<<<<<<< HEAD
// sendEvaluationEvent envia um evento para a fila SQS (aws-sdk-go-v2).
func (a *App) sendEvaluationEvent(userID, flagName string, result bool) {
	if a.SqsClient == nil || a.SqsQueueURL == "" {
		log.Printf("[SQS_DISABLED] Evento: User '%s', Flag '%s', Result '%t'",
			userID, flagName, result)
=======
// sendEvaluationEvent envia um evento para a fila SQS
func (a *App) sendEvaluationEvent(userID, flagName string, result bool) {
	// Se a URL da fila não foi configurada, apenas loga localmente e sai.
	if a.SqsSvc == nil || a.SqsQueueURL == "" {
		log.Printf("[SQS_DISABLED] Evento: User '%s', Flag '%s', Result '%t'", userID, flagName, result)
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
		return
	}

	event := EvaluationEvent{
		UserID:    userID,
		FlagName:  flagName,
		Result:    result,
		Timestamp: time.Now().UTC(),
	}

	body, err := json.Marshal(event)
	if err != nil {
		log.Printf("Erro ao serializar evento SQS: %v", err)
		return
	}

<<<<<<< HEAD
	// Timeout no contexto para não bloquear a goroutine indefinidamente.
	sendCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = a.SqsClient.SendMessage(sendCtx, &sqs.SendMessageInput{
=======
	// Envia a mensagem
	_, err = a.SqsSvc.SendMessage(&sqs.SendMessageInput{
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
		MessageBody: aws.String(string(body)),
		QueueUrl:    aws.String(a.SqsQueueURL),
	})

	if err != nil {
		log.Printf("Erro ao enviar mensagem para SQS: %v", err)
	} else {
		log.Printf("Evento de avaliação enviado para SQS (Flag: %s)", flagName)
	}
<<<<<<< HEAD
}
=======
}
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
