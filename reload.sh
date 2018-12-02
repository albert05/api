#! /bin/bash

# update code
git checkout .
git pull origin master

# copy config
rm -rf conf/app.conf
cp conf/app.conf.bat conf/app.conf

# stop fs
supervisorctl stop api

# get listen port
pid=`lsof -i:9999 | grep LISTEN | awk -F '[ ]+' '{print $2}'`

# kill port listen
kill ${pid}

# start fs
supervisorctl start api

# re grant authorization
chmod -R 755 reload.sh
chmod -R 755 restart.sh
