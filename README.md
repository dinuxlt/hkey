### HKey

It's imposible to paste to IPMI/iDrac/iLo remote/virtual console. This small util will type text for you to any active window, including virtual console.

### Usage

  -m, --message= Message to type  
  -w, --wait=    Time in seconds to wait before typing the message (default: 3)  
  -t, --time=    Time in ms between letters (default: 500)  
  -r, --return   Send ENTER key after message  

### Build
```
git clone git@github.com:dinuxlt/hkey.git
cd hkey
go build
```
