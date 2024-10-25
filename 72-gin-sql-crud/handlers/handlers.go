package handlers

import "demo/models"

var ChAudit chan models.Audit

func init() {
	ChAudit = make(chan models.Audit, 5)
}
