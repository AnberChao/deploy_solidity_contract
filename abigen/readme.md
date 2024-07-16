
生成abi和bin  solc为0.8.20版本
.\solc.exe --abi --bin --output-dir ./output erc20.sol

生成go文件
.\abigen.exe --bin .\output\MyToken.bin --abi .\output\MyToken.abi --pkg erc20 --type MyToken --out ..\contracts\mytoken\mytoken.go