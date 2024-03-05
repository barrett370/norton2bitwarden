package formats

import (
	"os"

	"github.com/gocarina/gocsv"
)

// Item Type,User Name,Password,Site Name,Login URL,Notes,Secure,Favorite

type NortonEntry struct {
	ItemType string `csv:"Item Type"`
	Username string `csv:"User Name"`
	Password string `csv:"Password"`
	Title    string `csv:"Site Name"`
	URL      string `csv:"Login URL"`
	Notes    string `csv:"Notes"`
	Secure   bool   `csv:"Secure"`
	Favorite bool   `csv:"Favorite"`
}

func DecodeNortonExport(filename string) ([]NortonEntry, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var out []NortonEntry
	err = gocsv.UnmarshalFile(f, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
