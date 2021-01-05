#!/bin/sh
# Trusted Wi-Fi Area - Do not lock the screen (Ubuntu/Mint with Cinnamon)
#
# git repo: https://github.com/cookiebinary1/trusted-wifi-area
# readme: https://github.com/cookiebinary1/trusted-wifi-area/blob/main/README.md
# author: cookie.binary@gmail.com
#

echo -n "Downloading trusted-wifi-area script to /usr/local/bin. Root permissions are required. Is this ok? (Y/n) "
read yesno </dev/tty

if [ "x$yesno" = "xy" ] || [ "x$yesno" = "x" ]; then
  # Yes
  if ! curl -sS https://raw.githubusercontent.com/cookiebinary1/trusted-wifi-area/main/trusted-wifi-area.sh | sudo tee /usr/local/bin/trusted-wifi-area > /dev/null; then
    echo "Copying failed. Setup terminated."
    exit 1
  fi
  sudo chmod +x /usr/local/bin/trusted-wifi-area
  echo -n "Set automatic start after each login (~/.profile)? (Y/n) "
  read -r yesno </dev/tty

  if [ "x$yesno" = "xy" ] || [ "x$yesno" = "x" ]; then
    # Yes
    sed -i -e '/trusted-wifi-area/d' ~/.profile
    echo "trusted-wifi-area &" >> ~/.profile
    echo "Auto start applied."
  else
    # No
    echo "Autostart was disabled. Start manually with typing: trusted-wifi-area"
  fi

  curl -sS https://raw.githubusercontent.com/cookiebinary1/trusted-wifi-area/main/bin/config | sudo tee /usr/local/bin/trusted-wifi-area-config > /dev/null
  sudo chmod +x /usr/local/bin/trusted-wifi-area-config
  trusted-wifi-area-config
  echo "Setup successful."
else
  # No
  echo "Setup terminated."
fi
