package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestBST(t *testing.T) {
	RegisterFailHandler(Fail)     // registers the fail handler from ginkgo
	RunSpecs(t, "MaxHeap tests ") // hands over control to the ginkgo testing framework
}

var _ = Describe("Test the bst", func() {
	// Initializing the structs used to test the Insert(), Search(), and Remove() methods
	var bst bst
	exStr1 := exStruct{
		val:  "Val1",
		rank: 15,
	}
	exStr2 := exStruct{
		val:  "Val2",
		rank: 8,
	}
	exStr3 := exStruct{
		val:  "Val3",
		rank: 4,
	}
	exStr4 := exStruct{
		val:  "Val4",
		rank: 2,
	}
	exStr5 := exStruct{
		val:  "Val5",
		rank: 3,
	}
	exStr6 := exStruct{
		val:  "Val6",
		rank: 5,
	}
	exStr7 := exStruct{
		val:  "Val7",
		rank: 11,
	}
	exStr8 := exStruct{
		val:  "Val8",
		rank: 10,
	}
	exStr9 := exStruct{
		val:  "Val9",
		rank: 9,
	}
	exStr10 := exStruct{
		val:  "Val10",
		rank: 12,
	}
	exStr11 := exStruct{
		val:  "Val11",
		rank: 13,
	}
	exStr12 := exStruct{
		val:  "Val12",
		rank: 14,
	}
	exStr13 := exStruct{
		val:  "Val13",
		rank: 22,
	}
	exStr14 := exStruct{
		val:  "Val14",
		rank: 19,
	}
	exStr15 := exStruct{
		val:  "Val15",
		rank: 31,
	}
	exStr16 := exStruct{
		val:  "Val16",
		rank: 29,
	}
	exStr17 := exStruct{
		val:  "Val17",
		rank: 35,
	}

	Context("Test the exStruct methods", func() {

		It("Test exStruct fields", func() {
			exStr := exStruct{
				val:  "Random value",
				rank: 45,
			}
			Expect(exStr.getRank()).Should(Equal(45))
			Expect(exStr.getVal()).Should(Equal("Random value"))
		})
	})

	Context("Test the insert method", func() {
		It("Searching for each node after inserting to ensure correct insertion, done by comparing their rank values", func() {
			bst.Insert(exStr1)
			Expect(bst.Search(exStr1).bstNode.getRank()).Should(Equal(exStr1.rank))
			bst.Insert(exStr2)
			Expect(bst.Search(exStr2).bstNode.getRank()).Should(Equal(exStr2.rank))
			bst.Insert(exStr3)
			Expect(bst.Search(exStr3).bstNode.getRank()).Should(Equal(exStr3.rank))
			bst.Insert(exStr4)
			Expect(bst.Search(exStr4).bstNode.getRank()).Should(Equal(exStr4.rank))
			bst.Insert(exStr5)
			Expect(bst.Search(exStr5).bstNode.getRank()).Should(Equal(exStr5.rank))
			bst.Insert(exStr6)
			Expect(bst.Search(exStr6).bstNode.getRank()).Should(Equal(exStr6.rank))
			bst.Insert(exStr7)
			Expect(bst.Search(exStr7).bstNode.getRank()).Should(Equal(exStr7.rank))
			bst.Insert(exStr8)
			Expect(bst.Search(exStr8).bstNode.getRank()).Should(Equal(exStr8.rank))
			bst.Insert(exStr9)
			Expect(bst.Search(exStr9).bstNode.getRank()).Should(Equal(exStr9.rank))
			bst.Insert(exStr10)
			Expect(bst.Search(exStr10).bstNode.getRank()).Should(Equal(exStr10.rank))
			bst.Insert(exStr11)
			Expect(bst.Search(exStr11).bstNode.getRank()).Should(Equal(exStr11.rank))
			bst.Insert(exStr12)
			Expect(bst.Search(exStr12).bstNode.getRank()).Should(Equal(exStr12.rank))
			bst.Insert(exStr13)
			Expect(bst.Search(exStr13).bstNode.getRank()).Should(Equal(exStr13.rank))
			bst.Insert(exStr14)
			Expect(bst.Search(exStr14).bstNode.getRank()).Should(Equal(exStr14.rank))
			bst.Insert(exStr15)
			Expect(bst.Search(exStr15).bstNode.getRank()).Should(Equal(exStr15.rank))
			bst.Insert(exStr16)
			Expect(bst.Search(exStr16).bstNode.getRank()).Should(Equal(exStr16.rank))
			bst.Insert(exStr17)
			Expect(bst.Search(exStr17).bstNode.getRank()).Should(Equal(exStr17.rank))
		})
		It("Tests the right node pointer of some nodes", func() {
			Expect(bst.root.right.bstNode.getRank()).Should(Equal(exStr13.rank))
			Expect(bst.Search(exStr2).right.bstNode.getRank()).Should(Equal(exStr7.rank))
			Expect(bst.Search(exStr7).right.bstNode.getRank()).Should(Equal(exStr10.rank))
		})
		It("Tests the parent node pointer of some nodes", func() {
			Expect(bst.root.left.parent.bstNode.getRank()).Should(Equal(exStr1.rank))
			Expect(bst.Search(exStr9).parent.bstNode.getRank()).Should(Equal(exStr8.rank))
			Expect(bst.Search(exStr3).parent.bstNode.getRank()).Should(Equal(exStr2.rank))
		})
		It("Tests the left node pointer of some nodes", func() {
			Expect(bst.root.left.bstNode.getRank()).Should(Equal(exStr2.rank))
			Expect(bst.Search(exStr15).left.bstNode.getRank()).Should(Equal(exStr16.rank))
			Expect(bst.Search(exStr7).left.bstNode.getRank()).Should(Equal(exStr8.rank))
		})
	})

	Context("Test the search method", func() {
		// Initializing two more structs to test Search()
		randomStruct1 := exStruct{
			val:  "Value 1",
			rank: 565,
		}
		randomStruct2 := exStruct{
			val:  "Value 2",
			rank: 439,
		}

		It("Test Node search method", func() {
			Expect(bst.Search(exStr7).bstNode.getVal()).Should(Equal("Val7"))
			bst.Insert(randomStruct1)
			bst.Remove(randomStruct1)
			bst.Insert(randomStruct2)
			bst.Remove(randomStruct2)
			Node := bst.Search(exStr9).right
			Expect(Node).ShouldNot(Equal(randomStruct2))
		})
	})

	Context("Test the remove method", func() {
		It("Case in which node to be removed has two children", func() {
			bst.Remove(exStr15)
			Expect(bst.root.right.right.bstNode.getRank()).Should(Equal(exStr16.rank))
			bst.Remove(exStr3)
			Expect(bst.root.left.left.bstNode.getRank()).Should(Equal(exStr5.rank))
			bst.Remove(exStr7)
			Expect(bst.root.left.right.bstNode.getRank()).Should(Equal(exStr8.rank))
			bst.Remove(exStr9)
			bst.Remove(exStr7) // Nothing will be removed here
			Expect(bst.root.left.right.bstNode.getRank()).Should(Equal(exStr8.rank))
			bst.Remove(exStr13)
			Expect(bst.root.right.bstNode.getRank()).Should(Equal(exStr14.rank))
			bst.Remove(exStr2)
			Expect(bst.root.left.bstNode.getRank()).Should(Equal(exStr6.rank))

		})
		It("Case in which node to be removed has one child", func() {
			bst.Remove(exStr16)
			Expect(bst.root.right.right.bstNode.getRank()).Should(Equal(exStr17.rank))
			//Expect(bst.Search(exStr16).bstNode.getRank()).ShouldNot(Equal(exStr16.rank))
			bst.Remove(exStr14)
			Expect(bst.root.right.bstNode.getRank()).Should(Equal(exStr17.rank))
			//fmt.Print("\n", bst.root.left.bstNode.getRank(), bst.root.left.right.bstNode.getRank(), bst.root.left.right.right.bstNode.getRank(), bst.root.left.right.right.right.bstNode.getRank(), bst.root.left.right.left.bstNode.getRank(), bst.root.left.right.left.parent.bstNode.getRank(), "\n")
			//Expect(bst.Search(exStr9).bstNode.getRank()).ShouldNot(Equal(exStr9.rank))
			Expect(bst.root.left.right.bstNode.getRank()).Should(Equal(exStr8.rank))
			bst.Remove(exStr10)
			Expect(bst.Search(exStr8).right.bstNode.getRank()).Should(Equal(exStr11.rank))
			bst.Remove(exStr11)
			Expect(bst.Search(exStr8).right.bstNode.getRank()).Should(Equal(exStr12.rank))

			bst.Remove(exStr5)
			Expect(bst.root.left.left.bstNode.getRank()).Should(Equal(exStr4.rank))
		})

		It("Case in which node to be removed has no children", func() {
			bst.Remove(exStr12)
			bst.Remove(exStr9)
			bst.Remove(exStr4)
			bst.Remove(exStr8)
			bst.Remove(exStr6)
			bst.Remove(exStr17)

		})
		It("Removing last nodes to empty the tree", func() {
			bst.Remove(exStr1)
			Expect(bst.isEmpty()).Should(Equal(true))
		})
		It("Remove test case in which node.right != nil but less than it's parent", func() {
			bst.Insert(exStr1)
			Expect(bst.isEmpty()).Should(Equal(false))
			bst.Insert(exStr2)
			bst.Insert(exStr7)
			bst.Remove(exStr2)
			bst.Remove(exStr7)
			bst.Remove(exStr1)
			Expect(bst.isEmpty()).Should(Equal(true))
		})

		It("Remove test case in which removeNode has 1 child on the left but is greater than it's parent", func() {
			bst.Insert(exStr1)
			bst.Insert(exStr17)
			bst.Insert(exStr16)
			Expect(bst.root.right.left.bstNode.getRank()).Should(Equal(exStr16.rank))
			bst.Remove(exStr17)
			bst.Remove(exStr16)
			bst.Remove(exStr1)
			Expect(bst.isEmpty()).Should(Equal(true))
		})
	})
})
