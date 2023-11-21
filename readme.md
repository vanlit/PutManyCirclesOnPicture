# PutManyCirclesOnPicture
Put many circles on picture!  
![demo output - circles](https://github.com/vanlit/PutManyCirclesOnPicture/releases/download/untagged-69f47198d54beb488daf/little_output.jpg)  
Simple as that. The circles are not really fancy, but the tool can still come handy if You'are seeking a way to put Your heat map on a picture knowing the pixel-wise coordinates.

# Build
You do not have to specify the GOOS and GOARCH if You want to build for the system You are building on.
``` bash
# for linux
GOOS=linux GOARCH=amd64 go build -o PutManyCirclesOnPicture_lin64
# for macos
GOOS=darwin GOARCH=amd64 go build -o PutManyCirclesOnPicture_darwin64
# for windows
GOOS=windows GOARCH=amd64 go build -o PutManyCirclesOnPicture_win64.exe 
```
You can also just grab executable(s) from the releases page.
The commands above are exactly how I've built them.

# Run
### CSV file content example:
```
X	Y	Size	Color
890	840	15	aa33aa
```

PutManyCirclesOnPicture <input_image_path> <output_image_path> <csvFileWithCirclesDescriptions>
The output will be saved in jpeg.


# Why in Go? Why not Javascript?
Tiz fasta! 
![demo output - timings](https://github.com/vanlit/PutManyCirclesOnPicture/releases/download/untagged-69f47198d54beb488daf/timings.png)  
Also I don't want anybody to install a JS interpreter on their machine just to run a little tool program.