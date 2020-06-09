package familytree_test

import (
	"fmt"
	"geektrust/familytree"
	"testing"
)

type inputs struct {
	input    string
	fMembers familytree.MemberList
}

func TestProcessInputs(t *testing.T) {

	members := make(familytree.MemberList)
	shan := familytree.Member{}
	shan.Name = "Shan"
	shan.Gender = "male"
	members["shan"] = &shan

	familyMembers, err := familytree.GenerateFamilyTree("../resources/family_members.txt")
	if err != nil {
		fmt.Println("Error while generating family tree ", err)
	}

	testData1 := []struct {
		testName  string
		inputData inputs
		result    string
	}{
		{
			"person not found",
			inputs{
				input:    "ADD_PARTNER Minga Leena Female",
				fMembers: members,
			},
			"PERSON_NOT_FOUND",
		},
		{
			"partner added successfully",
			inputs{
				input:    "ADD_PARTNER Laki Surya Female",
				fMembers: familyMembers,
			},
			"PARTNER_ADDITION_SUCCEEDED",
		},
		{
			"partner addition partner already added",
			inputs{
				input:    "ADD_PARTNER Shan Shakti Female",
				fMembers: familyMembers,
			},
			"PARTNER_ADDITION_FAILED",
		},
		{
			"partner additon wrong gender",
			inputs{
				input:    "ADD_PARTNER Shan Anga Male",
				fMembers: familyMembers,
			},
			"PARTNER_ADDITION_FAILED",
		},
		{
			"add child parents not created",
			inputs{
				input:    "ADD_CHILD Arjun Ahita Female",
				fMembers: familyMembers,
			},
			"PERSON_NOT_FOUND",
		},
		{
			"add child to Male",
			inputs{
				input:    "ADD_CHILD Shan Chit Male",
				fMembers: familyMembers,
			},
			"CHILD_ADDITION_FAILED",
		},
		{
			"add child valid input",
			inputs{
				input:    "ADD_CHILD Krithi Sagar Male",
				fMembers: familyMembers,
			},
			"CHILD_ADDITION_SUCCEEDED",
		},
		{
			"add child duplicate name",
			inputs{
				input:    "ADD_CHILD Anga Chit Male",
				fMembers: familyMembers,
			},
			"CHILD_ADDITION_FAILED",
		},
		{
			"get relationship son",
			inputs{
				input:    "GET_RELATIONSHIP Shan Son",
				fMembers: familyMembers,
			},
			"Chit Ish Vich Aras",
		},
		{
			"get relationship siblings none",
			inputs{
				input:    "GET_RELATIONSHIP Shan Siblings",
				fMembers: familyMembers,
			},
			"NONE",
		},
		{
			"get relationship Daughter",
			inputs{
				input:    "GET_RELATIONSHIP Shan Daughter",
				fMembers: familyMembers,
			},
			"Satya",
		},
		{
			"get relationship father none",
			inputs{
				input:    "GET_RELATIONSHIP Shan Father",
				fMembers: familyMembers,
			},
			"NONE",
		},
		{
			"get relationship mother none",
			inputs{
				input:    "GET_RELATIONSHIP Shan Mother",
				fMembers: familyMembers,
			},
			"NONE",
		},
		{
			"get relationship mother",
			inputs{
				input:    "GET_RELATIONSHIP Chit Mother",
				fMembers: familyMembers,
			},
			"Anga",
		},
		{
			"get relationship father",
			inputs{
				input:    "GET_RELATIONSHIP Chit Father",
				fMembers: familyMembers,
			},
			"Shan",
		},
		{
			"get relationship maternal-uncle",
			inputs{
				input:    "GET_RELATIONSHIP Yodhan Maternal-Uncle",
				fMembers: familyMembers,
			},
			"Vritha",
		},
		{
			"get relationship maternal-aunt none",
			inputs{
				input:    "GET_RELATIONSHIP Shan Maternal-Uncle",
				fMembers: familyMembers,
			},
			"NONE",
		},
		{
			"get relationship paternal-uncle none",
			inputs{
				input:    "GET_RELATIONSHIP Yodhan Paternal-Uncle",
				fMembers: familyMembers,
			},
			"NONE",
		},
		{
			"get relationship maternal aunt",
			inputs{
				input:    "GET_RELATIONSHIP Yodhan Maternal-Aunt",
				fMembers: familyMembers,
			},
			"Tritha",
		},
		{
			"get relationship paternal aunt",
			inputs{
				input:    "GET_RELATIONSHIP Vasa Paternal-Aunt",
				fMembers: familyMembers,
			},
			"Atya",
		},
		{
			"get relationship paternal uncle",
			inputs{
				input:    "GET_RELATIONSHIP Ahit Paternal-Uncle",
				fMembers: familyMembers,
			},
			"Chit Ish Vich",
		},
		{
			"get relationship siblings",
			inputs{
				input:    "GET_RELATIONSHIP Chit Siblings",
				fMembers: familyMembers,
			},
			"Ish Vich Aras Satya",
		},
		{
			"get relationship sister in law none",
			inputs{
				input:    "GET_RELATIONSHIP Shan Sister-In-Law",
				fMembers: familyMembers,
			},
			"NONE",
		},
		{
			"get relationship sister in law",
			inputs{
				input:    "GET_RELATIONSHIP Jaya Sister-In-Law",
				fMembers: familyMembers,
			},
			"Tritha",
		},
		{
			"get relationship brother in law none",
			inputs{
				input:    "GET_RELATIONSHIP Jnki Brother-In-Law",
				fMembers: familyMembers,
			},
			"NONE",
		},
		{
			"get relationship brother in law",
			inputs{
				input:    "GET_RELATIONSHIP Vyan Brother-In-Law",
				fMembers: familyMembers,
			},
			"Chit Ish Vich Aras",
		},
	}

	for _, td := range testData1 {
		t.Run(td.testName, func(t *testing.T) {
			result := familytree.ProcessInputs(td.inputData.input, td.inputData.fMembers)
			if result != td.result {
				t.Errorf("Expected output is %v for test case %v and got %v", td.result, td.testName, result)
			}
		})
	}
}
