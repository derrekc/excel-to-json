package main

import (
	"fmt"
	"strings"
)

// Recommendation will be translated to JSON format
type Recommendation struct {
	Title              string
	Description        string
	Source             string
	RecommendationDate string
	Category           string
	EGDecision         string
	ResponsibleParty   string
	CurrentStatus      string
	ProgressToDate     string
	CampusPartnerships string
}

// String pretty-prints the recommendation
func (r Recommendation) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Title: %s\n", r.Title))
	b.WriteString(fmt.Sprintf("Recommendation:\n%s\n", r.Description))
	b.WriteString(fmt.Sprintf("Source: %s\n", r.Source))
	b.WriteString(fmt.Sprintf("Recommendation Date: %s\n", r.RecommendationDate))
	b.WriteString(fmt.Sprintf("Category: %s\n", r.Category))
	b.WriteString(fmt.Sprintf("EG Decision: %s\n", r.EGDecision))
	b.WriteString(fmt.Sprintf("Responsible Party/Parties: %s\n", r.ResponsibleParty))
	b.WriteString(fmt.Sprintf("Current Status: %s\n", r.CurrentStatus))
	b.WriteString(fmt.Sprintf("Progress To Date:\n%s\n", r.ProgressToDate))
	b.WriteString(fmt.Sprintf("Campus Partnerships: %s\n", r.CampusPartnerships))

	return b.String()
}

// RowTitle for later use
const (
	RowTitle              = 0x00
	RowDescription        = 0x01
	RowSource             = 0x02
	RowRecommendationDate = 0x03
	RowCategory           = 0x04
	RowEGDecision         = 0x05
	RowResponsibleParty   = 0x06
	RowCurrentStatus      = 0x07
	RowProgressToDate     = 0x08
	RowCampusPartnerships = 0x09
	RowExtra              = 0x0a
)

// RecommendationService manages those Recommendations
type RecommendationService struct {
	Recommendations []Recommendation
}

// NewRecommendationService can be called from main or within a DIG container
func NewRecommendationService() *RecommendationService {
	return &RecommendationService{}
}

// AddRecommendationFromSpreadsheetRow accepts a row of spreadsheet
// cell data
func (rs *RecommendationService) AddRecommendationFromSpreadsheetRow(row []string) {
	r := Recommendation{}
	r.Title = row[RowTitle]
	r.Description = row[RowDescription]
	r.Source = row[RowSource]
	r.RecommendationDate = row[RowRecommendationDate]
	r.Category = row[RowCategory]
	r.EGDecision = row[RowEGDecision]

	if len(row) > RowResponsibleParty {
		r.ResponsibleParty = row[RowResponsibleParty]
	}

	if len(row) > RowCurrentStatus {
		r.CurrentStatus = row[RowCurrentStatus]
	}

	if len(row) > RowProgressToDate {
		r.ProgressToDate = row[RowProgressToDate]
	}

	if len(row) > RowCampusPartnerships {
		r.CampusPartnerships = row[RowCampusPartnerships]
	}

	rs.Recommendations = append(rs.Recommendations, r)
	fmt.Println(r)
}
