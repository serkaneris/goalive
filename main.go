package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

func checkSubdomain(subdomain string, verbose bool, timeout time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	protocols := []string{"http://", "https://"}

	for _, protocol := range protocols {
		url := protocol + subdomain
		client := http.Client{
			Timeout: timeout,
		}

		resp, err := client.Head(url)
		if err != nil {
			if verbose {
				if os.IsTimeout(err) {
					fmt.Printf("Timeout occurred while checking %s\n", url)
				} else {
					fmt.Printf("Error occurred while checking %s: %s\n", url, err.Error())
				}
			}
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode < 400 {
			if verbose {
				fmt.Printf("Active: %s\t Status Code: %d\n", url, resp.StatusCode)
			} else {
				fmt.Println(url)
			}
		} else {
			if verbose {
				fmt.Printf("Inactive: %s\t Status Code: %d\n", url, resp.StatusCode)
			}
		}
	}
}

func main() {
	inputFile := flag.String("i", "", "File containing subdomains")
	timeout := flag.Int("t", 5000, "Request timeout in milliseconds")
	verbose := flag.Bool("v", false, "Print active and inactive URLs with status code")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *verbose {
		// Print the number of logical CPUs available.
		fmt.Printf("Logical CPUs: %d\n", runtime.NumCPU())

		// Set the maximum number of CPUs that can be executing simultaneously and return the previous setting.
		prevMaxProcs := runtime.GOMAXPROCS(0)
		fmt.Printf("Previous GOMAXPROCS setting: %d\n", prevMaxProcs)
	}

	if *inputFile == "" {
		fmt.Println("Error: Input file not specified")
		os.Exit(1)
	}

	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Printf("Error: File not found: %s\n", *inputFile)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var subdomains []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			subdomains = append(subdomains, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	timeoutDuration := time.Duration(*timeout) * time.Millisecond
	var wg sync.WaitGroup

	for _, subdomain := range subdomains {
		wg.Add(1)
		go checkSubdomain(subdomain, *verbose, timeoutDuration, &wg)
	}

	wg.Wait()
}





// package main

// import (
// 	"bufio"
// 	"flag"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"time"
// )

// func checkSubdomain(subdomain string, verbose bool, timeout time.Duration) {
// 	protocols := []string{"http://", "https://"}

// 	for _, protocol := range protocols {
// 		url := protocol + subdomain
// 		client := http.Client{
// 			Timeout: timeout,
// 		}

// 		resp, err := client.Head(url)
// 		if err != nil {
// 			if verbose {
// 				if os.IsTimeout(err) {
// 					fmt.Printf("%s adresini kontrol ederken zaman aşımı oldu\n", url)
// 				} else {
// 					fmt.Printf("%s adresini kontrol ederken hata oluştu: %s\n", url, err.Error())
// 				}
// 			}
// 			continue
// 		}
// 		defer resp.Body.Close()

// 		if resp.StatusCode < 400 {
// 			if verbose {
// 				fmt.Printf("Aktif: %s\t Durum Kodu: %d\n", url, resp.StatusCode)
// 			} else {
// 				fmt.Println(url)
// 			}
// 		} else {
// 			if verbose {
// 				fmt.Printf("Pasif: %s\t Durum Kodu: %d\n", url, resp.StatusCode)
// 			}
// 		}
// 	}
// }

// func main() {
// 	inputFile := flag.String("i", "", "Alt alan adlarını içeren dosya")
// 	timeout := flag.Int("t", 5000000, "Request timeout millisecond")
// 	verbose := flag.Bool("v", false, "Aktif ve pasif URL'leri durum kodu ile birlikte yazdır")
// 	flag.Parse()

// 	if *inputFile == "" {
// 		fmt.Println("Hata: Girdi dosyası belirtilmedi")
// 		os.Exit(1)
// 	}

// 	file, err := os.Open(*inputFile)
// 	if err != nil {
// 		fmt.Printf("Hata: Dosya bulunamadı: %s\n", *inputFile)
// 		os.Exit(1)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	var subdomains []string
// 	for scanner.Scan() {
// 		line := strings.TrimSpace(scanner.Text())
// 		if line != "" {
// 			subdomains = append(subdomains, line)
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Printf("Hata: %s\n", err.Error())
// 		os.Exit(1)
// 	}

// 	timeoutDuration := time.Duration(*timeout) * time.Microsecond

// 	for _, subdomain := range subdomains {
// 		checkSubdomain(subdomain, *verbose, timeoutDuration)
// 	}
// }	
