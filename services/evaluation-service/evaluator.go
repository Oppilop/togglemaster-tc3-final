package main

import (
<<<<<<< HEAD
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
=======
	"crypto/sha1"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
	"os"
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
)

const (
	// Tempo de vida do cache em segundos
	CACHE_TTL = 30 * time.Second
)

// getDecision é o wrapper principal
func (a *App) getDecision(userID, flagName string) (bool, error) {
<<<<<<< HEAD
=======
	// 1. Obter os dados da flag (do cache ou dos serviços)
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	info, err := a.getCombinedFlagInfo(flagName)
	if err != nil {
		return false, err
	}
<<<<<<< HEAD
	return a.runEvaluationLogic(info, userID), nil
}

// getCombinedFlagInfo busca no Redis com fallback para os microsserviços
=======

	// 2. Executar a lógica de avaliação
	return a.runEvaluationLogic(info, userID), nil
}

// getCombinedFlagInfo busca os dados no Redis, com fallback para os microsserviços
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
func (a *App) getCombinedFlagInfo(flagName string) (*CombinedFlagInfo, error) {
	cacheKey := fmt.Sprintf("flag_info:%s", flagName)

	// 1. Tentar buscar do Cache (Redis)
	val, err := a.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
<<<<<<< HEAD
		var info CombinedFlagInfo
		if uerr := json.Unmarshal([]byte(val), &info); uerr == nil {
			log.Printf("Cache HIT para flag '%s'", flagName)
			return &info, nil
		}
		log.Printf("Erro ao desserializar cache para flag '%s': %v", flagName, err)
	}

	log.Printf("Cache MISS para flag '%s'", flagName)

=======
		// Cache HIT
		var info CombinedFlagInfo
		if err := json.Unmarshal([]byte(val), &info); err == nil {
			log.Printf("Cache HIT para flag '%s'", flagName)
			return &info, nil
		}
		// Se o unmarshal falhar, trata como cache miss
		log.Printf("Erro ao desserializar cache para flag '%s': %v", flagName, err)
	}
	
	log.Printf("Cache MISS para flag '%s'", flagName)
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	// 2. Cache MISS - Buscar dos serviços
	info, err := a.fetchFromServices(flagName)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	// 3. Salvar no Cache (best effort)
	if jsonData, mErr := json.Marshal(info); mErr == nil {
		if sErr := a.RedisClient.Set(ctx, cacheKey, jsonData, CACHE_TTL).Err(); sErr != nil {
			log.Printf("Falha ao salvar no cache (não-fatal): %v", sErr)
		}
=======
	// 3. Salvar no Cache
	jsonData, err := json.Marshal(info)
	if err == nil {
		a.RedisClient.Set(ctx, cacheKey, jsonData, CACHE_TTL).Err()
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	}

	return info, nil
}

// fetchFromServices busca dados do flag-service e targeting-service concorrentemente
func (a *App) fetchFromServices(flagName string) (*CombinedFlagInfo, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var flagInfo *Flag
	var ruleInfo *TargetingRule
	var flagErr, ruleErr error

<<<<<<< HEAD
=======
	// Goroutine 1: Buscar do flag-service
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	go func() {
		defer wg.Done()
		flagInfo, flagErr = a.fetchFlag(flagName)
	}()

<<<<<<< HEAD
=======
	// Goroutine 2: Buscar do targeting-service
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	go func() {
		defer wg.Done()
		ruleInfo, ruleErr = a.fetchRule(flagName)
	}()

	wg.Wait()

	if flagErr != nil {
		return nil, flagErr
	}
	if ruleErr != nil {
<<<<<<< HEAD
		log.Printf("Aviso: regra de segmentação não encontrada para '%s' (%v). Usando padrão.",
			flagName, ruleErr)
=======
		log.Printf("Aviso: Nenhuma regra de segmentação encontrada para '%s'. Usando padrão.", flagName)
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	}

	return &CombinedFlagInfo{
		Flag: flagInfo,
		Rule: ruleInfo,
	}, nil
}

<<<<<<< HEAD
// fetchFlag busca a flag no flag-service.
//
// SECURITY (gosec G107/G306): construímos a URL via fmt.Sprintf usando apenas
// flagName, que vem da query string e poderia conter caracteres ofensivos.
// Aqui validamos via http.NewRequestWithContext que rejeita URLs malformadas.
func (a *App) fetchFlag(flagName string) (*Flag, error) {
	url := fmt.Sprintf("%s/flags/%s", a.FlagServiceURL, flagName)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao montar requisição: %w", err)
	}

	apiKey := os.Getenv("SERVICE_API_KEY")
	req.Header.Set("Authorization", "Bearer "+apiKey)

