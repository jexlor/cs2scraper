## Cs2Scraper  [![PkgGoDev](https://pkg.go.dev/badge/github.com/jexlor/cs2scraper)](https://pkg.go.dev/github.com/jexlor/cs2scraper) (Beta)

--------------------------------------------

<strong>cs2scraper</strong> is a scraping tool to retrieve fresh data and wrap it up in convenient JSON file with clean structure. The idea is simple, you run script to target already pre configured target site and get data.

Example structure:

```json
{
    "name": "Inheritance",
    "weapon": "AK-47",
    "rarity": "Covert Rifle",
    "collection": "The Kilowatt Collection",
    "price": {
      "price_string": "$48.62 - $302.37",
      "price_stattrak_string": "$91.61 - $694.48",
      "currency": "USD",
      "min": {
        "value": 48.62,
        "stattrak_value": 91.61,
        "unit": "USD"
      },
      "max": {
        "value": 302.37,
        "stattrak_value": 694.48,
        "unit": "USD"
      },
      "updated_at": "2025-11-09T20:31:26+04:00"
    },
    "url": "https://www.example.com/images/AK-47_Inheritance.png"
}
```

## Ways of Use
--------------------------------------------

1. **Run locally** to generate a full JSON file containing all scraped skin data.
2. **Import the package as a library** to get a populated `[]Skin` struct returned directly from your Go code.
3. **Run as a Docker container** to execute the scraper in an isolated environment.  
   Useful for automated tasks, cron jobs, microservices, or running on servers without installing Go/Chromium locally.


<strong>Disclaimer!</strong>
--------------------------------------------

Since the script subtly “unofficially” targets online databases, in our case the popular site <a href="https://www.csgodatabase.com/">CSGO Database</a>, you must understand that all responsibility lies with you. Whether you use the data for a hobby project, a real website, or any form of monetization, you are solely accountable for the consequences.

Even though the script simply retrieves data that is already visible to you, but script is developed the way to bypass anti-bot systems, which can be considered a violation of the target website’s terms of service.

## About Technical Stuff
--------------------------------------------
The scraper works **synchronously** because the chance of success when running a headless browser with goroutines is virtually zero if the target site has any protection. In our case, the target has **Cloudflare** and strong anti-bot systems.  

The script uses a Chromium instance to make synchronous requests to pre-configured URLs, injects custom JavaScript to retrieve data specified by selectors, and then navigates to the next URL in sequence. This approach ensures reliable scraping while minimizing detection.

## Why Chromium?
--------------------------------------------
The scraper uses the `chromedp` library because it provides a **high-level interface for controlling Chromium** programmatically. This allows us to:

- Navigate websites like a real browser.
- Execute custom JavaScript on the page to extract dynamic content.
- Bypass certain basic anti-bot protections that block simple HTTP requests.
- Wait for elements to load and handle asynchronous content reliably.
- And it helps disguise the script as a real user because it operates through an actual browser.

## What you should have 
--------------------------------------------

1. Golang obviously
2. Chromium `sudo apt install chromium-browser`


## Now let's get started 
--------------------------------------------

## If you want to run script locally and get JSON

1. **Clone repository**

```bash
git clone https://github.com/jexlor/cs2scraper.git
```

2. **Ensure that dependencies are installed**

```bash
go mod tidy
```

3. **Run script in root**

```bash
go run .
```

## If you want to import script as an library 

1. **Import cs2scraper**

```bash
go get github.com/jexlor/cs2scraper
```

2. **Call func**

```go
skins, err := scraper.ScrapSkins()
```

3. **Configure scraper settings(optional)**

Here is literally everything that you can modify if you want

```go
	import "github.com/jexlor/cs2scraper/scraper/config"
```

## Future Plans & Current Situation
--------------------------------------------

Our main goal is to make the scraper as **flexible and configurable as possible**, so that virtually everything can be adjusted **without modifying the scraping logic itself**.  

Different targets use different classes, protections, subdomains, URLs, and overall page structures, which makes building a fully flexible scraper quite challenging.  

For now, the scraper works and is fairly configurable if you download it locally. Expanding it to new targets requires investigating their classes and page structures. The ultimate goal is to allow scraping new targets **by changing only the configuration**, without touching the core logic. As the scraper grows and begins to include scraping for **Agents, Stickers, Souvenir packages**, and other items, it is becoming even more challenging to maintain flexibility and configurability.

## Community Contributions
--------------------------------------------
The project greatly benefits from community input. If you have suggestions, improvements, or want to help expand support for new targets, configurations, or features, your contribution is welcome.  

Whether it's reporting issues, improving scripts, or submitting pull requests, every helping hand makes the scraper more powerful and flexible for everyone.

*— Jexlor*
