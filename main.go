package main

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/datatypes"
	"net/http"
	"sync"
	"time"
)

type MtbCall func(d *Data, wg *sync.WaitGroup, errChan chan error, done chan bool)

type Data struct {
	Nom int
	Per int
	Doc string
}

func AddNominee(d *Data, wg *sync.WaitGroup, errChan chan error, done chan bool) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Println("AddNominee: error occurred")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("AddNominee: ", d.Nom)
		}
	}
}

func AddData(d *Data, wg *sync.WaitGroup, errChan chan error, done chan bool) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Println("AddData : error occurred")
			return
		default:
			fmt.Println("AddData: ", d.Per)
			time.Sleep(5 * time.Second)
			fmt.Println("AddData1: ", d.Per)
		}
	}
}

func UploadDocument(d *Data, wg *sync.WaitGroup, errChan chan error, done chan bool) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Println("UploadDocument : error occurred")
			return
		default:
			fmt.Println("UploadDocument: ", d.Doc)
			errChan <- errors.New("UploadDocument error")
			close(done)
		}
	}
}

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		time.Sleep(5 * time.Second)
		fmt.Println("doing work...")
	}
}

func doJob(ticker *time.Ticker) {
	for t := range ticker.C {
		fmt.Println("doJob: ", t)
	}
}

type Steps struct {
	On One `json:"1"`
	Tw Two `json:"2"`
}

type One struct {
	CardTheme string `json:"card_theme"`
}

type Two struct {
	CardName string `json:"card_name"`
}

func service(name string) {
	headers := http.Header{
		"Content-Type": []string{"application/json"},
		"Accept":       []string{"*/*"},
	}

	headers.Add("user", name)

	fmt.Println(headers)
}

type Number struct {
	Id string `json:"id"`
}

type Person struct {
	Name string `json:"name,omitempty"`
}

type Combined struct {
	Number Number `json:"number,omitempty"`
	Person Person `json:"person,omitempty"`
}

type Details struct {
	ChunkId string                       `json:"chunk_id"`
	Com     datatypes.JSONType[CardInfo] `gorm:"type:jsonb"`
}

func SetBit(num, chunk int) int {
	return num | (1 << chunk)
}

func main() {

	chunkSub := 0
	for i := 0; i < 6; i++ {
		chunkSub = SetBit(chunkSub, i)
	}

	fmt.Println(chunkSub)

	//data := []string{"abir", "fahim", "others", "zunaed", "nishat"}
	//
	//// Separate "others" and sort the rest
	//sort.Strings(data)
	//
	//// Find and move "others" to the last position
	//for i, v := range data {
	//	if v == "others" {
	//		// Remove "others" from its current position
	//		data = append(data[:i], data[i+1:]...)
	//		// Add "others" to the end of the slice
	//		data = append(data, v)
	//		break
	//	}
	//}
	//
	//fmt.Println(data)

	//type pair struct {
	//	a int    `json:"a"`
	//	v string `json:"v"`
	//}
	//
	//data := []pair{
	//	{a: 100, v: "abir"},
	//	{a: 200, v: "fahim"},
	//	{a: 300, v: "others"},
	//	{a: 400, v: "zunaed"},
	//	{a: 500, v: "nishat"},
	//}
	//
	//sort.Slice(data, func(i, j int) bool {
	//	if data[i].v == "others" {
	//		return false
	//	}
	//	if data[j].v == "others" {
	//		return true
	//	}
	//	return data[i].v < data[j].v
	//})
	//
	//fmt.Println(data)
	// Find the submatches

	//"Home / Holding: -, Village / Road: , Tupamari, Post Office: Ramganj Hat-5300, Nilphamari Sadar, Nilphamari"
	//
	//address := "House: 10, Road: 12, Village: Mirpur, Thana: Kafrul, District: Dhaka, Postcode: 1216"
	//
	//address = "Home / Holding: Comilla Basti, Village / Road: P. Agargaon, Ward No. 2699, Post Office: Mohammadpur-1208, Sher-e-Bangla Nagar, Dhaka"

	// Regex to capture 4-digit postcode numbers

	// Regex to capture a 4-digit number
	//str := "Home / Holding: -, Village / Road: , Bahadurpur, Post Office: Dhanyakhola-6431, Sharsha, Jessore"

	// Regex to capture 4-digit numbers
	//re := regexp.MustCompile(`\b\d{4}\b`)

	// Find all 4-digit numbers in the string
	//re := regexp.MustCompile(`Post Office:.*(\d{4}),.*`)
	////re := regexp.MustCompile(`.*(\d{4}).*`)
	//
	//// Find the submatches
	//match := re.FindStringSubmatch(str)
	//if len(match) > 1 {
	//	fmt.Println(match[1])
	//} else {
	//	fmt.Println("not matched")
	//}

	//	// GORM PostgreSQL connection string
	//	dsn := "host=localhost user=username password=password dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//	if err != nil {
	//		log.Fatalf("Unable to connect to database: %v\n", err)
	//	}
	//
	//	// Function to fetch entries from the last 1 hour based on the latest entry timestamp
	//	fetchLastHourEntries := func() {
	//		var users []User
	//
	//		// Raw SQL query in GORM
	//		query := `
	//		SELECT id, name, status, created_at
	//		FROM users
	//		WHERE created_at >= (
	//			SELECT created_at FROM users ORDER BY created_at DESC LIMIT 1
	//		) - INTERVAL '1 hour';
	//	`
	//
	//		// Execute the raw SQL query using GORM
	//		result := db.Raw(query).Scan(&users)
	//
	//		if result.Error != nil {
	//			log.Printf("Query failed: %v\n", result.Error)
	//			return
	//		}
	//
	//		// Output the users fetched
	//		if len(users) > 0 {
	//			fmt.Println("Users added within the last hour based on the latest entry time:")
	//			for _, user := range users {
	//				fmt.Printf("ID: %d, Name: %s, Status: %s, CreatedAt: %s\n", user.ID, user.Name, user.Status, user.CreatedAt)
	//			}
	//		} else {
	//			fmt.Println("No users added within the last hour based on the latest entry.")
	//		}
	//	}
	//
	//	// Fetch entries every hour using a ticker
	//	ticker := time.NewTicker(1 * time.Hour)
	//	defer ticker.Stop()
	//
	//	// Fetch entries once immediately and then every hour
	//	fetchLastHourEntries()
	//	for {
	//		select {
	//		case <-ticker.C:
	//			fetchLastHourEntries()
	//		}
	//	}
	//}

}
