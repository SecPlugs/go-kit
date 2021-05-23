### url_scan.sh - Url Analysis Script
Use this script for submitting urls for analysis
To scan a url, pass it as a parameter  
E.g. below shows scanning of a Secplugs test url
```sh
./url_scan.sh https://example.com/?param=e81e043973f21d036e39fe
```
### file_scan.sh - File Analysis Script
Use this script for submitting files for analysis
To test connectivity run without no argument, it will default to scanning eicar.  
```sh
./file_scan.sh
```
To scan a named file, pass the file as a parameter  
E.g. to have it scan itself use below.
```sh
./file_scan.sh ./file_scan.sh
```
To scan the contents of named directory, specify a directory as parameter
E.g. to scan the contents of the /tmp
```sh
./file_scan.sh /tmp
```
