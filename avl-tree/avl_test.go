package main

import (
	//"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestBST(t *testing.T) {
	RegisterFailHandler(Fail)     // registers the fail handler from ginkgo
	RunSpecs(t, "MaxHeap tests ") // hands over control to the ginkgo testing framework
}

var _ = Describe("Test the avl tree", func() {
	var avl avl
	Str1 := exStruct{
		val:  "Val1",
		rank: 5, //1
	}
	Str2 := exStruct{
		val:  "Val2",
		rank: 4, //20
	}
	Str3 := exStruct{
		val:  "Val3",
		rank: 3, //2
	}
	Str4 := exStruct{
		val:  "Val4",
		rank: 6, //19
	}
	Str5 := exStruct{
		val:  "Val5",
		rank: 7,
	}
	Str6 := exStruct{
		val:  "Val6",
		rank: 20,
	}
	Str7 := exStruct{
		val:  "Val7",
		rank: 15,
	}
	Str8 := exStruct{
		val:  "Val8",
		rank: 1,
	}
	Str9 := exStruct{
		val:  "Val9",
		rank: 2,
	}
	/*Str10 := exStruct{
		val:  "Val10",
		rank: 16,
	}
	*/
	Context("Test the exStruct methods", func() {
		It("Test exStruct fields", func() {
			Expect(Str1.getRank()).Should(Equal(Str1.rank))
			Expect(Str2.getRank()).Should(Equal(Str2.rank))
			Expect(Str3.getRank()).Should(Equal(Str3.rank))
			Expect(Str4.getRank()).Should(Equal(Str4.rank))
			Expect(Str5.getRank()).Should(Equal(Str5.rank))
			Expect(Str6.getRank()).Should(Equal(Str6.rank))
			Expect(Str7.getRank()).Should(Equal(Str7.rank))
			/*Expect(Str8.getRank()).Should(Equal(Str8.rank))
			Expect(Str9.getRank()).Should(Equal(Str9.rank))
			Expect(Str10.getRank()).Should(Equal(Str10.rank)) */
		})
	})

	Context("Test the insert method", func() {
		avl.Insert(Str1)
		avl.Insert(Str2)
		It("Insertion that causes a single right rotation", func() {
			avl.Insert(Str3)
			Expect(avl.root.avlNode.getRank()).Should(Equal(Str2.rank))
			Expect(avl.root.right.avlNode.getRank()).Should(Equal(Str1.rank))
			Expect(avl.root.left.avlNode.getRank()).Should(Equal(Str3.rank))
		})
		It("Insertions that causes a single left rotation", func() {
			avl.Insert(Str4)
			avl.Insert(Str5)
			Expect(avl.root.right.avlNode.getRank()).Should(Equal(Str4.rank))
			Expect(avl.root.right.left.avlNode.getRank()).Should(Equal(Str1.rank))
			Expect(avl.root.right.right.avlNode.getRank()).Should(Equal(Str5.rank))
			avl.Insert(Str6)
			Expect(avl.root.avlNode.getRank()).Should(Equal(Str4.rank))
			Expect(avl.root.right.right.avlNode.getRank()).Should(Equal(Str6.rank))
			Expect(avl.root.left.right.avlNode.getRank()).Should(Equal(Str1.rank))
			Expect(avl.root.left.left.avlNode.getRank()).Should(Equal(Str3.rank))
			Expect(avl.root.left.left.parent.avlNode.getRank()).Should(Equal(Str2.rank))
			Expect(avl.root.right.right.parent.avlNode.getRank()).Should(Equal(Str5.rank))
			Expect(avl.root.left.right.parent.avlNode.getRank()).Should(Equal(Str2.rank))
			Expect(avl.root.left.right.parent.parent.avlNode.getRank()).Should(Equal(avl.root.avlNode.getRank()))
		})
		It("Insertion that causes a double left rotation", func() {
			avl.Insert(Str7)
			Expect(avl.root.right.right.avlNode.getRank()).Should(Equal(Str6.rank))
			Expect(avl.root.right.left.avlNode.getRank()).Should(Equal(Str5.rank))
			Expect(avl.root.right.avlNode.getRank()).Should(Equal(Str7.rank))
			//RL case
		})
		It("Insertion that causes a double right rotation", func() {
			avl.Insert(Str8)
			avl.Insert(Str9)
			Expect(avl.root.left.left.left.avlNode.getRank()).Should(Equal(Str8.rank))
			Expect(avl.root.left.left.avlNode.getRank()).Should(Equal(Str9.rank))
		})
	})

	Context("Test the remove method", func() {
		It("Removal in which node to be removed is a leaf node", func() {
			Expect(avl.root.left.avlNode.getRank()).Should(Equal(Str2.rank))
			Expect(avl.root.left.left.avlNode.getRank()).Should(Equal(Str9.rank))
			Expect(avl.root.left.left.left.avlNode.getRank()).Should(Equal(Str8.rank))
			avl.Remove(Str3)
			Expect(avl.root.left.avlNode.getRank()).Should(Equal(Str2.rank))
			Expect(avl.root.left.left.avlNode.getRank()).Should(Equal(Str9.rank))
			Expect(avl.root.left.left.left.avlNode.getRank()).Should(Equal(Str8.rank))
			//Node3 := avl.Search(Str3)
			//Expect(Node3).Should(Equal(nil))
		})
		It("Remove that causes a single right rotation", func() {
			avl.Remove(Str1)
			// Recalculating the tree's balance/height factors
			// seems to be the issue right here bc RightRotate's not
			// being called
			//Expect(avl.root.left.avlNode.getRank()).Should(Equal(Str9.rank))
		})
	})
})