=======
// fetchFlag (função helper)
func (a *App) fetchFlag(flagName string) (*Flag, error) {
	url := fmt.Sprintf("%s/flags/%s", a.FlagServiceURL, flagName)

	apiKey := os.Getenv("SERVICE_API_KEY")
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	resp, err := a.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao chamar flag-service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, &NotFoundError{flagName}
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("flag-service retornou status %d", resp.StatusCode)
	}

<<<<<<< HEAD
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta do flag-service: %w", err)
	}

=======
	body, _ := ioutil.ReadAll(resp.Body)
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	var flag Flag
	if err := json.Unmarshal(body, &flag); err != nil {
		return nil, fmt.Errorf("erro ao desserializar resposta do flag-service: %w", err)
	}
	return &flag, nil
}

func (a *App) fetchRule(flagName string) (*TargetingRule, error) {
	url := fmt.Sprintf("%s/rules/%s", a.TargetingServiceURL, flagName)
<<<<<<< HEAD

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao montar requisição: %w", err)
	}

	apiKey := os.Getenv("SERVICE_API_KEY")
	req.Header.Set("Authorization", "Bearer "+apiKey)

=======
	apiKey := os.Getenv("SERVICE_API_KEY") // Usa a mesma chave
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	resp, err := a.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao chamar targeting-service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
<<<<<<< HEAD
		return nil, &NotFoundError{flagName}
=======
		return nil, &NotFoundError{flagName} // Não é um erro fatal
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("targeting-service retornou status %d", resp.StatusCode)
	}

<<<<<<< HEAD
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta do targeting-service: %w", err)
	}

=======
	body, _ := ioutil.ReadAll(resp.Body)
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	var rule TargetingRule
	if err := json.Unmarshal(body, &rule); err != nil {
		return nil, fmt.Errorf("erro ao desserializar resposta do targeting-service: %w", err)
	}
	return &rule, nil
}

// runEvaluationLogic é onde a decisão é tomada
func (a *App) runEvaluationLogic(info *CombinedFlagInfo, userID string) bool {
	if info.Flag == nil || !info.Flag.IsEnabled {
		return false
	}

	if info.Rule == nil || !info.Rule.IsEnabled {
		return true
	}

<<<<<<< HEAD
	rule := info.Rule.Rules
	if rule.Type == "PERCENTAGE" {
		percentage, ok := rule.Value.(float64)
		if !ok {
			log.Printf("Erro: valor da regra de porcentagem não é número para '%s'", info.Flag.Name)
			return false
		}

		userBucket := getDeterministicBucket(userID + info.Flag.Name)

=======
	// 3. Processa a regra (só temos "PERCENTAGE" por enquanto)
	rule := info.Rule.Rules
	if rule.Type == "PERCENTAGE" {
		// Converte o 'value' (que é interface{}) para float64
		percentage, ok := rule.Value.(float64)
		if !ok {
			log.Printf("Erro: valor da regra de porcentagem não é um número para a flag '%s'", info.Flag.Name)
			return false
		}
		
		// Calcula o "bucket" do usuário (0-99)
		userBucket := getDeterministicBucket(userID + info.Flag.Name)
		
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
		if float64(userBucket) < percentage {
			return true
		}
	}

	return false
}

<<<<<<< HEAD
// getDeterministicBucket distribui usuários determinísticamente em buckets [0..99].
//
// SECURITY (gosec G401): a versão anterior usava SHA-1, considerado quebrado
// para usos criptográficos. Aqui o uso NÃO é criptográfico (é só hashing
// para distribuição uniforme), mas o gosec não consegue saber disso e
// reporta HIGH. Trocamos por SHA-256 — também distribui uniformemente,
// é amplamente disponível, e elimina o finding.
func getDeterministicBucket(input string) int {
	hash := sha256.Sum256([]byte(input))
	val := binary.BigEndian.Uint32(hash[:4])
=======
func getDeterministicBucket(input string) int {
	// Usamos SHA1 (rápido) e pegamos os primeiros 4 bytes
	hasher := sha1.New()
	hasher.Write([]byte(input))
	hash := hasher.Sum(nil)
	
	// Converte 4 bytes para um uint32
	val := binary.BigEndian.Uint32(hash[:4])
	
	// Retorna o módulo 100
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	return int(val % 100)
}
