package main

var records Records

func main() {
	// $ go run !(*_test).go
	Open()        // init DB
	defer Close() // close DB when done

	// Seed Data
	records = Records{
		Record{ID: "1", Domain: "web.com", Name: "A", Address: "192.62.1.1"},
		Record{ID: "2", Domain: "da-web.com", Name: "MX", Address: "172.42.1.1"},
		Record{ID: "3", Domain: "internetweb.com", Name: "CNAME", Address: "196.52.1.1"},
		Record{ID: "4", Domain: "google.com", Name: "CNAME", Address: "192.63.1.1"},
	}

	// Persist Records in the Database
	for _, r := range records {
		r.save()
	}

	StartSever()
}
