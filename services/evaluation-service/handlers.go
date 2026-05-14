package main

import (
	"encoding/json"
<<<<<<< HEAD
	"errors"
=======
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	"log"
	"net/http"
)

type EvaluationResponse struct {
	FlagName string `json:"flag_name"`
	UserID   string `json:"user_id"`
	Result   bool   `json:"result"`
}

<<<<<<< HEAD
// writeJSON centraliza serialização + tratamento do erro do Encode (gosec G104).
func writeJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("erro ao escrever resposta JSON: %v", err)
	}
}

func (a *App) healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (a *App) evaluationHandler(w http.ResponseWriter, r *http.Request) {
=======
func (a *App) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (a *App) evaluationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. Parsear os query parameters
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	userID := r.URL.Query().Get("user_id")
	flagName := r.URL.Query().Get("flag_name")

	if userID == "" || flagName == "" {
<<<<<<< HEAD
		writeJSON(w, http.StatusBadRequest,
			map[string]string{"error": "user_id e flag_name são obrigatórios"})
		return
	}

	// Obter a decisão (lógica de cache/serviço está em evaluator.go)
	result, err := a.getDecision(userID, flagName)
	if err != nil {
		// Se for "não encontrado", retornamos 'false' (fail-closed)
		var nfe *NotFoundError
		if errors.As(err, &nfe) {
			result = false
		} else {
			log.Printf("Erro ao avaliar flag '%s': %v", flagName, err)
			writeJSON(w, http.StatusBadGateway,
				map[string]string{"error": "Erro interno ao avaliar a flag"})
=======
		http.Error(w, `{"error": "user_id e flag_name são obrigatórios"}`, http.StatusBadRequest)
		return
	}

	// 2. Obter a decisão (lógica de cache/serviço está em evaluator.go)
	result, err := a.getDecision(userID, flagName)
	if err != nil {
		// Se o erro for "não encontrado", retornamos 'false' (comportamento seguro)
		if _, ok := err.(*NotFoundError); ok {
			result = false
		} else {
			// Outros erros (serviços offline, etc)
			log.Printf("Erro ao avaliar flag '%s': %v", flagName, err)
			http.Error(w, `{"error": "Erro interno ao avaliar a flag"}`, http.StatusBadGateway)
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
			return
		}
	}

<<<<<<< HEAD
	// Envia evento para SQS assincronamente (não bloqueia a resposta).
	go a.sendEvaluationEvent(userID, flagName, result)

	writeJSON(w, http.StatusOK, EvaluationResponse{
=======
	// 3. Enviar evento para SQS (assincronamente)
	// Isso não bloqueia a resposta para o cliente.
	go a.sendEvaluationEvent(userID, flagName, result)

	// 4. Retornar a resposta
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(EvaluationResponse{
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
		FlagName: flagName,
		UserID:   userID,
		Result:   result,
	})
<<<<<<< HEAD
}
=======
}
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
