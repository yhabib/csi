#!/bin/bash

my_array=(1 apple "long string" $(gdate))
echo "${my_array[1]}"
my_array[1]=orange
echo "${my_array[1]}"
echo ${my_array[2]}
echo ${my_array[3]}
