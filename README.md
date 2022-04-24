# CS:GO Server Scanner

It scan your favorite CS:GO community servers and send desktop notification when your favorite map is playing.

## Works only on Windows OS.

On action:


![Alt text](ss.png?raw=true "SS")

______________
## How to Install

``` sh
go build -o csgo-server-scanner.exe
./csgo-server-scanner.exe
```

Before running, place config.json with your binary path(csgo-server-scanner.exe). Example config at config.json

``` JSON
{
    "favoriteMap": "de_dust2",
    "serverList": [
        {
            "name": "Turkiye Cumhuriyeti",
            "host": "185.193.165.212:27015"
        },
        {
            "name": "La Casa",
            "host": "185.193.165.115:27015"
        }
    ]
}
```
