# Trusted Wi-Fi Area - Not to Lock Screen (Ubuntu/Mint with Cinnamon)

If you don't want to lock the screen saver on a specific Wi-Fi network (at home, etc.), try this solution.

### How it works?
The screensaver will not lock the screen on the listed Wi-Fi networks.
So far, the implementation is only for Cinnamon (Ubuntu, Linux Mint etc.).
Other desktop environments are being worked on.

## Installation
Notice: if you want to run script manually, please skip entire installation step.

Here are three simple steps. Copy script, set executable permissions and auto startup.

```bash
sudo cp trusted-wifi-area.sh /usr/local/bin
sudo chmod +x /usr/local/bin/trusted-wifi-area.sh
echo "trusted-wifi-area.sh &" >> ~/.profile
```

## Configuration

Create/Edit config file with command bellow.
```bash
nano ~/.trusted-wifi-area
```
..and enter list of the Wi-Fi names.
Press `Ctrl`+`O` to save and `Ctrl`+`X` to exit.

### Example
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

## Contributing
Pull requests are welcome. If anybody has any ideas, you are welcome.

### Warning
Use at your own risk.

## License
[MIT](https://choosealicense.com/licenses/mit/)