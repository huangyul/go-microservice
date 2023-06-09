package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])
		jData, _ := json.Marshal(map[string]int{
			"Data": a + b,
		})
		_, _ = w.Write(jData)
	})

	http.ListenAndServe(":3333", nil)
}

Get-ChildItem -Path . -Recurse | foreach {
$newname = $_.Name -replace "【加微信z2018005801赠送精品IT课程】", ""
Rename-Item -Path $_.FullName -NewName $newname
}