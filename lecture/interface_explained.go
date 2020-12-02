package lecture

// Type
// int, var, double, float, bool

// Structure = Type (Composition)

// Interface
// Empty interface = Any type/pointer. Syntax: interface{}
// Decoupling relationship

// Con nợ (Person) - chủ nợ (Person)
// Company is also "con no" or "chu no" // change 2

// Trong OOP, 1 class can only inherits one class but implements many interface

type ChuNoBehavior interface {
	ChoMuonTien(amount int, conNo ConNoBehavior)
	DoiTien()
}

type ConNoBehavior interface {
	MuonTien(amount int, chuNo ChuNoBehavior)
	TraTien()
}

type Company struct {
	Name  string
	Money int

	ConNo ConNoBehavior
	ChuNo ChuNoBehavior
}

type Bank struct {
	Name  string
	Money int
}

func (c *Company) ChoMuonTien(amount int, conNo ConNoBehavior) {
	c.Money -= amount
	c.ConNo = conNo
}

func (c *Company) DoiTien() {
	panic("implement me")
}

type Person struct {
	Name  string
	Money int
	ConNo *Person
	ChuNo *Person

	TienNo       int
	BankingMoney int // change 1
}

// Encapsulation
func (p *Person) NhanTien(amount int) {
	p.Money += amount
}

func (p *Person) ChoTien(amount int) {
	p.Money -= amount
}

func (p *Person) ChoMuonTien(amount int, conNo *Person) {
	p.ConNo = conNo
	p.ChoTien(amount)

	conNo.NhanTien(amount)
	//conNo.ChuNo = p
}

//func (p *Person) ChoMuonTien(conNo *Person, no int) {
//	p.Money -= no
//	conNo.Money += no
//	conNo.TienNo = no
//
//	p.ConNo = conNo
//	conNo.ChuNo = p
//}

func (p *Person) TraTien() {
	p.Money -= p.TienNo
	p.ChuNo.Money += p.TienNo

	p.ChuNo = nil
	p.ChuNo.ConNo = nil
}

func Test() {
	//me := Person{
	//	Name:  "Viet",
	//	Money: 1000,
	//}
	//
	//lam := Person{
	//	Name:  "Lam",
	//	Money: 1000,
	//}
	//
	//
	//me.ChoMuonTien(&lam, 100)
	//lam.TraTien()
}
