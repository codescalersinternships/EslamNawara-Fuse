#!/bin/sh

go run ./demo/main.go mnt &
sleep 0.3

String=$(cat ./mnt/Name)
if [[ $String != 'Eslam' ]]; then
    echo 'TEST FAILED: file "Name" does not match struct value'
    fusermount -zu ./mnt
    exit 1
fi

Int=$(cat ./mnt/Age)
if [[ $Int != '22' ]]; then
    echo 'TEST FAILED: file "Age" does not match struct value'
    fusermount -zu ./mnt
    exit 1
fi

find ./mnt/Sub >> /dev/null
if [[ $? != 0 ]]; then
    echo 'TEST FAILED: dir "Sub" does not exist'
fi

value=$(cat ./mnt/Sub/SomeValue)
if [[ $value != 20 ]]; then
    echo 'TEST FAILED: file "SomeValue" does not match struct value'
    echo $value
    fusermount -zu ./mnt
    exit 1
fi

sleep 2
Name=$(cat ./mnt/Name)
if [[ $Name != 'not Eslam' ]]; then
    echo 'TEST FAILED: file "Name" does not match struct value'
    fusermount -zu ./mnt
    exit 1
fi

fusermount -zu ./mnt
echo 'TEST PASSED'
