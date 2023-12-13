#!/bin/bash

# Array of container names
containers=('weatherapp-auth' 'weatherapp-db' 'weatherapp-redis' 'weatherapp-ui' 'weatherapp-weather')

# Flag to track if any container is not running
containers_not_running=0

for container in "${containers[@]}"
do 
    # Check if the container is running
    RUNNING=$(docker-compose ps | grep "$container" | grep -o Up)

    if [[ ${RUNNING} == Up ]]
    then 
        echo "$container container is in running state"
    else 
        echo "$container container is NOT in running state"
        containers_not_running=1  # Set the flag if any container is not running
    fi
done

# Check the flag to determine the overall status
if [ $containers_not_running -eq 1 ]
then
    exit 1  # Exit with status 1 if any container is not running
else
    exit 0  # Exit with status 0 if all containers are running
fi
