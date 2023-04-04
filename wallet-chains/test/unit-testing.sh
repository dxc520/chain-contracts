#!/bin/bash
#脚本执行说明
#./smartContract-test.sh smartContractString
#smartContractString为多个合约名称，用逗号间隔开

set -e
base_dir=$(cd `dirname $0` && pwd)
cd $base_dir

smartContractString=${SMARTCONTRACTSTRING:-"CellarBaseNftERC721.sol,CellarCoinBaseERC20.sol"}

#判断合约是否指定
if [ -z "$smartContractString" ]
then
	echo "USAGE: ./smartContract-test.sh smartContractString"
	exit
fi

time=$(date "+%Y%m%d%H%M%S")
htmlname="$time.html"
fileName="/tmp/$time.html"
touch $fileName

#写入html初始内容
htmlStr1="<!DOCTYPE html>\n<html>\n<meta charset='utf-8'>\n<head>\n<title>SmartContract test log</title>\n"
htmlStr2="<style>\nb {color:green;}\n</style>\n</head>\n<body>\n<pre>\nSmartContract test log .......\n"
echo -e "$htmlStr1" >> $fileName
echo -e "$htmlStr2" >> $fileName


cd ../ 
ls

#处理合约名称
smartContracts=(${smartContractString//,/ })
for smartContract in ${smartContracts[@]}
do
	echo "slither test for $smartContract"
	
	echo -e "<b> slither test for $smartContract </b>" >> $fileName
	slither --solc-remaps @openzeppelin/=./node_modules/@openzeppelin/ --filter-paths  "node_modules" contracts/$smartContract 2>&1 | tee -a $fileName || echo
	#echo "$response" >> $fileName
	#sleep 5
done


htmlEndStr="</pre>\n</body>\n</html>"
echo -e "$htmlEndStr" >> $fileName

cd $base_dir
echo "upload $htmlname to bfs"
#上传到bfs
response=`curl -X POST "https://pnode.solarfs.io/un/file" -H "accept: application/json" -H "Content-Type: multipart/form-data" -F "days=365" -F "file=@$fileName;type=text/html"`

echo "get afid"
# 提取afid
afid=`echo "$response" | awk -F '"' '{print $12}'`

#拼接预览页面
echo "https://pnode.solarfs.io/dn/short/$afid-$htmlname"

rm -rf $htmlname
