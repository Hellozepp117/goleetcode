package functions

func intToRoman(num int) string {
	var ans string
	for i:=1;i<=num/1000;i++{
		ans=ans+"M"
	}
	num=num%1000
	if num/900==1{
		ans =ans+"CM"
	}
	return
}
M1000
D500
C100
L50
X10
V5
I