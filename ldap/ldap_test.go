package ldap_test

import (
	"testing"

	"github.com/ecator/gomeeting/ldap"
)

func TestLogin(t *testing.T) {
	var (
		err  error
		info map[string]string
	)
	addr := "192.168.33.50:389"
	baseDN := "cn=users,dc=v0,dc=home"
	userName := "test1"
	password := "Asd123456"
	mapKey := map[string]string{
		"name":  "displayName",
		"email": "mail",
	}
	info, err = ldap.Login(addr, baseDN, userName, password, mapKey)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(info)
	}
}
