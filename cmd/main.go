package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"resizer/internal/http/request"
	service2 "resizer/internal/service/fetcher"
	service1 "resizer/internal/service/resizer"
)

func main() {
	http.HandleFunc("/resize", resizeHandler)
	port := 3000
	fmt.Printf("Server is running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func resizeHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody request.RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Fetch the image from the provided URL
	fmt.Println("fetching image from the source ...")
	resp, err := service2.Fetcher(reqBody)
	if err != nil {
		http.Error(w, "Failed to fetch image", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to fetch image", http.StatusInternalServerError)
		return
	}

	// Read the image data as a byte slice
	fmt.Println("Read the image as a slice of bytes ...")
	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read image data", http.StatusInternalServerError)
		return
	}

	// Resize the image
	fmt.Println("Resizing the image ...")
	img, err := service1.ResizeImage(imageData, reqBody.Width, reqBody.Height)
	if err != nil {
		http.Error(w, "Failed to resize image", http.StatusInternalServerError)
		return
	}

	// Write the resized image to response
	w.Header().Set("Content-Type", "image/jpeg")
	if _, err := w.Write(img); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return

	}
}
