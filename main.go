package main

import (
	"bytes"
	"fmt"
	"strings"
)

// ReadSatchaFromImage: สกัดคำสั่งที่ซ่อนอยู่ท้ายไฟล์รูปภาพ (Zero-Garbage)
func ReadSatchaFromImage(imgData []byte) string {
	// 1. ค้นหา Marker 'CMD:' ที่ฝังไว้ท้ายไฟล์
	marker := []byte("CMD:")
	idx := bytes.Index(imgData, marker)
	
	if idx == -1 {
		return "❌ No Command Found"
	}

	// 2. สกัดข้อมูลหลัง Marker ออกมา (Payload JSON)
	rawCmd := string(imgData[idx+len(marker):])
	return strings.TrimSpace(rawCmd)
}

func main() {
	// จำลองข้อมูลภาพที่ส่งมาจากปฏิทินหรือ R2
	mockImage := append([]byte{0x89, 0x50, 0x4E, 0x47}, []byte("\nCMD:{\"job\":\"IGNITE_V83\",\"id\":\"9333074\"}")...)
	
	satcha := ReadSatchaFromImage(mockImage)
	fmt.Printf("🎯 สัจจะที่ขุนพลพลายทองพบ: %s\n", satcha)
}
