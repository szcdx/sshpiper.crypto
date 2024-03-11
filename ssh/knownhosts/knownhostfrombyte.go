package knownhosts

import (
	"golang.org/x/crypto/ssh"
)

func NewFromByte(line []byte, filename string, linenum int) (ssh.HostKeyCallback, error) {
	db := newHostKeyDB()
	var err = db.parseLine(line, filename, linenum)
	if err != nil {
		return nil, err
	}

	var certChecker ssh.CertChecker
	certChecker.IsHostAuthority = db.IsHostAuthority
	certChecker.IsRevoked = db.IsRevoked
	certChecker.HostKeyFallback = db.check

	return certChecker.CheckHostKey, nil
}
