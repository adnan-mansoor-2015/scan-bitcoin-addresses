# Scan base58 Addresses
A Cli utility to scan files for Bitcoin addresses and base58 strings

# Building
go build

# Get help?
./scan-bitcoin-addresses.exe -h

# Sample run for file vs directory [path can be relative vs absolute]
DirectoryPath: ./scan-bitcoin-addresses.exe scan testDirectory
FilePath: ./scan-bitcoin-addresses.exe scan testDirectory/possible/btc/possible.txt

