package main

import (
	"fmt"
	"net/http"
	"os"
)

// SystemInfo represents the system's details and assigned person's information.
type SystemInfo struct {
	SystemName string
	Color      string
	Hostname   string
	PersonName string
	PersonID   string
}

// getSystemInfo handles requests and serves the system information.
func getSystemInfo(w http.ResponseWriter, r *http.Request) {
	systemName := os.Getenv("SYSTEM_NAME")
	color := os.Getenv("BACKGROUND_COLOR")
	hostname, _ := os.Hostname()
	personName := os.Getenv("PERSON_NAME")
	personID := os.Getenv("PERSON_ID")

	info := SystemInfo{
		SystemName: systemName,
		Color:      color,
		Hostname:   hostname,
		PersonName: personName,
		PersonID:   personID,
	}

	html := fmt.Sprintf(`
    <html>
    <head>
        <style>
            body { 
                background-color: %s;
                font-family: Arial, sans-serif; 
                text-align: center; 
                color: white; 
                padding: 50px;
            }
            .info {
                font-size: 24px;
                margin-bottom: 20px;
            }
            .hostname, .person {
                font-size: 18px;
                margin-bottom: 10px;
                color: #ddd;
            }
        </style>
    </head>
    <body>
        <div class="info">%s</div>
        <div class="hostname">Hostname: %s</div>
        <div class="person">Assigned to: %s (ID: %s)</div>
    </body>
    </html>
    `, info.Color, info.SystemName, info.Hostname, info.PersonName, info.PersonID)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", getSystemInfo)
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
