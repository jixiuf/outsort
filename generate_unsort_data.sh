#!/bin/sh
# 生成 乱序的数字到 in.txt 文件中
# 使用方式 ./generate_unsort_data.sh 100  生成100个数字
cnt=10000
if [ -n "$1"  ]; then
    cnt=$1;
fi
echo $cnt
rm -rf /tmp/seq.txt
for var in `seq $cnt`; do echo "$var">>/tmp/seq.txt; done;
while read i;do echo "$i $RANDOM";done</tmp/seq.txt|sort -k2n|cut -d" " -f1>in.txt
