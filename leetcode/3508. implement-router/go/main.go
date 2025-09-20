package leetcode3508

import "strconv"

type Packet struct {
	source      int
	destination int
	timestamp   int
}

type Router struct {
	memoryLimit   int
	packets       map[string]Packet
	destinationTs map[int][]int
	queue         []string
}

func Constructor(memoryLimit int) Router {
	return Router{
		memoryLimit:   memoryLimit,
		packets:       map[string]Packet{},
		destinationTs: map[int][]int{},
		queue:         []string{},
	}
}

func (r *Router) AddPacket(source int, destination int, timestamp int) bool {
	key := r.getKey(source, destination, timestamp)
	if _, exists := r.packets[key]; exists {
		return false
	}

	if len(r.packets) >= r.memoryLimit {
		r.ForwardPacket()
	}

	r.packets[key] = Packet{source, destination, timestamp}

	if _, exists := r.destinationTs[destination]; !exists {
		r.destinationTs[destination] = []int{}
	}

	r.destinationTs[destination] = append(r.destinationTs[destination], timestamp)
	r.queue = append(r.queue, key)

	return true
}

func (r *Router) ForwardPacket() []int {
	if len(r.packets) == 0 || len(r.queue) == 0 {
		return []int{}
	}

	key := r.queue[0]
	r.queue = r.queue[1:]

	packet := r.packets[key]
	delete(r.packets, key)

	destination := packet.destination
	r.destinationTs[destination] = r.destinationTs[destination][1:]

	return []int{packet.source, packet.destination, packet.timestamp}
}

func (r *Router) GetCount(destination int, startTime int, endTime int) int {
	timestamps, exists := r.destinationTs[destination]
	if !exists || len(timestamps) == 0 {
		return 0
	}

	left := r.lowerBound(timestamps, startTime)
	right := r.upperBound(timestamps, endTime)

	return right - left
}

func (r *Router) getKey(source int, destination int, timestamp int) string {
	return strconv.Itoa(source) + "_" + strconv.Itoa(destination) + "_" + strconv.Itoa(timestamp)
}

func (r *Router) lowerBound(timestamps []int, target int) int {
	low, high := 0, len(timestamps)

	for low < high {
		mid := (low + high) / 2
		if timestamps[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func (r *Router) upperBound(timestamps []int, target int) int {
	low, high := 0, len(timestamps)

	for low < high {
		mid := (low + high) / 2
		if timestamps[mid] <= target { // the difference is here
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
