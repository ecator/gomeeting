package ldap

import (
	"fmt"
	"strings"

	ldapv3 "github.com/go-ldap/ldap/v3"
)

// Login check it can login,
// return userInfo{username,name,email} if success
func Login(addr, baseDN, userName, password string, attrMapKey map[string]string) (map[string]string, error) {
	var (
		conn          *ldapv3.Conn
		entry         *ldapv3.Entry
		searchRequest *ldapv3.SearchRequest
		searchResult  *ldapv3.SearchResult
		err           error
	)
	info := map[string]string{}
	info["username"] = userName
	info["name"] = userName
	info["email"] = userName + "@example.com"
	conn, err = ldapv3.Dial("tcp", addr)
	if err != nil {
		return info, err
	}
	err = conn.Bind("cn="+userName+","+baseDN, password)
	if err != nil {
		conn.Close()
		return info, err
	}
	searchRequest = ldapv3.NewSearchRequest(
		baseDN,
		ldapv3.ScopeWholeSubtree, ldapv3.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(cn=%s))", userName),
		nil,
		nil,
	)

	searchResult, err = conn.Search(searchRequest)
	if err != nil {
		conn.Close()
		return info, err
	}

	if len(searchResult.Entries) != 1 {
		conn.Close()
		return info, fmt.Errorf("User does not exist or too many entries returned")
	}

	entry = searchResult.Entries[0]
	for _, attr := range entry.Attributes {
		for k, v := range attrMapKey {
			if strings.ToLower(attr.Name) == strings.ToLower(v) && len(attr.Values) > 0 {
				info[k] = attr.Values[0]
			}
		}
	}

	conn.Close()

	return info, nil
}
