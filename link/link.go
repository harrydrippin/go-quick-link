package link

import (
	"errors"
	"sync"
)

// QuickLink represents the link itself
type QuickLink struct {
	Command    string `json:"command"`
	ParamCount int    `json:"param_count"`
	Path       string `json:"path"`
}

var defaultQuickLink = QuickLink{"", 0, ""}

// Query represents the search conditions
type Query struct {
	Command    string
	ParamCount int
}

// Database stores many QuickLinks inside the map
type Database struct {
	sync.Mutex
	queryMap map[Query]QuickLink
}

var instance *Database
var singleton sync.Once

var errNotFound = errors.New("No quick link has found by given query")

// GetDatabase obtains a singleton instance of linkDatabase
func GetDatabase() *Database {
	singleton.Do(func() {
		instance = &Database{
			queryMap: make(map[Query]QuickLink),
		}
	})
	return instance
}

// GetQuickLinks returns every QuickLink on the database
func (db *Database) GetQuickLinks() []QuickLink {
	var result []QuickLink
	for key := range db.queryMap {
		result = append(result, db.queryMap[key])
	}
	return result
}

// AddQuickLink adds given QuickLink to the database
func (db *Database) AddQuickLink(link QuickLink) error {
	db.Lock()
	defer db.Unlock()

	// TODO(@harry): Need to check if command is URL-compatible
	query := Query{
		Command:    link.Command,
		ParamCount: int(link.ParamCount),
	}

	db.queryMap[query] = link

	return nil
}

// RemoveQuickLink removes the QuickLink by given Query
func (db *Database) RemoveQuickLink(query Query) error {
	db.Lock()
	defer db.Unlock()

	_, exist := db.queryMap[query]

	if !exist {
		return errNotFound
	}

	delete(db.queryMap, query)

	return nil
}

// QueryQuickLink finds the QuickLink by given Query
func (db *Database) QueryQuickLink(query Query) (QuickLink, error) {
	quickLink, exist := db.queryMap[query]

	if !exist {
		return defaultQuickLink, errNotFound
	}

	return quickLink, nil
}
