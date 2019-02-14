package Doco

func New() DocoSpecimen {
	//Add Single Page Tree Node
	//Add Pages
	docoMeta := DocoMeta{PdfVersion:1.7, Dimensions:DocoDimensions{Width:595, Height:842}, Font:DocoFont{}}
	docoPageTree := DocoPageTree{Children: &[]DocoPage{}}
	docoPage := DocoPage{PageNumber:1, Parent:&docoPageTree}
	*docoPageTree.Children = append(*docoPageTree.Children, docoPage)


	doco := Doco{Meta:docoMeta, Pages:[]DocoPageTree{docoPageTree}}
	return &doco
}

