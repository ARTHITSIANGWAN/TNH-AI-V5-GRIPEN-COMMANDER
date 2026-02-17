package main

import (
	"fmt"
	"net/http"
	"os"
)

// gripen engine: eric full stack, gemini intelligence active
func main() {
	port := os.Getenv("PORT")
	if port == "" { port = "8081" }

	// กำหนด Key ลับ (ในใช้งานจริงควรเก็บใน Environment Variable)
	sharedSecret := os.Getenv("A2A_SECRET_KEY") 

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		// 🛡️ CHECK 1: ตรวจสอบสิทธิ์ (Authentication)
		clientSecret := r.Header.Get("X-ThitNuea-Auth")
		if clientSecret == "" || clientSecret != sharedSecret {
			http.Error(w, "🚫 Unauthorized: Unknown Agent Access", http.StatusUnauthorized)
			return
		}

		// 🛡️ CHECK 2: เฉพาะ POST เท่านั้น (Method Restriction)
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		fmt.Fprintf(w, "gripen: a2a handshake success. gemini 2.0 flash is thinking...")
		// ที่นี่จะเป็นจุดเชื่อมต่อกับ Gemini API หรือ AI Model ต่อไป
	})

	fmt.Printf("🚀 Gripen Intelligence Engine started on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
