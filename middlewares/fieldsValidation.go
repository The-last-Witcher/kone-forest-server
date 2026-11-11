package middlewares

import (
	"kone-server/config"
)

func ValidateFields(req config.InquiryReq) bool{
	if req.FirstName == "" || req.LastName == "" || req.Email == "" || req.Topic == "" || req.Message == ""{
		return false
	}
	return true
}