package formats

import "time"

type BitwardenFile struct {
	Folders []BitwardenFolder `json:"folders,omitempty"`
	Items   []BitwardenItem   `json:"items"`
}

type BitwardenFolder struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type BitwardenItemType uint8

const (
	BitwardenItemTypePassword = iota + 1
)

type BitwardenItem struct {
	Type            BitwardenItemType      `json:"type"`
	PasswordHistory []PasswordHistoryEntry `json:"password_history,omitempty"`
	RevisionDate    *time.Time             `json:"revisionDate,omitempty"`
	CreationDate    *time.Time             `json:"creationDate,omitempty"`
	Name            string                 `json:"name"`
	Notes           string                 `json:"notes"`
	Login           BitwardenLogin         `json:"login"`
	FolderID        string                 `json:"folderId,omitempty"`
}

type PasswordHistoryEntry struct {
	LastUsedDate *time.Time `json:"lastUsedDate,omitempty"`
}

type BitwardenLogin struct {
	URIs     []BitwardenLoginURI `json:"uris,omitempty"`
	Username string              `json:"username"`
	Password string              `json:"password"`
	// TOPT
}

type BitwardenLoginURI struct {
	Match *string `json:"match"`
	URI   string  `json:"uri"`
}

func BitwardenFileFromNorton(entries []NortonEntry) (BitwardenFile, error) {
	bf := BitwardenFile{}

	for _, entry := range entries {
		bf.Items = append(bf.Items, BitwardenItem{
			Type:  BitwardenItemTypePassword,
			Name:  entry.Title,
			Notes: entry.Notes,
			Login: BitwardenLogin{
				URIs: []BitwardenLoginURI{
					{
						URI: entry.URL,
					},
				},
				Username: entry.Username,
				Password: entry.Password,
			},
		})
	}

	return bf, nil
}
