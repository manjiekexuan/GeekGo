package main

import "fmt"

func main()  {

	newPersonInfos :=[...][3]string{

		[3]string{"小强","男","在职"},
		[3]string{"小李","男","在职"},
		[3]string{"小苏","女","在职"},
	}

	for i,val :=range newPersonInfos{
		fmt.Println(i,val)
	}

}
