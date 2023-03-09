#!/bin/bash

# Looping every 5 minutes until we stop this script
while true; do

    # We just care about the pods running in default, of which in this case is just our nginx instance
    pods=($(kubectl get pods -n default -o custom-columns=name:.metadata.name --no-headers))

    for pod in "${pods[@]}"; do
        
        # Getting the logs and filtering for errors
        logs=$(kubectl logs -n default "$pod" | grep error)

        echo $logs
    done

    sleep 5m
done
