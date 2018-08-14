package crawler

import (
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"../utils"
)

func RandomAgents() string {

	rand.Seed(time.Now().Unix())
	var ua = []string{
		"Mozilla/5.0 (Amiga; U; AmigaOS 1.3; en; rv:1.8.1.19) Gecko/20081204 SeaMonkey/1.1.14",
		"Mozilla/5.0 (AmigaOS; U; AmigaOS 1.3; en-US; rv:1.8.1.21) Gecko/20090303 SeaMonkey/1.1.15",
		"Mozilla/5.0 (AmigaOS; U; AmigaOS 1.3; en; rv:1.8.1.19) Gecko/20081204 SeaMonkey/1.1.14",
		"Mozilla/5.0 (Android 2.2; Windows; U; Windows NT 6.1; en-US) AppleWebKit/533.19.4 (KHTML, like Gecko) Version/5.0.3 Safari/533.19.4",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/532.0 (KHTML, like Gecko) Chrome/4.0.208.0 Safari/532.0",
	}

	return ua[rand.Intn(len(ua)-1)]

}

func sendRequests(url string) {

	ua := RandomAgents()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   3 * time.Second,
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", ua)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("=> Try %s -- %d \n", url, resp.StatusCode)

	if resp.Header.Get("Location") != "" {
		fmt.Printf("[+] Redirect found: %s", url)
	}

}

//CrawlerCommonFiles make requests to endpoints
func CrawlerCommonFiles(url string) {

	var commonFiles = []string{
		"robots.txt",
		"sitemap.xml",
		"sitemap.xml.gz",
		"crossdomain.xml",
		"phpinfo.php",
		"test.php",
		"elmah.axd",
		"graphql",
		"server-status",
		"jmx-console",
		"admin-console",
		"web-console",
		"swagger.json",
		"api.json",
		"trace.axd",
		"admin",
		"app.js",
		"aws.js",
	}

	for _, endpoint := range commonFiles {
		schema := url + "/" + endpoint
		sendRequests(schema)
		continue

	}

}

// GetCommonHeaders get headers of response
func GetCommonHeaders(url string) {
	fmt.Println(utils.Yellow("[+] Try to get some headers"))
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   3 * time.Second,
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	headers := resp.Header
	for k, v := range headers {

		if strings.Contains(k, "X-Powered-By") {
			fmt.Printf("%s  :  %s", k, v)
			fmt.Println()
		}

		if strings.Contains(k, "Server") {
			fmt.Printf("%s  :  %s", k, v)
			fmt.Println()
		}

		if strings.Contains(k, "Access-Control-Allow-Origin") {
			fmt.Printf("%s  :  %s", k, v)
			fmt.Println()
		}

		if strings.Contains(k, "Access-Control-Allow-Credential") {
			fmt.Printf("%s  :  %s", k, v)
			fmt.Println()
		}

		if strings.Contains(k, "Access-Control-Allow-Methods") {
			fmt.Printf("%s  :  %s", k, v)
			fmt.Println()
		}

		if strings.Contains(k, "X-XSS-Protection") {
			fmt.Printf("%s  :  %s", k, v)
			fmt.Println(utils.Blue("Maybe some XSS atacks not be effective.."))
			fmt.Println()
		}

		if strings.Contains(k, "strict-transport-security") {
			fmt.Printf("%s  :  %s", k, v)
			fmt.Println(utils.Blue("HSTS attacks.. not be effective.."))
			fmt.Println()
		}

	}

}
