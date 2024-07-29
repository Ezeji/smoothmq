package models

import "net/http"

type TenantManager interface {
	GetTenant(r *http.Request) int64
	GetAWSSecretKey(accessKey string, region string) (tenantId int64, secretKey string, err error)
}
