# golook-cli
[![Build Status](https://travis-ci.org/ottenwbe/golook-cli.svg?branch=master)](https://travis-ci.org/ottenwbe/golook-cli)
[![codecov](https://codecov.io/gh/ottenwbe/golook-cli/branch/master/graph/badge.svg)](https://codecov.io/gh/ottenwbe/golook-cli)

Cli for [golook](https://github.com/ottenwbe/golook).

# Build

Ensure that go is installed and the GOPATH is set, then:

    go get github.com/ottenwbe/golook-cli

# Run

After that you can get started by calling the tool. It will show you a help message

    golook-cli 

For example, to call the /v1/info endpoint of a locally running golook broker, just type:    
    
    golook-cli info -u=http://localhost:8383
