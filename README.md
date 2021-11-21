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
./scan-bitcoin-addresses.exe scan testDirectory  
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

