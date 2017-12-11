#!/bin/bash

exec ./out/build//manpower -config ./config/config.json
#exec ./out/build/manpower \
#  -dbPort=${dbPort} \
#  -dbName=${dbName} \
#  -dbServer=${dbServer} \
#  -dbUserName=${dbUserName} \
#  -dbPassword=${dbPassword} \
#  -seedDataPath=${seedDataPath} \
#  -originAllowed=${originAllowed} \
#  $1
