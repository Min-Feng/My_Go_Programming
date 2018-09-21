package main

import "fmt"

type tree struct {
	key         int
	left, right *tree
}

//插入一個新的數值進入二元搜尋樹
func add(root *tree, value int) *tree {
	if root == nil {
		root = new(tree) //此步驟最為重要，讓root指標指向特定區塊，才可以賦值
		root.key = value
		return root
	}

	//二元搜尋樹的規則 ： Key(L)<Key(Current)<Key(R)
	if root.key > value {
		root.left = add(root.left, value)
		// root = add(root.left, value) 應該更換新的左子樹,而不是將左子樹取代整個root
	} else {
		root.right = add(root.right, value)
		// root = add(root.right, value) 錯誤想法 理由同上
	}
	return root
}

/* 根據二元搜尋樹規則
Key(L)<Key(Current)<Key(R)的性質
這正好與Inorder(LVR)之順序相同
因此，對整棵樹進行Inorder Traversal
就能夠對資料由小到大(依照Key)進行排序 sort */
func bulidBST(values []int) []int {

	//此步驟是將slice的數值,建構在搜尋二元樹
	var root *tree //此指標還沒有指向特定空間，為nil值
	for _, v := range values {
		root = add(root, v)
	}

	//建構搜尋二元樹，依照中序尋訪（Inorder Traversal）讀取樹的數值,就可以得到排序
	values = getDataOfBST(root, values[:0])
	return values //其實不需要回傳slice型別的values,因為都是對同一個底層陣列進行操作
}

func getDataOfBST(t *tree, values []int) []int {
	if t == nil {
		return values
	}

	values = getDataOfBST(t.left, values)
	values = append(values, t.key)
	values = getDataOfBST(t.right, values)
	return values
}

func main() {
	testData := []int{4, 6, 8, 3, 2, 0 ,-2,-4}
	returnData := bulidBST(testData)
	fmt.Println("returnData=", returnData) //證明都是對同一個底層陣列進行操作
	fmt.Println("testData=", testData)     //證明都是對同一個底層陣列進行操作
}
