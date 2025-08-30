#!/bin/sh
set -e

echo "PostgreSQL is up - running migrations"
./stellar-x migrate

echo "Starting StellarX"
exec ./stellar-x