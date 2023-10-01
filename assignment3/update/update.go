package update

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/adindazenn/assignment3-03/assignment3/model"
    "encoding/json"
	"net/http"
	"bytes"
)

func UpdateDataPeriodically() {
	for {
		// Generate nilai acak untuk "water" dan "wind" antara 1-100
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		// Buat data yang akan dikirim ke API
		dataToUpdate := model.Data{
			Water: water,
			Wind:  wind,
		}

		// Mengirim permintaan PUT ke API
		apiURL := "http://localhost:8080/api/update" 
		payload, _ := json.Marshal(dataToUpdate)
		req, err := http.NewRequest("PUT", apiURL, bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Error creating request:", err)
			time.Sleep(15 * time.Second) // Tunggu sebelum mencoba lagi
			continue
		}
	
		req.Header.Set("Content-Type", "application/json")
	
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			time.Sleep(15 * time.Second) // Tunggu sebelum mencoba lagi
			continue
		}
		defer resp.Body.Close()
	
		if resp.StatusCode != http.StatusOK {
			fmt.Println("API returned non-OK status:", resp.Status)
			time.Sleep(15 * time.Second) // Tunggu sebelum mencoba lagi
			continue
		}

		// Tentukan status berdasarkan nilai
		waterStatus := getStatus(water, 6, 8)
		windStatus := getStatus(wind, 7, 15)

		// Buat log dalam format JSON
		logData := model.Data{
			Water: water,
			Wind:  wind,
		}

		logJSON, _ := json.MarshalIndent(logData, "", "  ")
		fmt.Printf("%s\nstatus water: %s\nstatus wind: %s\n", string(logJSON), waterStatus, windStatus)

		// Tunggu 15 detik sebelum mengupdate data lagi
		time.Sleep(15 * time.Second)
	}
}

func getStatus(value, safeThreshold, dangerThreshold int) string {
	if value < safeThreshold {
		return "aman"
	} else if value >= safeThreshold && value <= dangerThreshold {
		return "siaga"
	} else {
		return "bahaya"
	}
}
