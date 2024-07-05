#!/bin/bash

# JSON payload
json_payload=$(cat <<EOF
{
    "firstName": "Ahmed",
    "lastName": "Aladdin"
}
EOF
)

# Send POST request using curl
curl -X POST http://localhost:8000/accounts \
     -H "Content-Type: application/json" \
     -d "$json_payload"
