package firestore

import pb "google.golang.org/genproto/googleapis/firestore/v1"

// GenerateNewDocumentSnapshot generates a new document from a protobuf Document
func (c *Client) GenerateNewDocumentSnapshot(proto *pb.Document) (*DocumentSnapshot, error) {
	docRef, err := pathToDoc(proto.Name, c)
	if err != nil {
		return nil, err
	}
	doc, err := newDocumentSnapshot(docRef, proto, c, proto.UpdateTime)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
