# Trusted Wi-Fi Area - Do not lock the screen (Ubuntu/Mint with Cinnamon)

If you don't want to lock the screen saver on a specific Wi-Fi network (at home, etc.), try this solution.

Test

### How it works?
The screensaver will not lock the screen on the listed Wi-Fi networks.
So far, the implementation is only for Cinnamon (Ubuntu, Linux Mint etc.).
Other desktop environments are being worked on.

## Installation (automatically)

Enter following command in terminal:

```bash
curl -sS https://raw.githubusercontent.com/cookiebinary1/trusted-wifi-area/main/setup.sh | sh
```

Setup will start. Follow the instructions.

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
After the logout / login (restart) the script will run automatically.

To run it manually enter this:
```bash
trusted-wifi-area
```

## Uninstallation

Remove script itself and its call from `.profile`:
```bash
sudo rm /usr/local/bin/trusted-wifi-area
sed -i -e '/trusted-wifi-area/d' ~/.profile
```

## Alternative Installation (manually)

Here are several simple steps. Download the script to target folder, set permission to execute and setup auto start.

```bash
curl -sS https://raw.githubusercontent.com/cookiebinary1/trusted-wifi-area/main/trusted-wifi-area.sh | sudo tee /usr/local/bin/trusted-wifi-area > /dev/null
sudo chmod +x /usr/local/bin/trusted-wifi-area
echo "trusted-wifi-area &" >> ~/.profile
```

## Contributing
Pull requests are welcome. If anybody has any ideas, you are welcome.

### Warning
Use at your own risk.

## License
[MIT](https://choosealicense.com/licenses/mit/)