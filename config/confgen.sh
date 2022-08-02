#!/bin/bash

confs=(
    '%%GRAPH_RPC%%=0.0.0.0:6070'
    '%%JUDGE_RPC%%=0.0.0.0:6080'
    '%%EV_RPC%%=0.0.0.0:8433'
    '%%REDIS%%=127.0.0.1:6379'
 )

configurer() {
    for i in "${confs[@]}"
    do
        search="${i%%=*}"
        replace="${i##*=}"

        uname=`uname`
        if [ "$uname" == "Darwin" ] ; then
            # Note the "" and -e  after -i, needed in OS X
            find ./config/*.json -type f -exec sed -i .tpl -e "s/${search}/${replace}/g" {} \;
        else
            find ./config/*.json -type f -exec sed -i "s/${search}/${replace}/g" {} \;
        fi
    done
}
configurer
