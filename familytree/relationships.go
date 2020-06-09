package familytree

import (
	"strings"
)

func (node *Member) findInLaws(gender string) []string {
	var partnersGender string
	if gender == "male" {
		partnersGender = "female"
	} else {
		partnersGender = "male"
	}
	resultList := make([]string, 0)
	elementNode := node
	if node.Mother != nil {
		node = node.Mother.Childrens
		for {
			if node.Name == elementNode.Name {
				if node.Siblings == nil {
					break
				}
			} else {
				if node.Gender == partnersGender && node.Partner != nil {
					resultList = append(resultList, node.Partner.Name)
				}
				if node.Siblings == nil {
					break
				}
			}
			node = node.returnSiblings()
		}
	}
	node = elementNode
	if node.Partner != nil {
		if node.Partner.Mother != nil {
			node = node.Partner.Mother.Childrens
			for {
				if node.Partner != nil {
					if node.Gender == gender && node.Partner.Name != elementNode.Name {
						resultList = append(resultList, node.Name)
					}
				} else if node.Gender == gender {
					resultList = append(resultList, node.Name)
				}
				if node.Siblings == nil {
					break
				}
				node = node.returnSiblings()
			}
		}
	}
	return resultList
}

func (node *Member) findChildren(gender string) []string {
	resultList := make([]string, 0)
	if node.Childrens != nil {
		node = node.Childrens
		if node.Gender == gender {
			resultList = append(resultList, node.Name)
		}
		if node.Siblings != nil {
			for {
				node = node.returnSiblings()
				if node != nil {
					if node.Gender == gender {
						resultList = append(resultList, node.Name)
					}
				} else {
					break
				}
			}
		}
	}
	return resultList
}

func (node *Member) findSiblings() ([]string, string) {
	name := node.Name
	resultList := make([]string, 0)
	if node.Mother == nil {
		return resultList, "NONE"
	}
	if node.Mother.Childrens != nil {
		node = node.Mother.Childrens
		if strings.ToLower(node.Name) != strings.ToLower(name) {
			resultList = append(resultList, node.Name)
		}
		for node.Siblings != nil {
			node = node.returnSiblings()
			if strings.ToLower(node.Name) != strings.ToLower(name) {
				resultList = append(resultList, node.Name)
			}
		}
	}
	return resultList, ""
}

func (node *Member) findPaternalAuntOrMaternalUncle(gender string) []string {
	resultList := make([]string, 0)
	for {
		if node != nil {
			if node.Gender == gender {
				resultList = append(resultList, node.Name)
			}
		} else {
			break
		}
		node = node.returnSiblings()
	}
	return resultList
}

func (node *Member) findMaternalAunt() ([]string, string) {
	resultList := make([]string, 0)
	if node.Mother == nil && node.Mother.Mother == nil {
	}
	motherNode := node.Mother
	node = node.Mother.Mother.Childrens
	for {
		if node != nil {
			if node.Gender == "female" && node.Name != motherNode.Name {
				resultList = append(resultList, node.Name)
			}
		} else {
			break
		}
		node = node.returnSiblings()
	}
	return resultList, ""
}

func (node *Member) findPaternalUncle() ([]string, string) {
	resultList := make([]string, 0)
	if node.Father == nil || node.Father.Mother == nil {
		return resultList, "NONE"
	}
	fatherNode := node.Father
	node = node.Father.Mother.Childrens
	for {
		if node != nil {
			if node.Gender == "male" && node.Name != fatherNode.Name {
				resultList = append(resultList, node.Name)
			}
		} else {
			break
		}
		node = node.returnSiblings()
	}
	return resultList, ""
}

func (node *Member) findMother() string {
	if node.Mother == nil {
		return "NONE"
	}
	return node.Mother.Name
}

func (node *Member) findFather() string {
	if node.Father == nil {
		return "NONE"
	}
	return node.Father.Name
}

func (node Member) returnSiblings() *Member {
	return node.Siblings
}
