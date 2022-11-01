func isHappy(n int) bool {
   var num,sum int 
    happy:= map[int]bool{}
    for {
        if n ==1 {
            return true
        }         
    for n!=1{
        num = n%10
        sum+=num*num
        n=n/10
    }
        if happy[sum] == true || sum == 0 {
            return false
        }
     happy[sum] = true 
    }
}
