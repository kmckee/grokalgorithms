#! /bin/bash
ls *.go | entr -s 'clear && richgo test'
