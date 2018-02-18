#!/bin/bash

# Build CLI
echo "Building CLI"
make cli-production
echo

# Upload to server
echo "Uploading new CLI"
scp bin/cli $USER@ec2-35-170-55-23.compute-1.amazonaws.com:/home/$USER/cli-new
echo

# Stop current server
echo "Stopping current CLI"
cat ./scripts/stop.sh | ssh $USER@ec2-35-170-55-23.compute-1.amazonaws.com /bin/bash
echo

# Move new version to the /var/app directory
echo "Replacing old CLI on server"
echo "cp -f cli-new /var/app/coins-cli" | ssh $USER@ec2-35-170-55-23.compute-1.amazonaws.com /bin/bash
echo

# Start up coin trader
echo "Starting CLI on server"
cat ./scripts/start.sh | ssh $USER@ec2-35-170-55-23.compute-1.amazonaws.com /bin/bash
echo