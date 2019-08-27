package main

import "fmt"

type t struct {
	ref   int
	value string
	link  *GlistNode
}

type GlistNode struct {
	utype string
	info  t
	next  *GlistNode
	pre   *GlistNode
}

type Glist struct {
	head   *GlistNode
	tail   *GlistNode
	length int
}

func Generate_List() *Glist {
	G := new(Glist)
	G.length = 0
	return G
}

func (li *Glist) add(info t, utype string, target *GlistNode) {
	p := new(GlistNode)
	p.info = info
	p.utype = utype
	if target == nil {
		if li.length == 0 {
			li.head = p
			li.tail = p
			//fmt.Print("Here\n")
		} else {
			fmt.Print("Error")
		}
	} else {
		var temp *GlistNode
		temp = target.next
		target.next = p
		p.next = temp
		temp.pre = p
		p.pre = target
		if target == li.tail {
			li.tail = target.next
		}
	}

	li.length++
}

func (li *Glist) delete(target *GlistNode) {
	if target == nil {
		if li.length == 0 {
			fmt.Print("Blank List\n")
		} else {
			fmt.Print("Invaild target\n")
		}

	} else {
		temp := target.pre
		temp.next = target.next
		target.next.pre = temp
	}
}

func search(item string, p *GlistNode) {
	for p != nil {
		if p.utype == "POINTER" {
			search(item, p.info.link)
		} else if p.utype == "DATA" {
			if item == p.info.value {
				fmt.Print("found")
			} else {
				p = p.next
			}
		}
	}

}

func (li *Glist) getoveralllength(len int32, p *GlistNode) int32 {
	for p != nil {
		if p.utype == "POINTER" {
			li.getoveralllength(len, p.info.link)
		} else if p.utype == "DATA" {
			p = p.next
			len++
		}
	}
	return len
}

func (li *Glist) searchp(item string) {
	search(item, li.head)
}
