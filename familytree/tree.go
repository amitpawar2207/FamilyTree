package familytree

import (
	"fmt"
	f "geektrust/fileoperations"
	"strings"
)

//GenerateFamilyTree creates the family tree of King shan
func GenerateFamilyTree(filePath string) (MemberList, error) {
	familyMembers := make(MemberList)
	inputs, fErr := f.ReadInputFile(filePath)
	if fErr != nil {
		return familyMembers, fErr
	}
	for _, input := range inputs {
		ProcessInputs(input, familyMembers)
	}
	return familyMembers, nil
}

//ProcessInputs method performs the operations according to the stated command
func ProcessInputs(input string, familyMembers MemberList) string {

	data := strings.Split(input, " ")
	if len(data) < 3 {
		return "Wrong Input : " + input
	}

	var result string
	action := classifyInput(input)

	switch action {

	case "add_child":
		result = familyMembers.addChildren(data[1], data[2], data[3])

	case "get_relationship":
		node := familyMembers.findNode(data[1])
		if node == nil {
			fmt.Println("PERSON_NOT_FOUND")
		} else {
			result = node.getRelationship(data[1], data[2])
		}

	case "create_root":
		familyMembers = createRootNode(data[1], data[2], familyMembers)

	case "add_partner":
		familyMembers, result = addPartner(data[1], data[2], data[3], familyMembers)

	default:
		fmt.Println("Wrong Input")
	}
	return result
}

func classifyInput(input string) string {
	input = strings.ToLower(input)
	if strings.Contains(input, "get_relationship") {
		return "get_relationship"
	} else if strings.Contains(input, "add_child") {
		return "add_child"
	} else if strings.Contains(input, "add_partner") {
		return "add_partner"
	} else if strings.Contains(input, "create_root") {
		return "create_root"
	}
	return ""
}

//FindRelationship returns the relation results or appropriate message
func (node *Member) findRelationship(name, relationship string) ([]string, string) {
	resultList := make([]string, 0)

	switch strings.ToLower(relationship) {

	case "sister-in-law":
		return node.findInLaws("female"), ""

	case "brother-in-law":
		return node.findInLaws("male"), ""

	case "son":
		return node.findChildren("male"), ""

	case "daughter":
		return node.findChildren("female"), ""

	case "siblings":
		return node.findSiblings()

	case "maternal-uncle":
		if node.Mother == nil || node.Mother.Mother == nil {
			return resultList, "NONE"
		}
		node = node.Mother.Mother.Childrens
		return node.findPaternalAuntOrMaternalUncle("male"), ""

	case "paternal-aunt":
		if node.Father == nil || node.Father.Mother == nil {
			return resultList, "NONE"
		}
		node = node.Father.Mother.Childrens
		return node.findPaternalAuntOrMaternalUncle("female"), ""

	case "maternal-aunt":
		return node.findMaternalAunt()

	case "paternal-uncle":
		return node.findPaternalUncle()

	case "mother":
		return nil, node.findMother()

	case "father":
		return nil, node.findFather()
	}
	return nil, ""
}

func (familyMembers MemberList) findNode(name string) *Member {
	if val, ok := familyMembers[strings.ToLower(name)]; ok {
		return val
	}
	return nil
}

func createRootNode(name, gender string, members MemberList) MemberList {
	node := Member{}
	node.Name = name
	node.Gender = strings.ToLower(gender)
	members[strings.ToLower(name)] = &node

	return members
}
