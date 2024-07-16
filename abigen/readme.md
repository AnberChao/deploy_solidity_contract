
生成abi和bin
.\solc.exe --abi --bin --output-dir ./output erc20.sol

生成go文件
.\abigen.exe --bin .\output\MyToken.bin --abi .\output\MyToken.abi --pkg erc20 --type MyToken --out ..\contracts\mytoken\mytoken.go