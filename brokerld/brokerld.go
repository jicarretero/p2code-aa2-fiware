package brokerld

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jicarretero/p2code-aa2-fiware/config"
	"github.com/jicarretero/p2code-aa2-fiware/idm"
	"github.com/jicarretero/p2code-aa2-fiware/models"
)

var url string
var context string
var walletType string
var walletAddress string
var walletToken string
var tenant string
var rq = 0
var dumpCurl = true

func SetConfig(cfg *config.Config) {
	url = cfg.Brokerld.URL
	context = cfg.Brokerld.Context
	walletType = cfg.Brokerld.WalletType
	walletAddress = cfg.Brokerld.WalletAddress
	walletToken = cfg.Brokerld.WalletToken
	tenant = cfg.Brokerld.Tenant
}

func GetFullPath(basePath string) string {
	if basePath[len(basePath)-1] == '/' {
		return basePath + "ngsi-ld/v1/entities"
	}
	return basePath + "/ngsi-ld/v1/entities"
}

func DumpCurl(r *http.Request, bodyBytes []byte) {
	if !dumpCurl {
		return
	}

	rq = rq + 1

	s := fmt.Sprintf("\\\n\\\ncurl -X %s ${NGSILD_ADDRESS}%s \\\n", r.Method, r.URL.Path)
	for key, value := range r.Header {
		s = fmt.Sprintf("%s -H \"%s: %s\" \\\n", s, key, value[0])
	}
	s = fmt.Sprintf("%s-d '%s'", s, string(bodyBytes))

	log.Printf("\n%s\n", s)

	tmpFile, err := os.OpenFile(fmt.Sprintf("/tmp/here-%d.req", rq), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer tmpFile.Close()

	// Write the byte array to the file
	if _, err := tmpFile.WriteString(s); err != nil {
		return
	}
}

// Send sends the JSON-LD payload to the NGSI-LD broker.
// This function will only do the HTTP POST request.
// It does not handle any response or errors.
func Post(payload []byte) (int, error) {
	postUrl := GetFullPath(url)
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return -1, err
	}

	token := idm.GetOIDCV4Token()
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	req.Header.Add("Link", context)
	if walletType != "" {
		req.Header.Add("Wallet-Type", walletType)
		req.Header.Add("Wallet-Address", walletAddress)
		req.Header.Add("Wallet-Token", walletToken)
	}
	if tenant != "" {
		req.Header.Add("NGSILD-TENANT", tenant)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-encoding", "identity")

	DumpCurl(req, payload)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return -1, err
	}
	defer resp.Body.Close()

	fmt.Println("POST Response status:", resp.Status)
	return resp.StatusCode, nil
}

func Patch(payload []byte, id string) (int, error) {
	postUrl := GetFullPath(url) + "/" + id
	op := "PATCH"

	// If using Canis Major, I need to use POST /attrs for updating entities
	if walletType != "" {
		op = "POST"
		postUrl = postUrl + "/attrs"
	}

	req, err := http.NewRequest(op, postUrl, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return -1, err
	}

	token := idm.GetOIDCV4Token()
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	req.Header.Add("Link", context)
	if walletType != "" {
		req.Header.Add("Wallet-Type", walletType)
		req.Header.Add("Wallet-Address", walletAddress)
		req.Header.Add("Wallet-Token", walletToken)
	}
	if tenant != "" {
		req.Header.Add("NGSILD-TENANT", tenant)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-encoding", "identity")

	DumpCurl(req, payload)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return -1, err
	}
	defer resp.Body.Close()

	fmt.Println("PATCH Response status:", resp.Status)
	return resp.StatusCode, nil
}

func Get(id string) (int, error) {
	getUrl := GetFullPath(url) + "/" + id
	req, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return -1, err
	}

	token := idm.GetOIDCV4Token()
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	req.Header.Add("Link", context)
	if tenant != "" {
		req.Header.Add("NGSILD-TENANT", tenant)
	}
	req.Header.Add("Accept-encoding", "identity")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return -1, err
	}
	defer resp.Body.Close()

	fmt.Println("Get Response status:", resp.Status)
	return resp.StatusCode, nil
}

func MapDeviceProfile(topic string, data []byte) {
	deviceData, err := models.DeserializeData(data)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	data_array := models.NewDeviceProfileJSONLD(deviceData)

	for _, d := range data_array {
		md, _ := json.Marshal(d)
		fmt.Println(string(md))

		s, err := Get(d.GetId())
		if err != nil {
			fmt.Printf("Error sending data to Broker-LD: %s\n", err)
		} else {
			if s >= 400 {
				_, err = Post(md)
			} else {
				_, err = Patch(md, d.GetId())
			}
			if err != nil {
				fmt.Printf("Error sending data to Broker-LD: %s\n", err)
			}
		}
	}
}
