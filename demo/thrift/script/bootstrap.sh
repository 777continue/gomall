#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/thrift"
exec "$CURDIR/bin/thrift"