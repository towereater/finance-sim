#!/bin/bash
set -e

## Mainframe library ##
echo "Cleaning lib project"
cd "mainframe-lib"

echo "Security..."
cd "security"
go mod tidy
cd ..

echo "User..."
cd "user"
go mod tidy
cd ..

echo "Account..."
cd "account"
go mod tidy
cd ..

echo "Checking account..."
cd "checking-account"
go mod tidy
cd ..

echo "Dossier..."
cd "dossier"
go mod tidy
cd ..

echo "XChanger..."
cd "xchanger"
go mod tidy
cd ..

cd ..
echo "Done."
echo ""

## Mainframe projects ##
echo "Cleaning main project"
cd "mainframe"

echo "Security..."
cd "security"
go mod tidy
go mod vendor
cd ..

echo "User..."
cd "user"
go mod tidy
go mod vendor
cd ..

echo "Account..."
cd "account"
go mod tidy
go mod vendor
cd ..

echo "Checking account..."
cd "checking-account"
go mod tidy
go mod vendor
cd ..

echo "Dossier..."
cd "dossier"
go mod tidy
go mod vendor
cd ..

cd ..
echo "Done."
echo ""

## BFF ##
echo "Cleaning BFF project"
cd "bff"
go mod tidy
go mod vendor
cd ..
echo "Done."
