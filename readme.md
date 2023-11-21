# PutManyCirclesOnPicture
Put many circles on picture!
Simple as that. The circles are not really fancy, but the tool can still come handy if You'are seeking a way to put Your heat map on a picture knowing the pixel-wise coordinates.

# Build
You do not have to specify the GOOS and GOARCH if You want to build for the system You are building on.
``` bash
# for linux
GOOS=darwin GOARCH=amd64 go build .
# for macos
GOOS=darwin GOARCH=amd64 go build .
# for windows
GOOS=windows GOARCH=amd64 go build .
```
You can also just grab executables from the releases page.


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

Also I don't want anybody to install a JS interpreter on their machine just to run a little tool program.