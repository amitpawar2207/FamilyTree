# FamilyTree
Problem Statement -
Write code to model out the King Shan family tree so that: 
• Given a ‘name’ and a ‘relationship’, you should output the people corresponding to the relationship in the order in which they were added   to the family tree.
• You should be able to add a child to any family in the tree through the mother

Predefined Family Tree -
Shan(M) + Anga(F)
Children of Shan and Anga -
Chit(M), Ish(M), Vich(M), Aras(M), Satya(F)
Partner of Members -
Chit(M)-Amba(F), Vich(M)-Lika(F), Aras(M)-Chitra(F), Satya(F)-Vyan(M)
Children of Members -
Chit + Amba = Dritha(F), Tritha(F), Vritha(M)
Vich + Lika = Vila(F), Chika(F)
Aras + Chitra = Jnki(F), Ahit(M)
Satya + Vyan = Asva(M), Vyas(M), Atya(F)
Partners -
Dritha(F)-Jaya(M), Jnki(F)-Arit(M), Asva(M)-Satvy(F), Vyas(M)-Krpi(F)
Childrens - 
Dritha + Jaya = Yodhan(M)
Jnki + Arit = Laki(M), Lavnya(F)
Satvy + Asva = Vasa(M)
Vyas + Krpi = Kriya(M), Krithi(F) 

Note - Generate The Family tree described above before processing inputs from file.

Sample Inputs and outputs -
Input format to add a child: ADD_CHILD ”Mother’s-Name" "Child's-Name" "Gender"

Input format to find the people belonging to a relationship: GET_RELATIONSHIP ”Name” “Relationship”

Example Test File  - 
ADD_CHILD Chitra Aria Female
GET_RELATIONSHIP Lavnya Maternal-Aunt
GET_RELATIONSHIP Aria Siblings

Output on finding the relationship:
CHILD_ADDITION_SUCCEEDED
Aria
Jnki Ahit

Sample Inputs and Outputs - 
Sample -1 
Input - 
ADD_CHILD Pjali Srutak Male
GET_RELATIONSHIP Pjali Son

Output -1 
PERSON_NOT_FOUND
PERSON_NOT_FOUND

Sample -2 
ADD_CHILD Asva Vani Female
GET_RELATIONSHIP Vasa Siblings

Output -2
CHILD_ADDITION_FAILED
NONE

Sample - 3
GET_RELATIONSHIP Atya Sister-In-Law

Output -3
Satvy Krpi


Run this Code using - 
go run main.go <absolute path of input file>
example - go run main.go /home/user/Documents/Inputfiles/input1.txt