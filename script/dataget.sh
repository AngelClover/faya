redis-cli get 000902min2022-09-08 | jq '.Content' | sed 's/\\\"/"/g' | sed 's/^\"//' | sed 's/\"$//' | jq .
