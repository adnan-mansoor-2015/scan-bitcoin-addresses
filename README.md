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

# Sample run for file vs directory [path can be relative vs absolute]
### DirectoryPath:
```console
./scan-bitcoin-addresses.exe scan testDirectory  
```
### FilePath: 
```console
./scan-bitcoin-addresses.exe scan testDirectory/possible/btc/possible.txt
```
