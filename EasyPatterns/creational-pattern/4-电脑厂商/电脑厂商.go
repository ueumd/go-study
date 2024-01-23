package main

import "fmt"

/*
练习：
设计一个电脑主板架构，电脑包括（显卡，内存，CPU）3个固定的插口
	1. 显卡具有显示功能（display，功能实现只要打印出意义即可）
	2. 内存具有存储功能（storage）
	3. cpu具有计算功能（calculate）

现有Intel厂商，nvidia厂商，Kingston厂商，均会生产以上三种硬件。
要求组装两台电脑，
	1台（Intel的CPU，Intel的显卡，Intel的内存）
	1台（Intel的CPU， nvidia的显卡，Kingston的内存）

用抽象工厂模式实现
*/

// ============================================ 抽象层 =========================================
type AbstractCpu interface {
	calculate()
}

type AbstractGpu interface {
	display()
}

type AbstractRam interface {
	storage()
}

/* ================================= Inter 产品族=========================== */
type InterCpu struct{}

func (this *InterCpu) calculate() {
	fmt.Println("Inter's CPU calculation")
}

type InterGpu struct{}

func (this *InterGpu) display() {
	fmt.Println("Inter's graphics display")
}

type InterRam struct{}

func (this *InterRam) storage() {
	fmt.Println("Inter's memory storage")
}

// ------------------------------- Inter工厂 -------------------------------
type InterFactory struct{}

// 生产cpu
func (this *InterFactory) CreateCpu() AbstractCpu {
	var cpu AbstractCpu
	cpu = new(InterCpu)
	return cpu
}

func (this *InterFactory) CreateGpu() AbstractGpu {
	var gpu AbstractGpu
	gpu = new(InterGpu)
	return gpu
}

func (this *InterFactory) CreateRam() AbstractRam {
	var ram AbstractRam
	ram = new(InterRam)
	return ram
}

/* ================================= Nvidia 产品族=========================== */
type NvidiaGpu struct{}

func (ng *NvidiaGpu) display() {
	fmt.Println("Nvidia GPU显示")
}

type NvidiaRam struct{}

func (nm *NvidiaRam) storage() {
	fmt.Println("Nvidia Mem存储")
}

type NvidiaCpu struct{}

func (nm *NvidiaCpu) calculate() {
	fmt.Println("Nvidia CPU计算")
}

// ------------------------------- Nvidia 工厂-------------------------------
type NvidiaFactory struct{}

func (n *NvidiaFactory) CreateGPU() AbstractGpu {
	var gpu AbstractGpu
	gpu = &NvidiaGpu{}
	return gpu
}

func (n *NvidiaFactory) CreateMem() AbstractRam {
	var mem AbstractRam
	mem = &NvidiaRam{}
	return mem
}

func (n *NvidiaFactory) CreateCPU() AbstractCpu {
	var cpu AbstractCpu
	cpu = &NvidiaCpu{}
	return cpu
}

/* ================================= Nvidia 产品族=========================== */
type KingstonGpu struct{}

func (k *KingstonGpu) display() {
	fmt.Println("Kingston GPU显示")
}

type KingstonRam struct{}

func (k *KingstonRam) storage() {
	fmt.Println("Kingston Mem存储")
}

type KingstonCpu struct{}

func (k *KingstonCpu) calculate() {
	fmt.Println("Kingston CPU计算")
}

// ------------------------------- kingston 工厂 ------------------------------
type kingstonFactory struct{}

func (n *kingstonFactory) CreateGPU() AbstractGpu {
	var gpu AbstractGpu
	gpu = &KingstonGpu{}
	return gpu
}

func (n *kingstonFactory) CreateRam() AbstractRam {
	var mem AbstractRam
	mem = &KingstonRam{}
	return mem
}

func (n *kingstonFactory) CreateCPU() AbstractCpu {
	var cpu AbstractCpu
	cpu = &KingstonCpu{}
	return cpu
}

/*=========================  业务逻辑层 =======================================*/
type ComputerFactory struct {
	CPU AbstractCpu
	GPU AbstractGpu
	Ram AbstractRam
}

func (c *ComputerFactory) show() {
	c.CPU.calculate()
	c.GPU.display()
	c.Ram.storage()
}

func main() {
	// 要求组装两台电脑，

	interFac := &InterFactory{}
	nvidiaFac := &NvidiaFactory{}
	kingstonFac := &kingstonFactory{}

	// 1台（Intel的CPU，Intel的显卡，Intel的内存）
	fmt.Println("---------- Intel的CPU，Intel的显卡，Intel的内存 -----------")
	computer1 := &ComputerFactory{
		CPU: interFac.CreateCpu(),
		GPU: interFac.CreateGpu(),
		Ram: interFac.CreateRam(),
	}
	computer1.show()

	// 1台（Intel的CPU， nvidia的显卡，Kingston的内存）
	fmt.Println("\n---------- Intel的CPU， nvidia的显卡，Kingston的内存 -----------")
	computer2 := &ComputerFactory{
		CPU: interFac.CreateCpu(),
		GPU: nvidiaFac.CreateGPU(),
		Ram: kingstonFac.CreateRam(),
	}
	computer2.show()
}
