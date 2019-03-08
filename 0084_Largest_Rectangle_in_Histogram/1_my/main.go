package main

type node struct {
	height int
	width  int
}

func largestRectangleArea(heights []int) int {
	max := 0
	// height: width (only the heighst heigt is recorded among who the same width)
	records := map[int]int{}
	for _, x := range heights {
		nearestNode := node{}
		for k := range records {
			if k < x {
				records[k]++
			} else if k > x {
				switch nearestNode.height {
				case 0:
					nearestNode = node{k, records[k]}
				default:
					if k < nearestNode.height {
						nearestNode = node{k, records[k]}
					}
				}
				area := records[k] * k
				if area > max {
					max = area
				}
				delete(records, k)
			}
		}
		if _, ok := records[x]; !ok {
			records[x] = nearestNode.width + 1
		} else {
			records[x]++
		}
	}

	for k := range records {
		area := records[k] * k
		if area > max {
			max = area
		}
	}
	return max
}
