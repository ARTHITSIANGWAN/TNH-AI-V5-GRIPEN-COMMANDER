package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/genai" // SDK ตัวจี๊ด
)

// 🛡️ Dark-Relay Security: เกราะป้องกันการเข้าถึงจากภายนอก
func isAuthorized(r *http.Request) bool {
	secret := os.Getenv("A2A_SECRET_KEY")
	return r.Header.Get("X-ThitNuea-Auth") == secret
}

// 🧠 Intelligence Engine: ส่วนที่เชื่อมต่อกับ Gemini 2.0 Flash
func askGemini(ctx context.Context, prompt string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGoogleAI,
	})
	if err != nil {
		return "", err
	}

	// ใช้รุ่น Flash เพื่อความเร็วและประหยัด (Zero Garbage Focus)
	resp, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", genai.Text(prompt), nil)
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
	}
	return "No response from AI", nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" { port = "8081" }

	fmt.Printf("🏹 Gripen Engine: Gemini Intelligence Active on Port %s\n", port)

	// Endpoint หลักที่ F-16 จะส่งงานมาให้ (A2A Handshake)
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		if !isAuthorized(r) {
			http.Error(w, "🚫 Unauthorized", http.StatusUnauthorized)
			return
		}

		var task struct {
			Action  string `json:"action"`
			Content string `json:"content"`
		}
		
		// 📥 รับงานแบบประหยัด Memory
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid Payload", http.StatusBadRequest)
			return
		}

		// 🐍 Snake Nudge: ส่งสัญญาณตอบรับเบื้องต้นว่ากำลังประมวลผล
		fmt.Printf("🎯 Gripen Start Task: %s\n", task.Action)

		// 🧠 เรียก AI มาช่วยคิด (งาน Edit หรือ Intelligence)
		aiResult, err := askGemini(r.Context(), task.Content)
		if err != nil {
			fmt.Fprintf(w, `{"status": "error", "message": "%v"}`, err)
			return
		}

		// 📤 ส่งผลลัพธ์ระดับ 100% Clarity กลับไป
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "completed",
			"result": aiResult,
		})
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
