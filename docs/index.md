![CICD](https://github.com/SecPlugs/go-kit/workflows/CICD/badge.svg)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Daily Test](https://github.com/SecPlugs/go-kit/workflows/DailyTest/badge.svg)
## Summary
{% include_relative _summary.md %}

## About
{% include_relative _detail.md %}

## Quick Start
Obtain the go module by simply running the commands below.
```console
GO111MODULE=on go get github.com/secplugs/go-kit/filescan
```
You'll now have the module in your `$GOPATH`

## Usage
Here, a very simple example of how to integrate file scan with your go code base is provided
{% include_relative usage.md %}

To avail yourself of the complete set of vendors and features provided by secplugs, and to have complete privacy
for the scans you submit, please create an accout at secplugs.com, login to the portal and create a new API key.
Please not that this API key is confidential.

After registering with secplugs.com and creating an API key, the only change to the code sample about would be
{% include_relative registered_usage.md %}

Everything else remains the same.

## Contact
Having trouble? [Contact Secplugs ](https://secplugs.com/contacts)
