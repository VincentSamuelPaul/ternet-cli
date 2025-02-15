#!/bin/bash

# URL to send the request to
URL="http://localhost:3000/createpost"

# Execute the curl request
curl --header "Content-Type: application/json" \
--request GET \
http://localhost:3000/getposts
# --data '{"username":"samuel","data":"hey niggas"}' \
# http://localhost:3000/createuser
# http://localhost:3000/loginuser