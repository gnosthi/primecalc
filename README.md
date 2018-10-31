PrimeCalc v0.4
==============
[![Build Status](https://travis-ci.com/gnosthi/primecalc.svg?branch=master)](https://travis-ci.com/gnosthi/primecalc)|[![codecov](https://codecov.io/gh/gnosthi/primecalc/branch/master/graph/badge.svg)](https://codecov.io/gh/gnosthi/primecalc)|[![CircleCI](https://circleci.com/gh/gnosthi/primecalc/tree/master.svg?style=svg)](https://circleci.com/gh/gnosthi/primecalc/tree/master)

---------
#PrimeCalc

### Do what thou wilt shall be the whole of the Law.
Prime calc is an small command line application which calculates whether a number is prime.
It takes any amount of numbers as an argument and calculates their primeness.
Additionally it will also tell their prime factorizations.

## Installation
A simple binary installation can be done by downloading the latest release from the release page
and untarring it to your local home bin/ directory.
```
tar -xzvf primecalc_<VERSION>_<OS>_<ARCH>.tar.gz -C $HOME/bin/
```
Or you could install it system wide.
```
sudo tar -xzvf primecalc_<VERSION>_<OS>_<ARCH>.tar.gz -C /usr/bin/
```

Additionally we currently have docker builds available.
```
docker run gnosthi/primecalc:latest <numbers> 
```

Or install from source.
```
go get -v -u github.com/gnosthi/primecalc
cd $GOPATH/src/github.com/gnosthi/primecalc
make && make install
```

## Usage
Currently 'PrimeCalc' is rather limited in its scope. However we will expand functionality as time goes by.
Simple usage of PrimeCalc is to simply run it with a set of numbers as arguments.
```
$ primecalc 3301 111 93 13 11
3301 : prime
------------------------------------------------
111 : not prime : 37 * 3 = 111
37 : prime
------------------------------------------------
93 : not prime : 31 * 3 = 93
31 : prime
------------------------------------------------
13 : prime
------------------------------------------------
11 : prime
------------------------------------------------
```
You can also run primecalc as a docker container
```
docker run gnosthi/primecalc:latest 313
313 : prime
------------------------------------------------
```
