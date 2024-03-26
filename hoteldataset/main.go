package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

const urlTemplate = "https://cdnqafox.azureedge.net/gz-master/00%s.csv?sp=rl&st=2024-03-25T09:50:27Z&se=2025-07-25T16:50:27Z&spr=https&sv=2022-11-02&sr=c&sig=YmlEZBSb4WPtsl9RAi%%2Fo3RLyPoFQaJ5IvRUFrSXw1l4%%3D"

func download(idx int, wg *sync.WaitGroup) {
	var url string
	defer wg.Done()
	if idx < 10 {
		url = fmt.Sprintf(urlTemplate, fmt.Sprintf("0%d", idx))
	} else {
		url = fmt.Sprintf(urlTemplate, strconv.Itoa(idx))
	}
	log.Printf("Downloading %s", url)
	fileName := fmt.Sprintf("dataset_hotel_%d.csv", idx)
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	out, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		panic(err)
	}
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go download(i, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Downloaded all files in %s", elapsed)

}
