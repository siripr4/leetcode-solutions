package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func averageOfLevels(root *TreeNode) []float64 {

	if root == nil {
		return []float64{}
	}

	currentLevelNodes := []*TreeNode{root}
	var nextLevelNodes []*TreeNode
	averagePerLevel := []float64{}

	for len(currentLevelNodes) != 0 {
		nextLevelNodes = []*TreeNode{}
		sum := 0.0
		for _, node := range currentLevelNodes {
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
			sum += float64(node.Val)
		}
		averagePerLevel = append(averagePerLevel, sum/float64(len(currentLevelNodes)))
		currentLevelNodes = nextLevelNodes
	}

	return averagePerLevel
}
