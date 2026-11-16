package main

import (
	"fmt"
	"net/http"
)

func main() {
	// تعریف مسیر اصلی (/) و پاسخ به درخواست‌ها
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Go")
	})

	// اجرای سرور روی پورت 8080
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}