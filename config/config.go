// main configuration file where you can set options for scraper,
// add/remove subdomains, change constants for scraper settings etc.

package config

import (
	"time"

	"github.com/chromedp/chromedp"
)

// here you should add subdomains (in our case it's name of collection/case)
var List = []string{
	"kilowatt-case", "revolution-case", "recoil-case", "dreams-nightmares-case", "sealed-genesis-terminal",
	"snakebite-case", "fracture-case", "prisma-2-case", "cs20-case", "prisma-case", "danger-zone-case",
	"horizon-case", "clutch-case", "spectrum-2-case", "operation-hydra-case", "spectrum-case", "glove-case",
	"gamma-2-case", "gamma-case", "chroma-3-case", "operation-wildfire-case", "revolver-case", "shadow-case", "falchion-case",
	"chroma-2-case", "chroma-case", "operation-vanguard-weapon-case", "operation-breakout-weapon-case", "huntsman-weapon-case",
	"operation-phoenix-weapon-case", "csgo-weapon-case-3", "winter-offensive-weapon-case", "csgo-weapon-case-2", "operation-bravo-case",
	"csgo-weapon-case", "fever-case", "gallery-case", "operation-riptide-case", "operation-broken-fang-case", "shattered-web-case", "esports-2014-summer-case",
	"esports-2013-winter-case", "esports-2013-case", "anubis-collection-package", "x-ray-p250-package",
}

// Scraper settings
const (
	DeadLine       = 120 * time.Second // time limit for context
	UrlLengthLimit = 60                // shorten url to specified length
	Delay          = 1 * time.Second   // delay to avoid triggering site protections
)

// allocator options
var Opts = append(chromedp.DefaultExecAllocatorOptions[:],
	chromedp.Flag("headless", true),
	chromedp.Flag("disable-blink-features", "AutomationControlled"),
	chromedp.Flag("exclude-switches", "enable-automation"),
	chromedp.Flag("disable-extensions", false),
	chromedp.Flag("start-maximized", false),
	chromedp.Flag("window-size", "800,600"),
	chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36"),
)
