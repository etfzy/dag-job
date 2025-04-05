package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/etfzy/dag-job/job"
)

func TestNormal(t *testing.T) {
	gjob := job.NewDAGJob()
	gjob.AddTask(job.NewTask("task1", func(ctx context.Context) error {
		fmt.Println("task1 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		fmt.Println("task1 end", time.Now().Format("2006-01-02 15:04:05.000"))
		return nil
	}))

	gjob.AddTask(job.NewTask("task2", func(ctx context.Context) error {
		fmt.Println("task2 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		fmt.Println("task2 end", time.Now().Format("2006-01-02 15:04:05.000"))
		return nil
	}))

	err := gjob.Build()
	if err != nil {
		panic(err)
	}

	state := gjob.Run(context.Background())
	s, _ := json.Marshal(state)
	fmt.Println(string(s))
}

func TestEndge(t *testing.T) {
	gjob := job.NewDAGJob()
	gjob.AddTask(job.NewTask("task1", func(ctx context.Context) error {
		fmt.Println("task1 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		fmt.Println("task1 end", time.Now().Format("2006-01-02 15:04:05.000"))
		return nil
	}))

	gjob.AddTask(job.NewTask("task2", func(ctx context.Context) error {
		fmt.Println("task2 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		fmt.Println("task2 end", time.Now().Format("2006-01-02 15:04:05.000"))
		return nil
	}))

	gjob.AddTask(job.NewTask("task3", func(ctx context.Context) error {
		fmt.Println("task3 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		fmt.Println("task3 end", time.Now().Format("2006-01-02 15:04:05.000"))
		return nil
	}))

	gjob.AddEdge("task1", "task2")

	err := gjob.Build()
	if err != nil {
		panic(err)
	}

	state := gjob.Run(context.Background())
	s, _ := json.Marshal(state)
	fmt.Println(string(s))
}

func TestMultiBranchs(t *testing.T) {
	gjob := job.NewDAGJob()
	gjob.AddTask(job.NewTask("task1", func(ctx context.Context) error {
		fmt.Println("task1 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return nil
	}))

	gjob.AddTask(job.NewTask("task2.1", func(ctx context.Context) error {
		fmt.Println("task2.1 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return nil
	}))

	gjob.AddTask(job.NewTask("task3", func(ctx context.Context) error {
		fmt.Println("task3 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return nil
	}))

	gjob.AddTask(job.NewTask("task4", func(ctx context.Context) error {
		fmt.Println("task4 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return nil
	}))

	gjob.AddTask(job.NewTask("task5", func(ctx context.Context) error {
		fmt.Println("task5 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return nil
	}))

	gjob.AddTask(job.NewTask("task2.2", func(ctx context.Context) error {
		fmt.Println("task2.2 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(10 * time.Second)
		return nil
	}))

	gjob.AddEdge("task1", "task2.1")
	gjob.AddEdge("task2.1", "task3")
	gjob.AddEdge("task2.1", "task4")
	gjob.AddEdge("task1", "task2.2")
	gjob.AddEdge("task3", "task5")
	gjob.AddEdge("task4", "task5")
	gjob.AddEdge("task2.2", "task5")

	err := gjob.Build()
	if err != nil {
		panic(err)
	}

	state := gjob.Run(context.Background())
	s, _ := json.Marshal(state)
	fmt.Println(string(s))

}

func TestCancel(t *testing.T) {
	gjob := job.NewDAGJob()
	gjob.AddTask(job.NewTask("task1", func(ctx context.Context) error {
		fmt.Println("task1 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return nil
	}))

	gjob.AddTask(job.NewTask("task2.1", func(ctx context.Context) error {
		fmt.Println("task2.1 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return nil
	}))

	gjob.AddTask(job.NewTask("task3", func(ctx context.Context) error {
		fmt.Println("task3 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return nil
	}))

	gjob.AddTask(job.NewTask("task4", func(ctx context.Context) error {
		fmt.Println("task4 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return nil
	}))

	gjob.AddTask(job.NewTask("task5", func(ctx context.Context) error {
		fmt.Println("task5 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return nil
	}))

	gjob.AddTask(job.NewTask("task2.2", func(ctx context.Context) error {
		fmt.Println("task2.2 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(10 * time.Second)
		return nil
	}))

	gjob.AddEdge("task1", "task2.1")
	gjob.AddEdge("task2.1", "task3")
	gjob.AddEdge("task2.1", "task4")
	gjob.AddEdge("task1", "task2.2")
	gjob.AddEdge("task3", "task5")
	gjob.AddEdge("task4", "task5")
	gjob.AddEdge("task2.2", "task5")

	err := gjob.Build()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(8 * time.Second)
		cancel()
	}()

	state := gjob.Run(ctx)
	s, _ := json.Marshal(state)
	fmt.Println(string(s))
}

// 环形回路
func TestError(t *testing.T) {
	gjob := job.NewDAGJob()
	gjob.AddTask(job.NewTask("task1", func(ctx context.Context) error {
		fmt.Println("task1 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return nil
	}))

	gjob.AddTask(job.NewTask("task2.1", func(ctx context.Context) error {
		fmt.Println("task2.1 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return nil
	}))

	gjob.AddTask(job.NewTask("task3", func(ctx context.Context) error {
		fmt.Println("task3 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return nil
	}))

	gjob.AddTask(job.NewTask("task4", func(ctx context.Context) error {
		fmt.Println("task4 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return errors.New("test4 error")
	}))

	gjob.AddTask(job.NewTask("task5", func(ctx context.Context) error {
		fmt.Println("task5 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		return nil
	}))

	gjob.AddTask(job.NewTask("task2.2", func(ctx context.Context) error {
		fmt.Println("task2.2 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(10 * time.Second)
		return nil
	}))

	gjob.AddEdge("task1", "task2.1")
	gjob.AddEdge("task2.1", "task3")
	gjob.AddEdge("task2.1", "task4")
	gjob.AddEdge("task1", "task2.2")
	gjob.AddEdge("task3", "task5")
	gjob.AddEdge("task4", "task5")
	gjob.AddEdge("task2.2", "task5")

	err := gjob.Build()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(8 * time.Second)
		cancel()
	}()

	state := gjob.Run(ctx)
	s, _ := json.Marshal(state)
	fmt.Println(string(s))
}

// 环形回路
func TestLoopBranchs(t *testing.T) {
	gjob := job.NewDAGJob()
	gjob.AddTask(job.NewTask("task1", func(ctx context.Context) error {
		fmt.Println("task1 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		fmt.Println("task1 end", time.Now().Format("2006-01-02 15:04:05.000"))
		return nil
	}))

	gjob.AddTask(job.NewTask("task2", func(ctx context.Context) error {
		fmt.Println("task2 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		fmt.Println("task2 end", time.Now().Format("2006-01-02 15:04:05.000"))
		return nil
	}))

	gjob.AddTask(job.NewTask("task3", func(ctx context.Context) error {
		fmt.Println("task3 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		fmt.Println("task3 end", time.Now().Format("2006-01-02 15:04:05.000"))
		return nil
	}))

	gjob.AddTask(job.NewTask("task4", func(ctx context.Context) error {
		fmt.Println("task4 start", time.Now().Format("2006-01-02 15:04:05.000"))
		time.Sleep(1 * time.Second)
		fmt.Println("task4 end", time.Now().Format("2006-01-02 15:04:05.000"))
		return nil
	}))

	gjob.AddEdge("task1", "task2")
	gjob.AddEdge("task1", "task3")
	gjob.AddEdge("task2", "task4")
	gjob.AddEdge("task3", "task4")
	gjob.AddEdge("task4", "task1")

	err := gjob.Build()
	if err != nil {
		panic(err)
	}
}
