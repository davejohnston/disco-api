// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Host struct {
	ID                       string                   `json:"id"`
	ParentID                 *string                  `json:"parentID"`
	Name                     string                   `json:"name"`
	Hostname                 *string                  `json:"hostname"`
	PrivateIPv4              string                   `json:"privateIPv4"`
	PrivateIPv6              *string                  `json:"privateIPv6"`
	PublicIPv4               *string                  `json:"publicIPv4"`
	PublicIPv6               *string                  `json:"publicIPv6"`
	DNSName                  *string                  `json:"dnsName"`
	PublicDNSName            *string                  `json:"publicDnsName"`
	CreatedAt                time.Time                `json:"createdAt"`
	UpdatedAt                *time.Time               `json:"updatedAt"`
	VulnerabiltiesConnection *VulnerabilityConnection `json:"vulnerabiltiesConnection"`
}

type HostConnection struct {
	TotalCount int          `json:"totalCount"`
	PageInfo   PageInfo     `json:"pageInfo"`
	Hosts      []*Host      `json:"hosts"`
	Edges      []*HostsEdge `json:"edges"`
}

type HostsEdge struct {
	Node   *Host  `json:"node"`
	Cursor string `json:"cursor"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

type VulnerabilitiesEdge struct {
	Node   *Vulnerability `json:"node"`
	Cursor string         `json:"cursor"`
}

type Vulnerability struct {
	ID                  string     `json:"id"`
	Name                *string    `json:"name"`
	RiskScore           *int       `json:"riskScore"`
	Summary             *string    `json:"summary"`
	Solution            *string    `json:"solution"`
	Cves                *string    `json:"cves"`
	CvssBaseScore       *int       `json:"cvssBaseScore"`
	CvssTemporalScore   *int       `json:"cvssTemporalScore"`
	CvssV3BaseScore     *int       `json:"cvssV3BaseScore"`
	CvssV3TemporalScore *int       `json:"cvssV3TemporalScore"`
	PatchAvailable      *bool      `json:"patchAvailable"`
	Exploitable         *bool      `json:"exploitable"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt"`
}

type VulnerabilityConnection struct {
	TotalCount      int                    `json:"totalCount"`
	PageInfo        PageInfo               `json:"pageInfo"`
	Vulnerabilities []*Vulnerability       `json:"vulnerabilities"`
	Edges           []*VulnerabilitiesEdge `json:"edges"`
}
