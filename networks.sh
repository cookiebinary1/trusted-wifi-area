#!/bin/bash

declare -a networks #the array where we will store all saved networks
n=0
for network in /etc/NetworkManager/system-connections/*; do
    networks[$n]="$(basename "$network")"
    (( n++ ))
done

#list all networks in a line
#echo ${networks[*]}

#list networks one by one
for (( i=0; i<${#networks[@]}; i++ )) {
    echo ${networks[$i]}
}

unset networks