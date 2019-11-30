

import "fmt"

func main()  {
	m := make([]string, 1, 15)
	fmt.Println(m)
	m[0] = "Digambar"
	changeMe(m)
	fmt.Println(m, len(m))

	mp := make(map[string]int)
	mp["age"] = 20
	changeMap(mp)
	fmt.Println(mp)
}

func changeMe(str []string){
	str = append(str, "Patil")
	fmt.Println(str)
	fmt.Println(len(str))
}

func changeMap(mp map[string]int)  {
	mp["age"] = 33
	fmt.Println(mp)
}