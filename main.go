package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/jexlor/cs2parser/config"
	"github.com/jexlor/cs2parser/custom"
)

func main() {

	// fancy ascii title ;)
	fmt.Print(`
           ___                                 
          |__ \                                
   ___ ___   ) |_ __   __ _ _ __ ___  ___ _ __ 
  / __/ __| / /| '_ \ / _` + "`" + ` | '__/ __|/ _ \ '__|
 | (__\__ \/ /_| |_) | (_| | |  \__ \  __/ |   
  \___|___/____| .__/ \__,_|_|  |___/\___|_|   
               | |                             
               |_|                             
`)

	//some extra flags to make sure script isn't detected
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("exclude-switches", "enable-automation"),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("start-maximized", false),
		chromedp.Flag("window-size", "800,600"),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	fmt.Println("\nCreating allocator context and applying flags...")
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	fmt.Println("Created allocator")

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	list := []string{
		"kilowatt-case", "revolution-case", "recoil-case", "dreams-nightmares-case", "sealed-genesis-terminal",
		"snakebite-case", "fracture-case", "prisma-2-case", "cs20-case", "prisma-case", "danger-zone-case",
		"horizon-case", "clutch-case", "spectrum-2-case", "operation-hydra-case", "spectrum-case", "glove-case",
	}

	var allSkins []config.Skin
	total := len(list)

	for i, item := range list {
		url := "https://www.csgodatabase.com/cases/" + item + "/"
		maxWidth := 60 // maximum width for the url display. adjust as needed.
		display := url
		if len(url) > maxWidth {
			display = url[:maxWidth-3] + "..."
		}
		fmt.Printf("\rProgress: (%d/%d) | Scraping: %-60s", i+1, total, display)
		var pageTitle string
		var rawData []map[string]string

		// lets get string containing our js code to pass for evaluation
		jsCode, err := os.ReadFile("config/script.js")
		if err != nil {
			log.Fatalf("Error reading JS file: %v", err)
		}

		err = chromedp.Run(ctx,
			chromedp.Navigate(url),
			// trying to evade webdriver detection
			chromedp.Evaluate(`
				Object.defineProperty(navigator, 'webdriver', {get: () => undefined});
				window.chrome = {runtime: {}};
				Object.defineProperty(navigator, 'plugins', {get: () => [1,2,3,4,5]});`, nil),

			// that's to avoid triggering cloudflare/site protections. 1 sec is enough but the longer the better.
			// you can even skip that line and comment it, but I don't recommend doing so.
			// chromedp.Sleep(1*time.Second),
			chromedp.Title(&pageTitle),
			// that js script goes straight to console of page and retrieves whatever selectors you write (its configurable)
			chromedp.Evaluate(string(jsCode), &rawData),
		)

		if err != nil {
			fmt.Printf("\nchromedp run failed: %v\n", err)
			continue
		}

		for _, data := range rawData {
			allSkins = append(allSkins, config.Skin{

				// something like data["example"] here is a raw response from script.
				//  you can adjust fields by extra functions like yourFunction(data["example"]).
				Name:       data["name"],
				Weapon:     custom.SpecialMark(data["weapon"]),
				Rarity:     data["rarity"],
				Collection: data["collection"],
				Price:      custom.ParsePrice(data["price"], data["stattrakPrice"]),
				URL:        data["url"],
			})
		}

		jsonData, err := json.MarshalIndent(allSkins, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling JSON: %v", err)
		}
		// if you want data to be printed on console
		// fmt.Println(string(jsonData))

		err = os.WriteFile("output.txt", jsonData, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}

	fmt.Println("\nScraped data written to output.txt")
}
