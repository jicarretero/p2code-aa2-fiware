package idm

import (
	"bytes"
	"crypto"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/jicarretero/p2code-aa2-fiware/config"
)

var serviceURL string
var useIdm bool
var scope string = "default"
var clientId string = "data-service"
var token string
var tokenTime time.Time

// OIDCConfiguration represents the OpenID Connect configuration
type OIDCConfiguration struct {
	TokenEndpoint string `json:"token_endpoint"`
}

// DIDDocument represents the DID document structure
type DIDDocument struct {
	ID string `json:"id"`
}

func SetConfig(cfg *config.Config) {
	useIdm = cfg.Brokerld.UseIDM

	if !useIdm {
		return
	}

	serviceURL = cfg.Brokerld.IdmServiceURL

	tmp_scope := cfg.Brokerld.IdmScope
	if tmp_scope != "" {
		scope = tmp_scope
	}

	tmp_client_id := cfg.Brokerld.IdmClientId
	if tmp_client_id != "" {
		clientId = tmp_client_id
	}

	helperScript := cfg.Brokerld.IdmWalletHelper
	if helperScript != "" {
		cmd, err := exec.Command("/bin/sh", helperScript).Output()
		if err != nil {
			log.Fatalf("Failed to execute IDM wallet helper script: %v", err)
		} else {
			fmt.Println(string(cmd))
		}
	}

	tokenTime = time.Now().Add(-1 * time.Hour)

}

func GetOIDCV4Token() string {
	if !useIdm {
		return ""
	}

	if time.Since(tokenTime) < time.Minute*5 && token != "" {
		return token
	}

	// Step 1: Get token endpoint from OIDC configuration
	tokenEndpoint, err := getTokenEndpoint(serviceURL)
	if err != nil {
		log.Fatalf("Failed to get token endpoint: %v", err)
	}

	// Step 2: Read holder DID from wallet identity
	holderDID, err := getHolderDID("/cert/did.json")
	if err != nil {
		log.Fatalf("Failed to get holder DID: %v", err)
	}

	// Step 3: Create verifiable presentation
	verifiableCredential := os.Getenv("VERIFIABLE_CREDENTIAL")
	// fmt.Println(verifiableCredential)
	vp, err := createVerifiablePresentation(holderDID, verifiableCredential)
	if err != nil {
		log.Fatalf("Failed to create verifiable presentation: %v", err)
	}

	// Step 4: Create and sign JWT
	jwt, err := createAndSignJWT(holderDID, vp, "/cert/private-key.pem")
	if err != nil {
		log.Fatalf("Failed to create and sign JWT: %v", err)
	}

	// Step 5: Encode JWT as VP token
	vpToken := base64URLEncode([]byte(jwt))

	// Step 6: Get access token from token endpoint
	accessToken, err := getAccessToken(tokenEndpoint, vpToken, scope)
	if err != nil {
		log.Fatalf("Failed to get access token: %v", err)
	}

	// fmt.Println(accessToken)
	token = accessToken

	return token
}

func getTokenEndpoint(serviceURL string) (string, error) {
	resp, err := http.Get(serviceURL + "/.well-known/openid-configuration")
	if err != nil {
		return "", fmt.Errorf("failed to fetch OIDC configuration: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var config OIDCConfiguration
	if err := json.Unmarshal(body, &config); err != nil {
		return "", fmt.Errorf("failed to parse OIDC configuration: %w", err)
	}

	if config.TokenEndpoint == "" {
		return "", fmt.Errorf("token endpoint not found in OIDC configuration")
	}

	return config.TokenEndpoint, nil
}

func getHolderDID(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read DID file: %w", err)
	}

	var didDoc DIDDocument
	if err := json.Unmarshal(file, &didDoc); err != nil {
		return "", fmt.Errorf("failed to parse DID document: %w", err)
	}

	if didDoc.ID == "" {
		return "", fmt.Errorf("DID not found in document")
	}

	return didDoc.ID, nil
}

func createVerifiablePresentation(holderDID, verifiableCredential string) (map[string]interface{}, error) {
	vp := map[string]interface{}{
		"@context": []string{"https://www.w3.org/2018/credentials/v1"},
		"type":     []string{"VerifiablePresentation"},
		"holder":   holderDID,
	}

	if verifiableCredential != "" {
		vp["verifiableCredential"] = []string{verifiableCredential}
	} else {
		vp["verifiableCredential"] = []string{}
	}

	return vp, nil
}

func createAndSignJWT(holderDID string, vp map[string]interface{}, privateKeyPath string) (string, error) {
	// Create JWT header
	header := map[string]interface{}{
		"alg": "ES256",
		"typ": "JWT",
		"kid": holderDID,
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", fmt.Errorf("failed to marshal header: %w", err)
	}
	headerEncoded := base64URLEncode(headerJSON)

	// Create JWT payload
	payload := map[string]interface{}{
		"iss": holderDID,
		"sub": holderDID,
		"vp":  vp,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}
	payloadEncoded := base64URLEncode(payloadJSON)

	// Sign the JWT
	signature, err := signData(headerEncoded+"."+payloadEncoded, privateKeyPath)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}

	return headerEncoded + "." + payloadEncoded + "." + signature, nil
}

func signData(data string, privateKeyPath string) (string, error) {
	// Read private key file
	keyFile, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return "", fmt.Errorf("failed to read private key file: %w", err)
	}

	// Parse PEM encoded private key
	block, _ := pem.Decode(keyFile)
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block containing private key")
	}

	// Parse ECDSA private key
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse ECDSA private key: %w", err)
	}

	// Hash the data
	hasher := sha256.New()
	hasher.Write([]byte(data))
	hashed := hasher.Sum(nil)

	// Sign the hashed data
	signature, err := privateKey.Sign(nil, hashed, crypto.SHA256)
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %w", err)
	}

	return base64URLEncode(signature), nil
}

func base64URLEncode(data []byte) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	encoded = strings.ReplaceAll(encoded, "+", "-")
	encoded = strings.ReplaceAll(encoded, "/", "_")
	encoded = strings.TrimRight(encoded, "=")
	return encoded
}

func getAccessToken(tokenEndpoint, vpToken, scope string) (string, error) {
	// Prepare form data
	formData := url.Values{}
	formData.Set("grant_type", "vp_token")
	formData.Set("client_id", clientId)
	formData.Set("vp_token", vpToken)
	formData.Set("scope", scope)

	// Create request
	req, err := http.NewRequest("POST", tokenEndpoint, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}

	// Parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "", fmt.Errorf("failed to parse token response: %w", err)
	}

	if tokenResponse.AccessToken == "" {
		return "", fmt.Errorf("access token not found in response")
	}

	return tokenResponse.AccessToken, nil
}
