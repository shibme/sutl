package utils

import (
	"errors"

	"xipher.org/xipher"
)

const (
	appNameLowerCase      = "xipher"
	xipherPublicKeyPrefix = "XPK_"
	xipherSecretKeyPrefix = "XSK_"
	xipherTxtPrefix       = "XCT_"
	xipherPubKeyFileExt   = ".xpk"
	xipherFileExt         = "." + appNameLowerCase
	secretKeyStrRegex     = "^" + xipherSecretKeyPrefix + "[A-Z2-7]{106}$"
)

var (
	pwdSecretKeyMap         = make(map[string]*xipher.SecretKey)
	preferredKeyQueryParams = []string{"xk", "xpw", "pw", "xw"}

	errInvalidXipherPubKey    = errors.New("invalid xipher public key")
	errInvalidXipherSecretKey = errors.New("invalid xipher secret key")
	errInvalidCipherText      = errors.New("invalid cipher text")
)
