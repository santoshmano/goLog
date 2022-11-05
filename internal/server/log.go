package server

import (
	"fmt"
	"sync"
)

// Log to record the events
type Log struct {
	mu      sync.Mutex
	records []Record
}

// Record - Individual records in the log
type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

func NewLog() *Log {
	return &Log{}
}

// Append - write to the log
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)

	return record.Offset, nil
}

// Read - read from the log
func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}

	return c.records[offset], nil
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")
