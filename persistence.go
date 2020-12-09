package pcookiejar

import (
	"encoding/json"
	"time"
)

// Jar implements the http.CookieJar interface from the net/http package.
type PersistentJar struct {
	Entries    map[string]map[string]entry `json:"entries"`
	NextSeqNum uint64                      `json:"nextSeqNum"`
}

func (j *Jar) MarshalJSON() ([]byte, error) {
	return json.Marshal(&PersistentJar{
		Entries:    j.entries,
		NextSeqNum: j.nextSeqNum,
	})
}

func (j *Jar) UnmarshalJSON(data []byte) error {
	var Jar PersistentJar
	if err := json.Unmarshal(data, &Jar); err != nil {
		return err
	}
	j.entries = Jar.Entries
	j.nextSeqNum = Jar.NextSeqNum
	return nil
}

type PersistentEntry struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	Domain     string `json:"domain"`
	Path       string `json:"path"`
	SameSite   string `json:"sameSite"`
	Secure     bool `json:"secure"`
	HttpOnly   bool `json:"httpOnly"`
	Persistent bool `json:"persistent"`
	HostOnly   bool `json:"hostOnly"`
	Expires    time.Time `json:"expires"`
	Creation   time.Time `json:"creation"`
	LastAccess time.Time `json:"lastAccess"`
	SeqNum     uint64 `json:"seqNum"`
}

func (e entry) MarshalJSON() ([]byte, error) {
	return json.Marshal(PersistentEntry{
		Name:       e.Name,
		Value:      e.Value,
		Domain:     e.Domain,
		Path:       e.Path,
		SameSite:   e.SameSite,
		Secure:     e.Secure,
		HttpOnly:   e.HttpOnly,
		Persistent: e.Persistent,
		HostOnly:   e.HostOnly,
		Expires:    e.Expires,
		Creation:   e.Creation,
		LastAccess: e.LastAccess,
		SeqNum:     e.seqNum,
	})
}

func (e *entry) UnmarshalJSON(data []byte) error {
	var Entry PersistentEntry
	if err := json.Unmarshal(data, &Entry); err != nil {
		return err
	}
	e.Name = Entry.Name
	e.Value = Entry.Value
	e.Domain = Entry.Domain
	e.Path = Entry.Path
	e.SameSite = Entry.SameSite
	e.Secure = Entry.Secure
	e.HttpOnly = Entry.HttpOnly
	e.Persistent = Entry.Persistent
	e.HostOnly = Entry.HostOnly
	e.Expires = Entry.Expires
	e.Creation = Entry.Creation
	e.LastAccess = Entry.LastAccess
	e.seqNum = Entry.SeqNum
	return nil
}
