# Scan base58 Addresses
A Cli utility to scan files for Bitcoin addresses and base58 strings

# Building
```console
cd scan-bitcoin-addresses  
go build
```

# Get help?
```console
cd scan-bitcoin-addresses  
./scan-bitcoin-addresses.exe -h
```

# Sample run for file vs directory [path can be relative or absolute]
### DirectoryPath:
```console
./scan-bitcoin-addresses.exe scan directory1 directory2 directory3  
```
### FilePath: 
```console
./scan-bitcoin-addresses.exe scan testDirectory/possible/btc/possible.txt
```

# Sample output
### Perfect match
```console
{
    "MatchType": "perfect",
    "MatchFile": "testDirectory\\perfect\\btc\\perfect.txt",
    "MatchedLine": "1Q1pE5vPGEEMqRcVRMbtBK842Y6Pzo6nK9 some words after the address",
    "MatchedWord": "1Q1pE5vPGEEMqRcVRMbtBK842Y6Pzo6nK9"
}
```

### Possible match
```console
{
    "MatchType": "possible",
    "MatchFile": "testDirectory\\possible\\btc\\possible.txt",
    "MatchedLine": "btc:1NKRhS7iYUGTaAfaR5z8BueAJesqaTyc4a btc:19ck9VKC6KjGxR9LJg4DNMRc45qFrJguvV address then possible address then some words case:2",
    "MatchedWord": "btc:1NKRhS7iYUGTaAfaR5z8BueAJesqaTyc4a"
}
```

## Verbose mode

```console
go run main.go scan -v testDirectory


Scanning: Dir(1 of 1) File(1 of 6): testDirectory\possible\btc\possible.txt
{
    "MatchType": "possible",
    "MatchFile": "testDirectory\\possible\\btc\\possible.txt",
    "MatchedLine": "btc:1Q1pE5vPGEEMqRcVRMbtBK842Y6Pzo6nK9 btc:1Q1pE5vPGEEMqRcVRMbtBK842Y6Pzo6nK9 address then possible address then some words case:2",
    "MatchedLineNo": 1,
    "MatchedWord": "btc:1Q1pE5vPGEEMqRcVRMbtBK842Y6Pzo6nK9"
}
...

Scan Complete: Dir(1 of 1) File(1 of 6): testDirectory\possible\btc\possible.txt, size: 1 KB(s), line(s): 18, timeTaken: 65 ms

---

Scanning: Dir(1 of 1) File(6 of 6): testDirectory\perfect\skycoin\perfect.txt
{
    "MatchType": "perfect",
    "MatchFile": "testDirectory\\perfect\\skycoin\\perfect.txt",
    "MatchedLine": "1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq some words after the address",
    "MatchedLineNo": 1,
    "MatchedWord": "1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq"
}
...

Scan Complete: Dir(1 of 1) File(6 of 6): testDirectory\perfect\skycoin\perfect.txt, size: 529 Byte(s), line(s): 9, timeTaken: 21 ms```
