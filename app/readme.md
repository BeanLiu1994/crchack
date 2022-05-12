# app

this program accepts a src file and a dst file 
and will create a new file has same size & crc64 by appending bytes to src file.

dst file should be bigger than src file

# example

```shell
$ ./app -src a.png -dst g.jpg -out o.png
src crc64 ECMA is cd501f701f24a286, len 8918, filename is a.png
dst crc64 ECMA is fd8a308b60e1b862, len 83728, filename is g.jpg
out crc64 ECMA is fd8a308b60e1b862, len 83728, filename is o.png
```