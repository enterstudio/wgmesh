
package main


import (

	"fmt"
)

func test_NodeInformationBase() {

	var err error

	fmt.Println("Testing Node Information Base")

	nib := NewNodeInformationBase()

	ni, _ := NewNodeInfo("hoge", "1.1.1.1", "1.1.1.1:80")


	nib.AddNodeInfo(ni)

	fmt.Printf("Add: %v, %p\n", ni, ni)

	find_ni := nib.FindNodeInfo("hoge")
	if find_ni != nil {
		fmt.Printf("found: %v, %p\n", find_ni, find_ni)
	}

	fmt.Println("add new node asdf")
	ni, err = NewNodeInfo("asdf", "1.1.1.1.1", "1.1.1.1:80")
	if err != nil {
		fmt.Println(err)
	} else {
		nib.AddNodeInfo(ni)
	}

	fmt.Println("add new node qwer")
	ni, err = NewNodeInfo("qwer", "1.1.1.1", "1.1.1.1:asdf")
	if err != nil {
		fmt.Println(err)
	} else {
		nib.AddNodeInfo(ni)
	}

	ni, err = NewNodeInfo("qwer", "1.1.1.1", "1.1.1.1:9999")
	if err != nil {
		fmt.Println(err)
	} else {
		nib.AddNodeInfo(ni)
	}


	fmt.Println("add new node huge")
	ni, _ = NewNodeInfo("huga", "1.1.1.1", "1.1.1.1:80")
	nib.AddNodeInfo(ni)

	fmt.Println("Walk node table")
	for nii := range nib.ForeachNodeInfo() {
		fmt.Printf("%v\n", nii)
	}

	fmt.Println("del node qwer")
	nib.DelNodeInfo("qwer")

	fmt.Println("del node qwer again")
	err = nib.DelNodeInfo("qwer")
	if err != nil {
		fmt.Printf("delnodeinfo err: %s\n", err)
	}

	fmt.Println("Walk node table")
	for nii := range nib.ForeachNodeInfo() {
		fmt.Printf("%v\n", nii)
	}

	fmt.Println("del node 12345")
	err = nib.DelNodeInfo("12345")
	if err != nil {
		fmt.Printf("delnodeinfo err: %s\n", err)
	}
}

func Test(opts Options) {

	test_NodeInformationBase()

}
