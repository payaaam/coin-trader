#!/bin/bash

HOST=ec2-35-170-55-23.compute-1.amazonaws.com

# Use $USER val if the DEPLOY_USER variable is unset
if [[ -z "${DEPLOY_USER:-}" ]]; then
  DEPLOY_USER=$USER
fi

# Ensure user has access
echo "Verifying user access"
status=$(ssh $DEPLOY_USER@${HOST} echo ok)
if [[ "$status" != "ok" ]]; then
  echo "ERROR: User $DEPLOY_USER does not have access"
  exit 1
fi
echo

# Ensure tests pass
echo "Running Tests"
make test
if [[ "$?" -ne 0 ]]; then
  echo "Warning: Tests do not pass..."
  exit 1
fi
echo

# Build CLI
echo "Building CLI"
make cli-production
echo

# Upload to server
echo "Uploading new CLI"
scp bin/cli $DEPLOY_USER@${HOST}:/home/$DEPLOY_USER/cli-new
echo

# Stop current server
echo "Stopping current CLI"
cat ./scripts/stop.sh | ssh $DEPLOY_USER@${HOST} /bin/bash
echo

# Move new version to the /var/app directory
echo "Replacing old CLI on server"
echo "cp -f cli-new /var/app/coins-cli" | ssh $DEPLOY_USER@${HOST} /bin/bash
echo

# Start up coin trader
echo "Starting CLI on server"
cat ./scripts/start.sh | ssh $DEPLOY_USER@${HOST} /bin/bash
echo