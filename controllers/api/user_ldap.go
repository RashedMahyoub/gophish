package api

import (
	"log"
	"net/http"

	"github.com/jtblin/go-ldap-client"
)

// Users contains functions to retrieve a list of existing users or create a
// new user. Users with the ModifySystem permissions can view and create users.
func (as *Server) ImportLdapUsers(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		client := &ldap.LDAPClient{
			Base:         "dc=resilience,dc=local",
			Host:         "192.168.0.204",
			Port:         389,
			UseSSL:       false,
			BindDN:       "uid=readonlysuer,ou=People,dc=resilience,dc=local",
			BindPassword: "Admin@123",
			UserFilter:   "(uid=%s)",
			GroupFilter: "(memberUid=%s)",
			Attributes:   []string{"givenName", "sn", "mail", "uid"},
		}
		// It is the responsibility of the caller to close the connection
		defer client.Close()
	
		ok, user, err := client.Authenticate("Administrator", "Admin@123")
		if err != nil {
			log.Fatalf("Error authenticating user %s: %+v", "username", err)
		}
		if !ok {
			log.Fatalf("Authenticating failed for user %s", "username")
		}
		log.Printf("User: %+v", user)
		
		groups, err := client.GetGroupsOfUser("Administrator")
		if err != nil {
			log.Fatalf("Error getting groups for user %s: %+v", "username", err)
		}
		log.Printf("Groups: %+v", groups) 
		JSONResponse(w, groups, http.StatusOK)
		return
	case r.Method == "POST":
		return
	}
}

