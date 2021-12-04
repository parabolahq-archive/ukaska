#!/bin/bash
echo -e "version: 1.4\n\njobs:\n     Backup:\n        cmd: /app/main\n        time: ${JOBBER_STRING}" > /root/.jobber
set -m
mongod --bind_ip_all  &

/jobber/jobberrunner -u /var/jobber/0/cmd.sock /root/.jobber

fg %1