package teamwork

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

const emptyString = ""

type DomainOccur struct {
	domain string
	occur  uint
}

func sortEmailsWithOccurs(filename string) []DomainOccur {

	file := openFile(filename)
	defer closeFile(file)
	reader := csv.NewReader(file)
	domainsOccurs := make(map[string]uint)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Error while reading the line:", err)
			break
		}
		if len(row) > 0 {
			wg.Add(1)
			go func(email string) {
				defer wg.Done()
				domain := extractEmailDomain(email)
				if domain != emptyString {
					mutex.Lock()
					domainsOccurs[domain]++
					mutex.Unlock()
				}
			}(row[2])
		}
	}
	wg.Wait()
	var result []DomainOccur
	for domain, count := range domainsOccurs {
		result = append(result, DomainOccur{domain: domain, occur: count})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].domain < result[j].domain
	})

	return result

}

func extractEmailDomain(email string) string {

	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	match := emailRegex.FindString(email)
	if match == emptyString {
		log.Println("No email address found in the input string:", email)
		return emptyString
	}
	parts := strings.Split(match, "@")
	if len(parts) != 2 {
		log.Println("Invalid email address format", email)
		return emptyString
	}
	return parts[1]

}

func openFile(filename string) *os.File {

	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening file:", err)
		return nil
	}
	return file

}

func closeFile(file *os.File) {

	err := file.Close()
	if err != nil {
		log.Println("Error while closing file:", err)
	}

}
