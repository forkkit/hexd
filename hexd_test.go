package hexd

import (
	"encoding/base64"
	"testing"
)

func TestWrite(t *testing.T) {
	data, _ := base64.StdEncoding.DecodeString(dataRaw)
	dump, _ := base64.StdEncoding.DecodeString(dumpRaw)
	IncludeHeader = true
	if string(Dump(data)) != string(dump) {
		t.Fatal("mismatch")
	}
	if DumpString(data) != string(dump) {
		t.Fatal("mismatch")
	}
}

const dataRaw = `
U1FMaXRlIGZvcm1hdCAzAAQAAQEAQCAgAAAAGQAAA2AAAAAAAAAAAAAAACIAAAABAAAAAAAAAAAAAAA
BAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAZAC3iHgUAAAAHA90AAAAAGQP7A/YD8Q
PsA+cD4gPdAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAhGIEBxcfHwGJE3RhYmxlY
3VzdG9tZXJzY3VzdG9tZXJzBUNSRUFURSBUQUJMRSAiY3VzdG9tZXJzIg0KKA0KICAgIFtDdXN0b21l
cklkXSBJTlRFR0VSIFBSSU1BUlkgS0VZIEFVVE9JTkNSRU1FTlQgTk9UIE5VTEwsDQogICAgW0ZpcnN
0TmFtZV0gTlZBUkNIQVIoNDApICBOT1QgTlVMTCwNCiAgICBbTGFzdE5hbWVdIE5WQVJDSEFSKDIwKS
AgTk9UIE5VTEwsDQogICAgW0NvbXBhbnldIE5WQVJDSEFSKDgwKSwNCiAgICBbQWRkcmVzc10gTlZBU
kNIQVIoNzApLA0KICAgIFtDaXR5XSBOVkFSQ0hBUig0MCksDQogICAgW1N0YXRlXSBOVkFSQ0hBUig0
MCksDQogICAgW0NvdW50cnldIE5WQVJDSEFSKDQwKSwNCiAgICBbUG9zdGFsQ29kZV0gTlZBUkNIQVI
oMTApLA0KICAgIFtQaG9uZV0gTlZBUkNIQVIoMjQpLA0KICAgIFtGYXhdIE5WQVJDSEFSKDI0KSwNCi
AgICBbRW1haWxdIE5WQVJDSEFSKDYwKSAgTk9UIE5VTEwsDQogICAgW1N1cHBvcnRSZXBJZF0gSU5UR
UdFUiwNCiAgICBGT1JFSUdOIEtFWSAoW1N1cHBvcnRSZXBJZF0pIFJFRkVSRU5DRVMgImVtcGxveWVl
cyIgKFtFbXBsb3llZUlkXSkgDQoJCU9OIERFTEVURSBOTyBBQ1RJT04gT04gVVBEQVRFIE5PIEFDVEl
PTg0KKQUAAAABA/sAAAABlgP7AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAA
`

