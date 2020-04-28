//Author: Rithika Hebbar
package main
import "fmt"

func main() {
    
    var input string
    var key int
    fmt.Println("Enter the string to be encrypted:")
    fmt.Scanf("%s",&input)
    fmt.Println("Enter the key:")
    fmt.Scanf("%d",&key)
    for _, char:= range input{
        
        if rune(char)>='A' && rune(char)<='Z'{
            
            fmt.Printf(string(rotate(rune(char), 'A', key)))
        }else if rune(char)>='a' && rune(char)<='z'{
            
            fmt.Printf(string(rotate(rune(char), 'a', key)))
        }else{
            
            fmt.Printf(string(char))
        }
    }
    
}
func rotate(r rune,base, key int) rune{
    temp:=((int(r)-base)+key)%26
    temp=temp+base 
    return rune(temp)
}