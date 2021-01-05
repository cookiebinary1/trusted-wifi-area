#!/bin/sh

current_essid=""
config_file=~/.trusted-wifi-area
while :
do
  new_essid=$(iwgetid -r)
  if [ "$new_essid" = "" ]
  then
  	new_essid="no-wifi"
  fi

  if [ "$new_essid" != "$current_essid" ]
  then
  	  current_essid="$new_essid"
  	  echo "New wifi: $new_essid" 
  	  killall cinnamon-screensaver
	  sleep 5s

	  if grep -q "$new_essid" "$config_file"; then
	    cinnamon-screensaver --disable-locking & echo "Restart screensaver without locking..."
	  else
	    cinnamon-screensaver & echo "Restart screensaver with locking..."
	  fi
  fi
  sleep 1s
done
