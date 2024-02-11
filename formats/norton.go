package formats

import (
	"os"

	"github.com/gocarina/gocsv"
)

// Username, Password, Title, Login, URL, Notes

type NortonEntry struct {
	Username string `csv:"Username"`
	Password string `csv:"Password"`
	Title    string `csv:"Title"`
	Login    string `csv:"Login"`
	URL      string `csv:"URL"`
	Notes    string `csv:"Notes"`
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
