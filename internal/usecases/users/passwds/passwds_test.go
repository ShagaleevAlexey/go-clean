package passwds

import "testing"

var passwordsTests = []string{
	"test",
	"admin",
	"moderator",
}

func Test_passwordHash(t *testing.T) {
	var err error
	for _, pass := range passwordsTests {
		res := GetPasswordHash(pass)
		err = MatchPasswordHash(res, pass)
		if err != nil {
			t.Errorf("Password '%s' not matched with '%s'", pass, res)
		}
	}
}
