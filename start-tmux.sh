#!/bin/bash

# Ensure the 'godiscogs' session exists
if ! tmux has-session -t godiscogs 2>/dev/null; then
  # Create a new session named 'godiscogs', detached
  cd /workspaces/godiscogs
  tmux new-session -d -s godiscogs
 fi
