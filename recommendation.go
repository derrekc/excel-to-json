package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// Recommendation will be translated to JSON format
type Recommendation struct {
	Title              string `json:"title"`
	Description        string `json:"body"`
	Source             string `json:"source"`
	RecommendationDate string `json:"recommendation_date"`
	Category           string `json:"recommendation_category"`
	EGDecision         string `json:"eg_decision"`
	ResponsibleParty   string `json:"responsible_party"`
	CurrentStatus      string `json:"current_status"`
	ProgressToDate     string `json:"recommendation_progress"`
	CampusPartnerships string `json:"campus_partnerships"`
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
	Recommendations       []Recommendation  `json:"recommendations"`
	CategoryTerms         map[string]string `json:"category_terms"`
	ResponsiblePartyTerms map[string]string `json:"responsible_party_terms"`
}

// NewRecommendationService can be called from main or within a DIG container
func NewRecommendationService() *RecommendationService {
	return &RecommendationService{
		CategoryTerms:         map[string]string{},
		ResponsiblePartyTerms: map[string]string{},
	}
}

// AddRecommendationFromSpreadsheetRow accepts a row of spreadsheet
// cell data
func (rs *RecommendationService) AddRecommendationFromSpreadsheetRow(row []string) {
	r := Recommendation{}
	r.Title = row[RowTitle]
	r.Description = row[RowDescription]
	r.Source = row[RowSource]
	r.RecommendationDate = correctedDateString(row[RowRecommendationDate])
	r.Category = row[RowCategory]
	rs.captureTerms("CategoryTerms", r.Category)

	r.EGDecision = row[RowEGDecision]

	if len(row) > RowResponsibleParty {
		r.ResponsibleParty = row[RowResponsibleParty]
		rs.captureTerms("ResponsiblePartyTerms", r.ResponsibleParty)
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
}

func (rs *RecommendationService) captureTerms(which string, t string) {
	var pMap map[string]string

	if which == "CategoryTerms" {
		pMap = rs.CategoryTerms
	} else if which == "ResponsiblePartyTerms" {
		pMap = rs.ResponsiblePartyTerms
	}
	re := regexp.MustCompile(",\\s?")
	split := re.Split(t, -1)
	for i := range split {
		pMap[split[i]] = "y"
	}
}

// WriteJSON to some io.Writer, writer, provided by the caller.
// If "writer" is null, print to stdout
func (rs *RecommendationService) WriteJSON(w io.Writer) {
	if w == nil {
		w = os.Stdout
	}

	resultB, _ := json.MarshalIndent(rs, "", "  ")
	io.WriteString(w, string(resultB))
}

// Return a date string suitable for using with a
// Drupal "date" field (in this case, 'YYYY-MM-DD')
func correctedDateString(ds string) string {
	var splits []string
	dayOfMonth := "01"

	r, _ := regexp.Compile("^\\d{4}")
	if r.MatchString(ds) {
		splits = strings.Split(ds, "-")
		return fmt.Sprintf("%s-%s-%s", splits[0], splits[1], dayOfMonth)
	}

	return ds
}
