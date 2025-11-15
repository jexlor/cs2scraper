package scraper

import (
	"context"
	_ "embed"
	"fmt"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/jexlor/cs2scraper/custom"
	"github.com/jexlor/cs2scraper/internal"
	"github.com/jexlor/cs2scraper/scraper/config"
)

//go:embed config/scripts/script.js
var ScriptJS string

//go:embed config/scripts/config.js
var ConfigJS string

// Run func executes the scraping process and returns a slice of Skin structs with error if any.
func ScrapSkins() ([]config.Skin, error) {
	// fancy ascii title ;)
	fmt.Print(`
   ____  ____     ____    ____      ____    ____        _       ____   U _____ u   ____     
U /"___|/ __"| u |___"\  / __"| uU /"___|U |  _"\ u U  /"\  u U|  _"\ u\| ___"|/U |  _"\ u  
\| | u <\___ \/  U __) |<\___ \/ \| | u   \| |_) |/  \/ _ \/  \| |_) |/ |  _|"   \| |_) |/  
 | |/__ u___) |  \/ __/ \u___) |  | |/__   |  _ <    / ___ \   |  __/   | |___    |  _ <    
  \____||____/>> |_____|u|____/>>  \____|  |_| \_\  /_/   \_\  |_|      |_____|   |_| \_\   
 _// \\  )(  (__)<<  //   )(  (__)_// \\   //   \\_  \\    >>  ||>>_    <<   >>   //   \\_  
(__)(__)(__)    (__)(__) (__)    (__)(__) (__)  (__)(__)  (__)(__)__)  (__) (__) (__)  (__) 
		`)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), config.Opts...)
	defer cancel()

	fmt.Println("\nCreating allocator context and applying opts...")
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	fmt.Println("Created allocator")

	ctx, cancel = context.WithTimeout(ctx, config.DeadLine)
	defer cancel()

	//essentials
	var allSkins []config.Skin

	total := len(config.List)

	for i, item := range config.List {
		url := config.Target + "/cases/" + item + "/"
		display := url
		if len(url) > config.UrlLengthLimit {
			display = url[:config.UrlLengthLimit-3] + "..."
		}
		fmt.Printf("\rProgress: (%d/%d) | Scraping: %-60s", i+1, total, display)
		var pageTitle string
		var rawData []map[string]string

		// lets get string containing our js code to pass for evaluation
		// jsInject, err := os.ReadFile("config/scripts/script.js")
		// if err != nil {
		// 	log.Fatalf("Error reading JS file: %v", err)
		// }
		// jsWebDriver, err := os.ReadFile("config/scripts/config.js")
		// if err != nil {
		// 	log.Fatalf("Error reading JS file: %v", err)
		// }

		err := chromedp.Run(ctx,
			chromedp.Navigate(url),
			// trying to evade webdriver detection
			chromedp.Evaluate(string(ConfigJS), nil),

			// that's to avoid triggering cloudflare/site protections. 1 sec is enough but the longer the better.
			// you can even skip that line and comment it, depends on site tolerance, but I don't recommend doing so.
			chromedp.Sleep(config.Delay),
			chromedp.Title(&pageTitle), // get page title to check if we got blocked
			// that js script goes straight to console of page and retrieves whatever selectors you write (its configurable)
			chromedp.Evaluate(string(ScriptJS), &rawData),
		)

		if err != nil {
			fmt.Printf("\nchromedp run failed: %v\n", err)
			continue
		}

		if strings.Contains(strings.ToLower(pageTitle), "verify") ||
			strings.Contains(strings.ToLower(pageTitle), "human") ||
			strings.Contains(strings.ToLower(pageTitle), "just a moment") {
			fmt.Printf("\n[!] We got detected %s\n", url)
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
	}

	uniqueSkins := internal.RemoveDuplicates(allSkins)
	return uniqueSkins, nil
}
