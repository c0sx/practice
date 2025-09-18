package leetcode3408

import (
	"container/heap"
)

type Task struct {
	taskId   int
	priority int
}

type TaskHeap []Task

func (th TaskHeap) Len() int {
	return len(th)
}

func (th TaskHeap) Less(i, j int) bool {
	if th[i].priority == th[j].priority {
		return th[i].taskId > th[j].taskId
	}

	return th[i].priority > th[j].priority
}

func (th TaskHeap) Swap(i, j int) {
	th[i], th[j] = th[j], th[i]
}

func (th *TaskHeap) Push(x any) {
	*th = append(*th, x.(Task))
}

func (th *TaskHeap) Pop() any {
	old := *th
	n := len(old)
	x := old[n-1]
	*th = old[0 : n-1]

	return x
}

type TaskManager struct {
	h          *TaskHeap
	users      map[int]int
	priorities map[int]int
}

func Constructor(tasks [][]int) TaskManager {
	h := &TaskHeap{}
	taskUsers := make(map[int]int)
	taskPriorities := make(map[int]int)

	for _, task := range tasks {
		heap.Push(h, Task{
			taskId:   task[1],
			priority: task[2],
		})

		taskUsers[task[1]] = task[0]
		taskPriorities[task[1]] = task[2]
	}

	return TaskManager{
		h:          h,
		users:      taskUsers,
		priorities: taskPriorities,
	}
}

func (tm *TaskManager) Add(userId int, taskId int, priority int) {
	heap.Push(tm.h, Task{
		taskId:   taskId,
		priority: priority,
	})

	tm.users[taskId] = userId
	tm.priorities[taskId] = priority
}

func (tm *TaskManager) Edit(taskId int, newPriority int) {
	heap.Push(tm.h, Task{
		taskId:   taskId,
		priority: newPriority,
	})

	tm.priorities[taskId] = newPriority
}

func (tm *TaskManager) Rmv(taskId int) {
	delete(tm.priorities, taskId)
	delete(tm.users, taskId)
}

func (tm *TaskManager) ExecTop() int {
	for tm.h.Len() > 0 {
		top := heap.Pop(tm.h).(Task)
		priority := -1
		if p, exists := tm.priorities[top.taskId]; exists {
			priority = p
		}

		if priority == top.priority {
			userId := tm.users[top.taskId]

			delete(tm.priorities, top.taskId)
			delete(tm.users, top.taskId)

			return userId
		}
	}

	return -1
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
