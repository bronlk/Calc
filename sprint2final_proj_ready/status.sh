#!/bin/bash


while true; do

    clear
  date 
  echo list_agents
  curl -X POST 192.168.1.11:8090/list_agents
  echo
  echo -------------------------------
  echo list_expressions
  curl -X POST 192.168.1.11:8090/list_expressions
    
 

  sleep 2
done
