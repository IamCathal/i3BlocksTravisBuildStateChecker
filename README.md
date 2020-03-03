# i3 blocks Travis CI build state checker

This is a nice little i3blocks indicator script to help you monitor the current build state of a github repository.
![](https://i.imgur.com/tZXUt7Q.png)
## Installation

1. Get the url for the build indicator from your repository e.g view the indicator as an image to get the direct url for it. The link will look like `https://api.travis-ci.org/raysan5/raylib.svg?branch=master`
2. Build the script with `go build buildState.go`
3. Place the executable in your i3blocks command folder. I'm on Ubuntu and it's `/usr/share/i3blocks`
4. Edit your i3blocks config file which can be found at `.config/i3/i3blocks.conf` and add this in: 
````
[buildStateGo]
color=#4ceb34
interval=30
````

## Future improvements

In the future I plan on adding functionality to track multiple repos and give more information about them.

