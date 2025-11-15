package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jexlor/cs2scraper/scraper"
	"github.com/jexlor/cs2scraper/scraper/config"
)

func main() {

	// here we run our scraper
	skins, err := scraper.ScrapSkins()
	if err != nil {
		log.Fatalf("Error during scraping: %v", err)
	}
	jsonData, err := json.MarshalIndent(skins, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// if you want data to be printed to console
	if config.ConsoleLog {
		fmt.Println(string(jsonData))
	}

	// generate filename with current date
	timestamp := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("output_%s.json", timestamp)

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	fmt.Printf("\n[+] Scraped data written to %s\n", filename)
}
