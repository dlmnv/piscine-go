#! /bin/bash

find -name '*.sh' | cut -d '.' -f2 | sed 's/\///g' | sed 's/test//g' | sed 's/sh//g'