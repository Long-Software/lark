package controllers

import "github.com/Long-Software/Bex/apps/cmd/everia/utils"

type PasswordController struct{}

func (p *PasswordController) UnlockVault(password string) map[string]interface{} {
	data, err := utils.LoadEncryptedFile("vault.enc")
	if err != nil {
		return map[string]interface{}{"success": false}
	}

	vData, err := utils.DecryptVault(data, password)
	if err != nil {
		return map[string]interface{}{"success": false}
	}

	return map[string]interface{}{
		"success": true,
		"entries": vData.Entries,
	}
}