#!/usr/bin/env bash
# Loads sample data into a running stack through the public HTTP APIs.
# Works with the Go, Node and Java implementations alike.
#
#   PAYMENTS_URL   default http://localhost:3001
#   ANALYTICS_URL  default http://localhost:3002
#   TX_COUNT       number of merchant transactions to create, default 3000
set -euo pipefail

PAYMENTS_URL="${PAYMENTS_URL:-http://localhost:3001}"
ANALYTICS_URL="${ANALYTICS_URL:-http://localhost:3002}"
TX_COUNT="${TX_COUNT:-3000}"

echo "Creating accounts on $PAYMENTS_URL"
curl -s -X POST "$PAYMENTS_URL/accounts" -H 'Content-Type: application/json' \
  -d '{"id":"acc-alice","ownerName":"Alice","balance":100,"merchantId":""}'; echo
curl -s -X POST "$PAYMENTS_URL/accounts" -H 'Content-Type: application/json' \
  -d '{"id":"acc-bob","ownerName":"Bob","balance":50,"merchantId":""}'; echo
curl -s -X POST "$PAYMENTS_URL/accounts" -H 'Content-Type: application/json' \
  -d '{"id":"acc-merchant","ownerName":"Demo Merchant","balance":0,"merchantId":"merch-1"}'; echo

TODAY="$(date -u +%F)"
echo "Creating $TX_COUNT transactions for merch-1 on $TODAY via $ANALYTICS_URL"
for i in $(seq 1 "$TX_COUNT"); do
  case $((i % 3)) in
    0) STATUS=captured ;;
    1) STATUS=failed ;;
    2) STATUS=refunded ;;
  esac
  curl -s -o /dev/null -X POST "$ANALYTICS_URL/transactions" -H 'Content-Type: application/json' \
    -d "{\"merchantId\":\"merch-1\",\"amount\":25,\"status\":\"$STATUS\",\"createdAt\":\"${TODAY}T12:00:00Z\"}"
  if (( i % 500 == 0 )); then echo "  $i / $TX_COUNT"; fi
done

echo "Done. Try:"
echo "  curl $PAYMENTS_URL/accounts/acc-alice"
echo "  curl '$ANALYTICS_URL/merchants/merch-1/daily-totals?date=$TODAY'"
