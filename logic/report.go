package logic

import "time"

// Report represent the reports listed in the get reports loader
type Report struct {
	ID               string    `json:"id"`
	Source           string    `json:"source"`
	SourceIdentityID string    `json:"source_identity_id"`
	Reference        Reference `json:"reference"`
	State            string    `json:"state"`
	Payload          Payload   `json:"payload"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// Reference represents the field of Report
type Reference struct {
	ReferenceID   string `json:"reference_id"`
	ReferenceType string `json:"reference_type"`
}

// Payload represent a field of Report
type Payload struct {
	Source                string `json:"source"`
	ReportType            string `json:"report_type"`
	Message               string `json:"message"`
	ReportID              string `json:"report_id"`
	ReferenceResourceID   string `json:"reference_resource_id"`
	ReferenceResourceType string `json:"reference_resource_type"`
}
