package model

import (
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/harbor/src/pkg/notification/policy/model"
	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// WebhookPolicy ...
type WebhookPolicy struct {
	*model.Policy
}

// ToSwagger ...
func (n *WebhookPolicy) ToSwagger() *models.WebhookPolicy {
	return &models.WebhookPolicy{
		ID:           n.ID,
		CreationTime: strfmt.DateTime(n.CreationTime),
		UpdateTime:   strfmt.DateTime(n.UpdateTime),
		Creator:      n.Creator,
		Description:  n.Description,
		Enabled:      n.Enabled,
		EventTypes:   n.EventTypes,
		Name:         n.Name,
		ProjectID:    n.ProjectID,
		Targets:      n.ToTargets(),
	}
}

// ToTargets ...
func (n *WebhookPolicy) ToTargets() []*models.WebhookTargetObject {
	var results []*models.WebhookTargetObject
	for _, t := range n.Targets {
		results = append(results, &models.WebhookTargetObject{
			Type:           t.Type,
			Address:        t.Address,
			AuthHeader:     t.AuthHeader,
			SkipCertVerify: t.SkipCertVerify,
		})
	}
	return results
}

// NewNotifiactionPolicy ...
func NewWebhookPolicy(p *model.Policy) *WebhookPolicy {
	return &WebhookPolicy{
		Policy: p,
	}
}
