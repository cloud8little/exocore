#!/bin/bash
rm -rf ~/.tmp-exocored
echo "Setup single validator node"
source ./local_node.sh
EXOCORE_PID=$!
echo "started exocore node, pid=$EXOCORE_PID"

ORACLE_START_BLOCK="20"

# Wait for the process to start
echo "Waiting for Exocore process to start..."
sleep 60  # Adjust the sleep time as needed

# check current height,if less than 20, wait for 3 seconds
echo "Checking Exocore process status..."

get_block_height() {
    curl -s http://localhost:26657/block | jq ".result.block.header.height"
}

RETRY_COUNT=5

for i in $(seq 1 $RETRY_COUNT); do
    CURRENT_HEIGHT=$(get_block_height)
    echo $CURRENT_HEIGHT
    CURRENT_HEIGHT=$(echo "$CURRENT_HEIGHT" | sed 's/"//g')
    echo "Current block height: $CURRENT_HEIGHT"

    if [ "$CURRENT_HEIGHT" -gt "$ORACLE_START_BLOCK" ]; then
        echo "Target block height $ORACLE_START_BLOCK reached."
        break
    fi

    # Wait for a few seconds before checking again
    sleep 1
done

# if the current height is still less than the target block height, exit with an error
if [ "$CURRENT_HEIGHT" -lt "$ORACLE_START_BLOCK" ]; then
    echo "Error: Current block height is less than the target block height."
    
    # Kill the process if needed
    if [ -n "$EXOCORE_PID" ]; then
        echo "Killing process with PID $EXOCORE_PID"
        kill $EXOCORE_PID
    fi
    
    # Exit with an error code
    exit 1
fi

echo "Checking oracle price..."
DEFAULT_PRICE=$(curl -X GET "http://localhost:1317/ExocoreNetwork/exocore/oracle/latest_price/1" -H "accept: application/json" | jq ".price.price")

if [ -z "$DEFAULT_PRICE" ]; then
    echo "Error: DEFAULT_PRICE is empty."
    
    # Kill the process if needed
    if [ -n "$EXOCORE_PID" ]; then
        echo "Killing process with PID $EXOCORE_PID"
        kill $EXOCORE_PID
    fi
    
    # Exit with an error code
    exit 1
fi

echo "Oracle price is $DEFAULT_PRICE"
echo "Sanity testing passed"
kill $EXOCORE_PID
pgrep exocored | xargs kill -9