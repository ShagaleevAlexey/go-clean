package access_token

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

var sigTests = []string{
	"asd",
}

func Test_generateTokens(t *testing.T) {
	for _, sig := range sigTests {
		uid := uuid.NewV4()

		at, rt, _, err := GenerateTokens(uid, []byte(sig))
		if err != nil {
			t.Errorf("Generate token has error '%s'", err)
		}

		atClaims, err := ValidAccessToken(at, []byte(sig))
		if err != nil {
			t.Errorf("Validation access token has error '%s'", err)
		}
		if atClaims.Uid != uid.String() {
			t.Errorf("Uid '%s' not matched with access token uid '%s'", uid, atClaims.Uid)
		}

		rtClaims, err := ValidAccessToken(rt, []byte(sig))
		if err != nil {
			t.Errorf("Validation refresh token has error '%s'", err)
		}
		if rtClaims.Uid != uid.String() {
			t.Errorf("Uid '%s' not matched with refresh token uid '%s'", uid, rtClaims.Uid)
		}
	}
}

func Test_validateTokens(t *testing.T) {
	var err error
	sig := sigTests[0]
	fakeSig := "jkasdkj"
	uid := uuid.NewV4()

	at, rt, _, err := GenerateTokens(uid, []byte(sig))
	if err != nil {
		t.Errorf("Generate token has error '%s'", err)
	}

	_, err = ValidAccessToken(at, []byte(fakeSig))
	if err == nil {
		t.Errorf("Validation access token not raised error with fake sig")
	}

	_, err = ValidAccessToken(rt, []byte(fakeSig))
	if err == nil {
		t.Errorf("Validation refresh token not raised error with fake sig")
	}
}


