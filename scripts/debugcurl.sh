#!/bin/bash
curl -s -X POST "https://notely-db-rohansriv-110.aws-ap-south-1.turso.io/v2/pipeline" -H "Authorization: Bearer ${DATABASE_URL#*authToken=}" -H "Content-Type: application/json" -d '{"requests":[{"type":"execute","stmt":{"sql":"select 1"}},{"type":"close"}]}'