const dumpRaw = `
T2Zmc2V0IChoKSAgMDAgMDEgMDIgMDMgMDQgMDUgMDYgMDcgMDggMDkgMEEgMEIgMEMgMEQgMEUgMEY
KMDAwMDAwMDAwMCAgNTMgNTEgNEMgNjkgNzQgNjUgMjAgNjYgNkYgNzIgNkQgNjEgNzQgMjAgMzMgMD
AgIFNRTGl0ZSBmb3JtYXQgMy4KMDAwMDAwMDAxMCAgMDQgMDAgMDEgMDEgMDAgNDAgMjAgMjAgMDAgM
DAgMDAgMTkgMDAgMDAgMDMgNjAgIC4uLi4uQCAgLi4uLi4uLmAKMDAwMDAwMDAyMCAgMDAgMDAgMDAg
MDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMjIgMDAgMDAgMDAgMDEgIC4uLi4uLi4uLi4uIi4uLi4KMDA
wMDAwMDAzMCAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDEgMDAgMDAgMDAgMDAgIC
4uLi4uLi4uLi4uLi4uLi4KMDAwMDAwMDA0MCAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgM
DAgMDAgMDAgMDAgMDAgMDAgIC4uLi4uLi4uLi4uLi4uLi4KMDAwMDAwMDA1MCAgMDAgMDAgMDAgMDAg
MDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMTkgIC4uLi4uLi4uLi4uLi4uLi4KMDAwMDA
wMDA2MCAgMDAgMkQgRTIgMUUgMDUgMDAgMDAgMDAgMDcgMDMgREQgMDAgMDAgMDAgMDAgMTkgIC4tw6
IuLi4uLi4uw50uLi4uLgowMDAwMDAwMDcwICAwMyBGQiAwMyBGNiAwMyBGMSAwMyBFQyAwMyBFNyAwM
yBFMiAwMyBERCAwMCAwMCAgLsO7LsO2LsOxLsOsLsOnLsOiLsOdLi4KMDAwMDAwMDA4MCAgMDAgMDAg
MDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgIC4uLi4uLi4uLi4uLi4uLi4
KMDAwMDAwMDA5MCAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMD
AgIC4uLi4uLi4uLi4uLi4uLi4KMDAwMDAwMDBhMCAgMDAgMDAgODQgNjIgMDQgMDcgMTcgMUYgMUYgM
DEgODkgMTMgNzQgNjEgNjIgNkMgIC4uwoRiLi4uLi4uwokudGFibAowMDAwMDAwMGIwICA2NSA2MyA3
NSA3MyA3NCA2RiA2RCA2NSA3MiA3MyA2MyA3NSA3MyA3NCA2RiA2RCAgZWN1c3RvbWVyc2N1c3RvbQo
wMDAwMDAwMGMwICA2NSA3MiA3MyAwNSA0MyA1MiA0NSA0MSA1NCA0NSAyMCA1NCA0MSA0MiA0QyA0NS
AgZXJzLkNSRUFURSBUQUJMRQowMDAwMDAwMGQwICAyMCAyMiA2MyA3NSA3MyA3NCA2RiA2RCA2NSA3M
iA3MyAyMiAwRCAwQSAyOCAwRCAgICJjdXN0b21lcnMiLi4oLgowMDAwMDAwMGUwICAwQSAyMCAyMCAy
MCAyMCA1QiA0MyA3NSA3MyA3NCA2RiA2RCA2NSA3MiA0OSA2NCAgLiAgICBbQ3VzdG9tZXJJZAowMDA
wMDAwMGYwICA1RCAyMCA0OSA0RSA1NCA0NSA0NyA0NSA1MiAyMCA1MCA1MiA0OSA0RCA0MSA1MiAgXS
BJTlRFR0VSIFBSSU1BUgowMDAwMDAwMTAwICA1OSAyMCA0QiA0NSA1OSAyMCA0MSA1NSA1NCA0RiA0O
SA0RSA0MyA1MiA0NSA0RCAgWSBLRVkgQVVUT0lOQ1JFTQowMDAwMDAwMTEwICA0NSA0RSA1NCAyMCA0
RSA0RiA1NCAyMCA0RSA1NSA0QyA0QyAyQyAwRCAwQSAyMCAgRU5UIE5PVCBOVUxMLC4uIAowMDAwMDA
wMTIwICAyMCAyMCAyMCA1QiA0NiA2OSA3MiA3MyA3NCA0RSA2MSA2RCA2NSA1RCAyMCA0RSAgICAgW0
ZpcnN0TmFtZV0gTgowMDAwMDAwMTMwICA1NiA0MSA1MiA0MyA0OCA0MSA1MiAyOCAzNCAzMCAyOSAyM
CAyMCA0RSA0RiA1NCAgVkFSQ0hBUig0MCkgIE5PVAowMDAwMDAwMTQwICAyMCA0RSA1NSA0QyA0QyAy
QyAwRCAwQSAyMCAyMCAyMCAyMCA1QiA0QyA2MSA3MyAgIE5VTEwsLi4gICAgW0xhcwowMDAwMDAwMTU
wICA3NCA0RSA2MSA2RCA2NSA1RCAyMCA0RSA1NiA0MSA1MiA0MyA0OCA0MSA1MiAyOCAgdE5hbWVdIE
5WQVJDSEFSKAowMDAwMDAwMTYwICAzMiAzMCAyOSAyMCAyMCA0RSA0RiA1NCAyMCA0RSA1NSA0QyA0Q
yAyQyAwRCAwQSAgMjApICBOT1QgTlVMTCwuLgowMDAwMDAwMTcwICAyMCAyMCAyMCAyMCA1QiA0MyA2
RiA2RCA3MCA2MSA2RSA3OSA1RCAyMCA0RSA1NiAgICAgIFtDb21wYW55XSBOVgowMDAwMDAwMTgwICA
0MSA1MiA0MyA0OCA0MSA1MiAyOCAzOCAzMCAyOSAyQyAwRCAwQSAyMCAyMCAyMCAgQVJDSEFSKDgwKS
wuLiAgIAowMDAwMDAwMTkwICAyMCA1QiA0MSA2NCA2NCA3MiA2NSA3MyA3MyA1RCAyMCA0RSA1NiA0M
SA1MiA0MyAgIFtBZGRyZXNzXSBOVkFSQwowMDAwMDAwMWEwICA0OCA0MSA1MiAyOCAzNyAzMCAyOSAy
QyAwRCAwQSAyMCAyMCAyMCAyMCA1QiA0MyAgSEFSKDcwKSwuLiAgICBbQwowMDAwMDAwMWIwICA2OSA
3NCA3OSA1RCAyMCA0RSA1NiA0MSA1MiA0MyA0OCA0MSA1MiAyOCAzNCAzMCAgaXR5XSBOVkFSQ0hBUi
g0MAowMDAwMDAwMWMwICAyOSAyQyAwRCAwQSAyMCAyMCAyMCAyMCA1QiA1MyA3NCA2MSA3NCA2NSA1R
CAyMCAgKSwuLiAgICBbU3RhdGVdIAowMDAwMDAwMWQwICA0RSA1NiA0MSA1MiA0MyA0OCA0MSA1MiAy
OCAzNCAzMCAyOSAyQyAwRCAwQSAyMCAgTlZBUkNIQVIoNDApLC4uIAowMDAwMDAwMWUwICAyMCAyMCA
yMCA1QiA0MyA2RiA3NSA2RSA3NCA3MiA3OSA1RCAyMCA0RSA1NiA0MSAgICAgW0NvdW50cnldIE5WQQ
owMDAwMDAwMWYwICA1MiA0MyA0OCA0MSA1MiAyOCAzNCAzMCAyOSAyQyAwRCAwQSAyMCAyMCAyMCAyM
CAgUkNIQVIoNDApLC4uICAgIAowMDAwMDAwMjAwICA1QiA1MCA2RiA3MyA3NCA2MSA2QyA0MyA2RiA2
NCA2NSA1RCAyMCA0RSA1NiA0MSAgW1Bvc3RhbENvZGVdIE5WQQowMDAwMDAwMjEwICA1MiA0MyA0OCA
0MSA1MiAyOCAzMSAzMCAyOSAyQyAwRCAwQSAyMCAyMCAyMCAyMCAgUkNIQVIoMTApLC4uICAgIAowMD
AwMDAwMjIwICA1QiA1MCA2OCA2RiA2RSA2NSA1RCAyMCA0RSA1NiA0MSA1MiA0MyA0OCA0MSA1MiAgW
1Bob25lXSBOVkFSQ0hBUgowMDAwMDAwMjMwICAyOCAzMiAzNCAyOSAyQyAwRCAwQSAyMCAyMCAyMCAy
MCA1QiA0NiA2MSA3OCA1RCAgKDI0KSwuLiAgICBbRmF4XQowMDAwMDAwMjQwICAyMCA0RSA1NiA0MSA
1MiA0MyA0OCA0MSA1MiAyOCAzMiAzNCAyOSAyQyAwRCAwQSAgIE5WQVJDSEFSKDI0KSwuLgowMDAwMD
AwMjUwICAyMCAyMCAyMCAyMCA1QiA0NSA2RCA2MSA2OSA2QyA1RCAyMCA0RSA1NiA0MSA1MiAgICAgI
FtFbWFpbF0gTlZBUgowMDAwMDAwMjYwICA0MyA0OCA0MSA1MiAyOCAzNiAzMCAyOSAyMCAyMCA0RSA0
RiA1NCAyMCA0RSA1NSAgQ0hBUig2MCkgIE5PVCBOVQowMDAwMDAwMjcwICA0QyA0QyAyQyAwRCAwQSA
yMCAyMCAyMCAyMCA1QiA1MyA3NSA3MCA3MCA2RiA3MiAgTEwsLi4gICAgW1N1cHBvcgowMDAwMDAwMj
gwICA3NCA1MiA2NSA3MCA0OSA2NCA1RCAyMCA0OSA0RSA1NCA0NSA0NyA0NSA1MiAyQyAgdFJlcElkX
SBJTlRFR0VSLAowMDAwMDAwMjkwICAwRCAwQSAyMCAyMCAyMCAyMCA0NiA0RiA1MiA0NSA0OSA0NyA0
RSAyMCA0QiA0NSAgLi4gICAgRk9SRUlHTiBLRQowMDAwMDAwMmEwICA1OSAyMCAyOCA1QiA1MyA3NSA
3MCA3MCA2RiA3MiA3NCA1MiA2NSA3MCA0OSA2NCAgWSAoW1N1cHBvcnRSZXBJZAowMDAwMDAwMmIwIC
A1RCAyOSAyMCA1MiA0NSA0NiA0NSA1MiA0NSA0RSA0MyA0NSA1MyAyMCAyMiA2NSAgXSkgUkVGRVJFT
kNFUyAiZQowMDAwMDAwMmMwICA2RCA3MCA2QyA2RiA3OSA2NSA2NSA3MyAyMiAyMCAyOCA1QiA0NSA2
RCA3MCA2QyAgbXBsb3llZXMiIChbRW1wbAowMDAwMDAwMmQwICA2RiA3OSA2NSA2NSA0OSA2NCA1RCA
yOSAyMCAwRCAwQSAwOSAwOSA0RiA0RSAyMCAgb3llZUlkXSkgLi4uLk9OIAowMDAwMDAwMmUwICA0NC
A0NSA0QyA0NSA1NCA0NSAyMCA0RSA0RiAyMCA0MSA0MyA1NCA0OSA0RiA0RSAgREVMRVRFIE5PIEFDV
ElPTgowMDAwMDAwMmYwICAyMCA0RiA0RSAyMCA1NSA1MCA0NCA0MSA1NCA0NSAyMCA0RSA0RiAyMCA0
MSA0MyAgIE9OIFVQREFURSBOTyBBQwowMDAwMDAwMzAwICA1NCA0OSA0RiA0RSAwRCAwQSAyOSAwNSA
wMCAwMCAwMCAwMSAwMyBGQiAwMCAwMCAgVElPTi4uKS4uLi4uLsO7Li4KMDAwMDAwMDMxMCAgMDAgMD
EgOTYgMDMgRkIgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgMDAgIC4uwpYuw7suLi4uLi4uL
i4uLgowMDAwMDAwMzIwICAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAw
MCAwMCAgLi4uLi4uLi4uLi4uLi4uLgowMDAwMDAwMzMwICAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCA
wMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAgLi4uLi4uLi4uLi4uLi4uLgowMDAwMDAwMzQwICAwMCAwMC
AwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAgLi4uLi4uLi4uLi4uLi4uL
gowMDAwMDAwMzUwICAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAw
MCAgLi4uLi4uLi4uLi4uLi4uLgowMDAwMDAwMzYwICAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCA
wMCAwMCAwMCAwMCAwMCAwMCAwMCAgLi4uLi4uLi4uLi4uLi4uLgowMDAwMDAwMzcwICAwMCAwMCAwMC
AwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAgLi4uLi4uLi4uLi4uLi4uLgowM
DAwMDAwMzgwICAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAg
Li4uLi4uLi4uLi4uLi4uLgowMDAwMDAwMzkwICAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCA
wMCAwMCAwMCAwMCAwMCAwMCAgLi4uLi4uLi4uLi4uLi4uLgowMDAwMDAwM2EwICAwMCAwMCAwMCAwMC
AwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAgLi4uLi4uLi4uLi4uLi4uLgowMDAwM
DAwM2IwICAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAgLi4u
Li4uLi4uLi4uLi4uLgowMDAwMDAwM2MwICAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCA
wMCAwMCAwMCAwMCAwMCAgLi4uLi4uLi4uLi4uLi4uLgowMDAwMDAwM2QwICAwMCAwMCAwMCAwMCAwMC
AwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAgLi4uLi4uLi4uLi4uLi4uLgowMDAwMDAwM
2UwICAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAgLi4uLi4u
Li4uLi4uLi4uLgowMDAwMDAwM2YwICAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCAwMCA
gICAgICAgICAgICAgLi4uLi4uLi4uLi4uICAgIAo=
`
