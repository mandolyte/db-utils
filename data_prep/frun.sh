#!/bin/sh
cd $HOME/data/hier
FileReporter /bin hier_bin.csv
FileReporter /dev hier_dev.csv
FileReporter /etc hier_etc.csv
FileReporter /usr hier_usr.csv
FileReporter /lib hier_lib.csv
FileReporter /sbin hier_sbin.csv
FileReporter /sys hier_sys.csv
FileReporter /var hier_var.csv
catcsv -o hier_dataset.csv hier_*.csv
splitcsv -c 5 -keep=false < hier_dataset.csv | sort -u > hier_base.csv
wc -l hier_base.csv 
