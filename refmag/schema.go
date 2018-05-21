package refmag


/************************************************************************************************
* The idea is to represent a reference database as a matrix.
* A reference is represented as a row; the row index is the reference id.
* Each column is a key word; 1 represents existence of the key word in a refernce, 0 absence.
*************************************************************************************************/

// represent the database with matrix
type matrix struct {
	data []byte
	rows int
	cols int
}

// get the id (int representation) of a keyword
type keywords map[string]byte

// find the reference name from record id
type refs map[int]string