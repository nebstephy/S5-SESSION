#!/bin/bash
ip_address=$(curl -s https://api.ipify.org ; echo)
ip_port=3000
url="http://$ip_address:$ip_port/login"
expected_title='<title>Weather app login</title>'

response=$(curl -s -o /dev/null -w "%{http_code}" "$url")  # Check response code

if [[ $response -ne 200 ]]; then
    echo "The website is not responding (HTTP status code: $response), please check"
    exit 1
fi

title_check=$(curl -s "$url" | grep -c "$expected_title")

if [[ $title_check -eq 0 ]]; then
    echo "The expected title '$expected_title' was not found. The website might be down or changed its structure."
    exit 1
else
    echo "The website is up and the expected title '$expected_title' was found."
    exit 0
fi
