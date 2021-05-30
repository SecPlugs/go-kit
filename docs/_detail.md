#### What is the Go Kit?

A go module (`github.com/secplugs/go-kit/filescan`) that can be readily integrated into your code base to submit data for malware analysis.  
Currently the plugin supports file scans.

#### What are the features?

- __File Scanning__ - Easy methods to allow uploading of files for scanning
- __Secplugs Portal__ - With a registered API key you can access all the core Secplugs features via the portal.

#### How does it work?

The library supports all the standard Secplugs functionality allowing access to file analysis functionality and makes it easier than integrating directly with the REST APIs.

#### How do I get started?

To get started download the module (via `go get github.com/secplugs/go-kit/filescan`) and use it in your Go code. This will work out of the box.

To use additional features and the privacy of your own account, after registering in Secplugs.com, login with your username and create an API key to use with your Go code. 
Create a new instance of the scan client with the API key and use it in your code.
Use can then use the Secplugs console to view activity, run reports and do deeper retrospective threat analysis.

