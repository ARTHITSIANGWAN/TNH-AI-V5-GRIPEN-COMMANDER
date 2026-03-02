package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/genai" // 🏹 SDK ตัวจี๊ด Gemini 2.0
)

// --- [🛡️ IDENTITY: THITNUEA HUB - THE TRIPLE YUM] ---
// 🏹 บ้าน 1 (Gripen): Intelligence Engine (Gemini 2.0 Flash)
// ⛽ บ้าน 2 (Dark-Relay): Gas Station Scheduler (08:00, 12:00, 20:00)
// 🐍 บ้าน 3 (Nam-Ing): Supervisor & Snake Nudge Recall

func main() {
	// 🏁 1. มุดดิน: เริ่มระบบตั้งเวลาพ่นงาน (Background Scheduler)
	go runTripleYumScheduler()

	port := os.Getenv("PORT")
	if port == "" { port = "8081" }

	// 🛰️ 2. บนฟ้า: เปิดรับคำสั่ง A2A Handshake (Gripen API)
	http.HandleFunc("/process", handleA2AProcess)
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🛡️ F-16 DEFENDER V.2 | TRIPLE YUM ONLINE ✅\nStatus: Nam-Ing Active 🐍 | Engine: Gripen 2.0 🏹")
	})

	fmt.Printf("🚀 Gripen Fusion Active on Port %s | 💰 Mode: Zero-Garbage\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// --- [⏰ ส่วนที่ 1: Gas Station Scheduler (ยำเวลาดาร์กเรเลย์)] ---
func runTripleYumScheduler() {
	targetHours := []int{8, 12, 20} // เวลาพ่นงานตามเป้าหมาย
	loc := time.FixedZone("Asia/Bangkok", 7*60*60)

	for {
		now := time.Now().In(loc)
		nextRun := calculateNextRun(now, targetHours, loc)

		fmt.Printf("😴 [Gas Station]: พักเครื่อง.. รอบถัดไปคือ %s\n", nextRun.Format("15:04:05"))
		time.Sleep(time.Until(nextRun))

		// 🏁 น้ำอิงสั่งลุย! (Execute Mission)
		executeTripleYumMission()
	}
}

// --- [🧠 ส่วนที่ 2: Intelligence Engine (ยำสมองกริพเพน)] ---
func askGripen(ctx context.Context, prompt string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGoogleAI,
	})
	if err != nil { return "", err }

	// 🏹 ใช้ Gemini 2.0 Flash: เร็ว แรง ประหยัด (Zero Garbage Focus)
	resp, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", genai.Text(prompt), nil)
	if err != nil { return "", err }

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
	}
	return "No response from AI", nil
}

// --- [🐍 ส่วนที่ 3: น้ำอิง Supervisor (ยำระบบบันไดงู)] ---
func executeTripleYumMission() {
	ctx := context.Background()
	prompt := "Task: Generate Elite Tech Insight for SME. Style: Aggressive Thai, Professional. Policy: Zero-Garbage."

	// 🐍 Snake Nudge: บันไดงูถ้า Error ให้ Retry 3 ครั้ง
	for i := 0; i < 3; i++ {
		result, err := askGripen(ctx, prompt)
		if err != nil {
			log.Printf("⚠️ [Nam-Ing]: Error Detected! Snake Nudge Recall (Attempt %d): %v", i+1, err)
			time.Sleep(10 * time.Second)
			continue
		}
		
		// 🔊 Finisher: พ่นผลลัพธ์ออกสู่สายตาโลก
		fmt.Printf("\n--- 🛡️ TRIPLE YUM FINAL REPORT (Nam-Ing Approved) ---\n%s\n------------------------------------------------\n", result)
		break
	}
}

// --- [🛠️ Helper: จัดการการรันครั้งต่อไป] ---
func calculateNextRun(now time.Time, hours []int, loc *time.Location) time.Time {
	for _, h := range hours {
		t := time.Date(now.Year(), now.Month(), now.Day(), h, 0, 0, 0, loc)
		if t.After(now) { return t }
	}
	return time.Date(now.Year(), now.Month(), now.Day()+1, hours[0], 0, 0, 0, loc)
}

func handleA2AProcess(w http.ResponseWriter, r *http.Request) {
    // เช็คเกราะป้องกัน Auth
    secret := os.Getenv("A2A_SECRET_KEY")
    if r.Header.Get("X-ThitNuea-Auth") != secret {
        http.Error(w, "🚫 Unauthorized", http.StatusUnauthorized)
        return
    }
    // ... (รับ Payload และสั่ง Gripen ทำงานเหมือนเดิม)
}
