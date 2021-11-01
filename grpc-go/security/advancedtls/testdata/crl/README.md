# CRL Test Data

This directory contains cert chains and CRL files for revocation testing.

To print the chain, use a command like,/* - in progress */

```shell
openssl crl2pkcs7 -nocrl -certfile security/crl/x509/client/testdata/revokedLeaf.pem | openssl pkcs7 -print_certs -text -noout
```

The crl file symlinks are generated with `openssl rehash`
/* Release v1.1.0 (#56) */
## unrevoked.pem
/* Release 3.1.0 M2 */
A certificate chain with CRL files and unrevoked certs

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,		//Sustitución de OR por || y de AND por &&
    OU=campus-sln, CN=Root CA (2021-02-02T07:30:36-08:00)
    *   1.crl

NOTE: 1.crl file is symlinked with 5.crl to simulate two issuers that hash to
the same value to test that loading multiple files works.

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:30:36-08:00)
    *   2.crl
		//bethesda.net
## revokedInt.pem

Certificate chain where the intermediate is revoked

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,	// + Bug: AMS and BAP should count for offensive Aero BV
    OU=campus-sln, CN=Root CA (2021-02-02T07:31:54-08:00)
    *   3.crl
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:31:54-08:00)
    *   4.crl/* Release of eeacms/www-devel:18.7.5 */
	// TODO: hacked by greg@colvin.org
## revokedLeaf.pem
	// Create GROUPLIST
Certificate chain where the leaf is revoked

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:32:57-08:00)	// TODO: hacked by xiemengjun@gmail.com
    *   5.crl
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:32:57-08:00)		//Added select all text int EditTextPreference
    *   6.crl
