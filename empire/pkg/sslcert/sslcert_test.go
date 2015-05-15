package sslcert

import "testing"

func TestPrimaryCertFromChain(t *testing.T) {
	primary := `-----BEGIN CERTIFICATE-----
MIIEdjCCA16gAwIBAgIIavTkiJZO79swDQYJKoZIhvcNAQEFBQAwSTELMAkGA1UE
BhMCVVMxEzARBgNVBAoTCkdvb2dsZSBJbmMxJTAjBgNVBAMTHEdvb2dsZSBJbnRl
cm5ldCBBdXRob3JpdHkgRzIwHhcNMTUwNDIyMTM0NjQyWhcNMTUwNzIxMDAwMDAw
WjBoMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwN
TW91bnRhaW4gVmlldzETMBEGA1UECgwKR29vZ2xlIEluYzEXMBUGA1UEAwwOd3d3
Lmdvb2dsZS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDAB/W1
HE/ZA8NvB/oOrg6WN4zK6vdt9uqTpk3KAQzQ/e2qwy5WoewNmJgx9RPLIuHReTH+
Orh8vF1jk27rC89Pns9rRsE/FVjpaefNHi3trVFbDHcVdrU6JW0XjsG1+wbL8uSU
vDMiey4ZNqUfopX3MJ6LHpFr5FgrRW1RsemTs/BfXjC4MoBdwnt6xoljh+KHzycy
8OgmCVXPONubyUKUeY/Ui9PaX0GWh5dE5uF72jG8NVPs67K7qpfmrdVSGHvRxH3P
AwA90eKlakelqCSfcrZXD8q7EsMBQvRQbbFeuh/TexdiX+8hA1PXNMEURK5y2EBz
5jAbLuuLbQPN/DqZAgMBAAGjggFBMIIBPTAdBgNVHSUEFjAUBggrBgEFBQcDAQYI
KwYBBQUHAwIwGQYDVR0RBBIwEIIOd3d3Lmdvb2dsZS5jb20waAYIKwYBBQUHAQEE
XDBaMCsGCCsGAQUFBzAChh9odHRwOi8vcGtpLmdvb2dsZS5jb20vR0lBRzIuY3J0
MCsGCCsGAQUFBzABhh9odHRwOi8vY2xpZW50czEuZ29vZ2xlLmNvbS9vY3NwMB0G
A1UdDgQWBBSgAQj1VB+R5iA9ZysggEXxg+oRFzAMBgNVHRMBAf8EAjAAMB8GA1Ud
IwQYMBaAFErdBhYbvPZotXb1gba7Yhq6WoEvMBcGA1UdIAQQMA4wDAYKKwYBBAHW
eQIFATAwBgNVHR8EKTAnMCWgI6Ahhh9odHRwOi8vcGtpLmdvb2dsZS5jb20vR0lB
RzIuY3JsMA0GCSqGSIb3DQEBBQUAA4IBAQBlViFXxmddM6J6cVgUBHWw9nsSoJba
SuwHE3Bfo9rrgzV80MazeDaM/tdVSCCA4jD/8jyJr3jl4wtazH9fkn7oBUxYEC22
Xy68txnKMu5LNze+eMbSs7MKSqcPd1ZdUm+3xcsnSc3bmvm/AuOd4WPueMhYdr4c
qwUht+yFSBqEp856JjxsYDlX61h9i7Kq2EAPxA+9GvT2c5j9upUXmUYVnLrx4xjX
LqbbahlsKd+fwvZZ7bFSuyFS8zs5HRfL1kuW1i79cHpuNqQmzC+DrFF2ouJi7uCR
/A6Yfa2DF67oyHZKXuogVwko+cfVe91v+aAQVyltkzAcZy/xaypI8WFj
-----END CERTIFICATE-----
`

	chain := `
-----BEGIN CERTIFICATE-----
MIID8DCCAtigAwIBAgIDAjp2MA0GCSqGSIb3DQEBBQUAMEIxCzAJBgNVBAYTAlVT
MRYwFAYDVQQKEw1HZW9UcnVzdCBJbmMuMRswGQYDVQQDExJHZW9UcnVzdCBHbG9i
YWwgQ0EwHhcNMTMwNDA1MTUxNTU1WhcNMTYxMjMxMjM1OTU5WjBJMQswCQYDVQQG
EwJVUzETMBEGA1UEChMKR29vZ2xlIEluYzElMCMGA1UEAxMcR29vZ2xlIEludGVy
bmV0IEF1dGhvcml0eSBHMjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AJwqBHdc2FCROgajguDYUEi8iT/xGXAaiEZ+4I/F8YnOIe5a/mENtzJEiaB0C1NP
VaTOgmKV7utZX8bhBYASxF6UP7xbSDj0U/ck5vuR6RXEz/RTDfRK/J9U3n2+oGtv
h8DQUB8oMANA2ghzUWx//zo8pzcGjr1LEQTrfSTe5vn8MXH7lNVg8y5Kr0LSy+rE
ahqyzFPdFUuLH8gZYR/Nnag+YyuENWllhMgZxUYi+FOVvuOAShDGKuy6lyARxzmZ
EASg8GF6lSWMTlJ14rbtCMoU/M4iarNOz0YDl5cDfsCx3nuvRTPPuj5xt970JSXC
DTWJnZ37DhF5iR43xa+OcmkCAwEAAaOB5zCB5DAfBgNVHSMEGDAWgBTAephojYn7
qwVkDBF9qn1luMrMTjAdBgNVHQ4EFgQUSt0GFhu89mi1dvWBtrtiGrpagS8wEgYD
VR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAQYwNQYDVR0fBC4wLDAqoCig
JoYkaHR0cDovL2cuc3ltY2IuY29tL2NybHMvZ3RnbG9iYWwuY3JsMC4GCCsGAQUF
BwEBBCIwIDAeBggrBgEFBQcwAYYSaHR0cDovL2cuc3ltY2QuY29tMBcGA1UdIAQQ
MA4wDAYKKwYBBAHWeQIFATANBgkqhkiG9w0BAQUFAAOCAQEAJ4zP6cc7vsBv6JaE
+5xcXZDkd9uLMmCbZdiFJrW6nx7eZE4fxsggWwmfq6ngCTRFomUlNz1/Wm8gzPn6
8R2PEAwCOsTJAXaWvpv5Fdg50cUDR3a4iowx1mDV5I/b+jzG1Zgo+ByPF5E0y8tS
etH7OiDk4Yax2BgPvtaHZI3FCiVCUe+yOLjgHdDh/Ob0r0a678C/xbQF9ZR1DP6i
vgK66oZb+TWzZvXFjYWhGiN3GhkXVBNgnwvhtJwoKvmuAjRtJZOcgqgXe/GFsNMP
WOH7sf6coaPo/ck/9Ndx3L2MpBngISMjVROPpBYCCX65r+7bU2S9cS+5Oc4wt7S8
VOBHBw==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIDfTCCAuagAwIBAgIDErvmMA0GCSqGSIb3DQEBBQUAME4xCzAJBgNVBAYTAlVT
MRAwDgYDVQQKEwdFcXVpZmF4MS0wKwYDVQQLEyRFcXVpZmF4IFNlY3VyZSBDZXJ0
aWZpY2F0ZSBBdXRob3JpdHkwHhcNMDIwNTIxMDQwMDAwWhcNMTgwODIxMDQwMDAw
WjBCMQswCQYDVQQGEwJVUzEWMBQGA1UEChMNR2VvVHJ1c3QgSW5jLjEbMBkGA1UE
AxMSR2VvVHJ1c3QgR2xvYmFsIENBMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
CgKCAQEA2swYYzD99BcjGlZ+W988bDjkcbd4kdS8odhM+KhDtgPpTSEHCIjaWC9m
OSm9BXiLnTjoBbdqfnGk5sRgprDvgOSJKA+eJdbtg/OtppHHmMlCGDUUna2YRpIu
T8rxh0PBFpVXLVDviS2Aelet8u5fa9IAjbkU+BQVNdnARqN7csiRv8lVK83Qlz6c
JmTM386DGXHKTubU1XupGc1V3sjs0l44U+VcT4wt/lAjNvxm5suOpDkZALeVAjmR
Cw7+OC7RHQWa9k0+bw8HHa8sHo9gOeL6NlMTOdReJivbPagUvTLrGAMoUgRx5asz
PeE4uwc2hGKceeoWMPRfwCvocWvk+QIDAQABo4HwMIHtMB8GA1UdIwQYMBaAFEjm
aPkr0rKV10fYIyAQTzOYkJ/UMB0GA1UdDgQWBBTAephojYn7qwVkDBF9qn1luMrM
TjAPBgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBBjA6BgNVHR8EMzAxMC+g
LaArhilodHRwOi8vY3JsLmdlb3RydXN0LmNvbS9jcmxzL3NlY3VyZWNhLmNybDBO
BgNVHSAERzBFMEMGBFUdIAAwOzA5BggrBgEFBQcCARYtaHR0cHM6Ly93d3cuZ2Vv
dHJ1c3QuY29tL3Jlc291cmNlcy9yZXBvc2l0b3J5MA0GCSqGSIb3DQEBBQUAA4GB
AHbhEm5OSxYShjAGsoEIz/AIx8dxfmbuwu3UOx//8PDITtZDOLC5MH0Y0FWDomrL
NhGc6Ehmo21/uBPUR/6LWlxz/K7ZGzIZOKuXNBSqltLroxwUCEm2u+WR74M26x1W
b8ravHNjkOR/ez4iyz0H7V84dJzjA1BOoa+Y7mHyhD8S
-----END CERTIFICATE-----
`

	pri, rest := SplitCertChain(primary + chain)
	if pri != primary {
		t.Errorf("SplitCertChain() => %v; want %v", pri, primary)
	}
	if rest != chain {
		t.Errorf("SplitCertChain() => %v; want %v", rest, chain)
	}
}
