package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// 🧠 กึ๋นของ Gripen: การประมวลผลงานระดับ "Edit"
func processEditTask(task string) string {
	// ในอนาคตจุดนี้คือการส่งต่อไปยัง Gemini API
	// ตอนนี้ทำ Mockup การ "Edit" แบบไวๆ ให้เห็นภาพ
	return fmt.Sprintf("Gripen-Edit: [Optimized] %s at %v", task, time.Now().Unix())
}

func gripenHandler(w http.ResponseWriter, r *http.Request) {
	// 🛡️ SECURITY CHECK: ไส้ในต้องมีเกราะ
	secret := os.Getenv("A2A_SECRET_KEY")
	if r.Header.Get("X-ThitNuea-Auth") != secret {
		http.Error(w, "🚫 Breach Attempt Detected!", http.StatusUnauthorized)
		return
	}

	// 📥 RECEIVE & DECODE: ถอดรหัสคำสั่ง (เน้น Clean Code)
	var incomingTask AIDispatch
	if err := json.NewDecoder(r.Body).Decode(&incomingTask); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// ⚙️ THE EDIT ENGINE: เริ่มการประมวลผล "ตัวจี๊ด"
	fmt.Printf("🎯 Gripen Processing: %s from %s\n", incomingTask.Action, incomingTask.Sender)
	
	result := processEditTask(incomingTask.Action)

	// 📤 RESPONSE: ส่งผลลัพธ์กลับแบบรวดเร็ว
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "completed", "result": "%s"}`, result)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" { port = "8081" }

	fmt.Printf("🏹 Gripen A2A (Enterprise/Edit Mode) Engine Started on %s\n", port)
	
	http.HandleFunc("/process", gripenHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "💓 Gripen Engine: Stable")
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
