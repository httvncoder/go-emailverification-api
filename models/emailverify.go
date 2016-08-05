package models

import (
	"net"
	"time"
)

// Question struct
type Question struct {
	JobEmailAddress string    `json:"JobEmailAddress"`
	JobStatus       string    `json:"JobStatus"`
	JobMessage      string    `json:"JobMessage"`
	JobTime         time.Time `json:"JobTime"`
}

// Answer struct
type Answer struct {
	EmailAddress     string    `json:"EmailAddress"`
	EmailUser        string    `json:"EmailUser"`
	EmailDomain      string    `json:"EmailDomain"`
	ValidationResult string    `json:"ValidationResult"`
	MXRecords        []*net.MX `json:"MXRecords"`
}

// Message for retunring
type Message struct {
	Question Question `json:"Question"`
	Answer   Answer   `json:"Answer"`
}
