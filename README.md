# Trusted Wi-Fi Area - Do not lock the screen (Ubuntu/Mint with Cinnamon)

If you don't want to lock the screen saver on a specific Wi-Fi network (at home, etc.), try this solution.

### How it works?
The screensaver will not lock the screen on the listed Wi-Fi networks.
So far, the implementation is only for Cinnamon (Ubuntu, Linux Mint etc.).
Other desktop environments are being worked on.

## Download 

Download and unzip by:

```bash
curl -sS https://codeload.github.com/cookiebinary1/trusted-wifi-area/zip/main > trusted-wifi-area.zip 
unzip trusted-wifi-area.zip 
rm trusted-wifi-area.zip
cd trusted-wifi-area-main
```

## Installation (automatically)

Notice: if you want to run script manually, please skip this step.

Start setup by entering:

```bash
sh setup.sh
```

Follow the instructions.

## Alternative Installation (manually)

Here are three simple steps. Copy script, set executable permissions and auto startup.

```bash
sudo cp trusted-wifi-area.sh /usr/local/bin
sudo chmod +x /usr/local/bin/trusted-wifi-area
echo "trusted-wifi-area.sh &" >> ~/.profile
```

## Configuration

Create/Edit config file with command bellow.
```bash
nano ~/.trusted-wifi-area
```
..and enter list of the Wi-Fi names.
Press `Ctrl`+`O` to save and `Ctrl`+`X` to exit.

#### Example
```
mywifi01
dlink1234567
HUAWEI-f3rDe
```

## Start
After the restart, the script will run automatically.
To run it manually, enter this:
```bash
sh trusted-wifi-area.sh
```

## Uninstallation

Remove script itself and its call from `.profile`:
```bash
sudo rm /usr/local/bin/trusted-wifi-area
sed -i -e '/trusted-wifi-area/d' ~/.profile
```

## Contributing
Pull requests are welcome. If anybody has any ideas, you are welcome.

### Warning
Use at your own risk.

## License
[MIT](https://choosealicense.com/licenses/mit/)