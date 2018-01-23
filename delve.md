Hey @payam @reuven, if you want to use the VS Code integrated debugging with Go you can give this a try:
[8:25 PM]
Guide for VS Code Go debugging

This is the official guide, and it was helpful but I ran into a lot of other issues that I detail fixes for below:
https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code

Here’s what works for me.

Prereq is the `delve` debugger, which is a little annoying to install: 
`brew install go-delve/delve/delve`
You’ll see an error when it tries to install the cert. When that happens, try this:
`cd ~/Library/Caches/homebrew`
`tar -xvf delve-1.0.0-rc.2.tar.gz`
`cd  delve-1.0.0-rc.2/scripts`
`./gencert.sh`
It will ask you for your password. Of course the `1.0.0-rc.2`part of the filename will change next delve release, hopefully they’ve resolved the cert issue by then.
Then run `brew install go-delve/delve/delve` again, it will work this time.

Also, if you don’t have `/Library/Developer/CommandLineTools/Library/PrivateFrameworks/LLDB.framework/Versions/A/Resources/debugserver` then you need to have Xcode command line tools installed from here: https://developer.apple.com/download/more/

Note: for some reason installing the command line tools via `xcode-select --install` or the App Store seem to not include the LLDB debugserver.


If `which lldb-server` turns up nothing, try `ln -s /Applications/Xcode.app/Contents/SharedFrameworks/LLDB.framework/Versions/A/Resources/lldb-server /usr/local/bin/lldb-server` to get it in your path.


I also had to set `"go.inferGopath": true` in the VS Code user settings.

Make a new launch config and make sure program and env are set correctly, e.g.:

```      "program": "${workspaceRoot}/cmd/server",
      "env": {
        "PORT": 50001
      },
```


I had to set at least one breakpoint before I started the debug process, try that if your breakpoints were added after starting the debugging session and are marked as “unverified”. (edited)
GitHub
Microsoft/vscode-go
vscode-go - An extension for VS Code which provides support for the Go language.
[8:28 PM]
If ^ works for you guys, let me know and I’ll clean it up and add it to the readme or wiki or wherever you think makes sense.