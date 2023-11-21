package myaws

import (
	"testing"
)

func TestGetAccounts(t *testing.T) {
	GetAccountsList()
}

func TestGenConfig(t *testing.T) {
	GenConfig()
}

func TestGetRoot(t *testing.T) {
	getRoot()
}

func TestGenSSOConfig(t *testing.T) {
	GenSSOConfig("awsso", "awsso", "AWSAdministratorAccess")
}
