// generated from ng_file_enum.ts.go
export enum GongdocCommandType {
	// insertion point	
	DIAGRAM_BASICFIELD_CREATE = "DIAGRAM_BASICFIELD_CREATE",
	DIAGRAM_BASICFIELD_DELETE = "DIAGRAM_BASICFIELD_DELETE",
	DIAGRAM_ELEMENT_CREATE = "DIAGRAM_ELEMENT_CREATE",
	DIAGRAM_ELEMENT_DELETE = "DIAGRAM_ELEMENT_DELETE",
	DIAGRAM_GONGSTRUCT_CREATE = "DIAGRAM_GONGSTRUCT_CREATE",
	DIAGRAM_GONGSTRUCT_DELETE = "DIAGRAM_GONGSTRUCT_DELETE",
	DIAGRAM_POINTER_TO_GONGSTRUCT_CREATE = "DIAGRAM_POINTER_TO_GONGSTRUCT_CREATE",
	DIAGRAM_POINTER_TO_GONGSTRUCT_DELETE = "DIAGRAM_POINTER_TO_GONGSTRUCT_DELETE",
	DIAGRAM_SLICE_OF_POINTERS_TO_GONGSTRUCT_CREATE = "DIAGRAM_SLICE_OF_POINTERS_TO_GONGSTRUCT_CREATE",
	DIAGRAM_SLICE_OF_POINTERS_TO_GONGSTRUCT_DELETE = "DIAGRAM_SLICE_OF_POINTERS_TO_GONGSTRUCT_DELETE",
	MARSHALL_DIAGRAM = "MARSHALL_ALL_DIAGRAMS",
	PRINT_ALL_DOCUMENTS = "PRINT_ALL_DOCUMENTS",
}

export interface GongdocCommandTypeSelect {
	value: string;
	viewValue: string;
}

export const GongdocCommandTypeList: GongdocCommandTypeSelect[] = [ // insertion point	
	{ value: GongdocCommandType.DIAGRAM_BASICFIELD_CREATE, viewValue: "DIAGRAM_BASICFIELD_CREATE" },
	{ value: GongdocCommandType.DIAGRAM_BASICFIELD_DELETE, viewValue: "DIAGRAM_BASICFIELD_DELETE" },
	{ value: GongdocCommandType.DIAGRAM_ELEMENT_CREATE, viewValue: "DIAGRAM_ELEMENT_CREATE" },
	{ value: GongdocCommandType.DIAGRAM_ELEMENT_DELETE, viewValue: "DIAGRAM_ELEMENT_DELETE" },
	{ value: GongdocCommandType.DIAGRAM_GONGSTRUCT_CREATE, viewValue: "DIAGRAM_GONGSTRUCT_CREATE" },
	{ value: GongdocCommandType.DIAGRAM_GONGSTRUCT_DELETE, viewValue: "DIAGRAM_GONGSTRUCT_DELETE" },
	{ value: GongdocCommandType.DIAGRAM_POINTER_TO_GONGSTRUCT_CREATE, viewValue: "DIAGRAM_POINTER_TO_GONGSTRUCT_CREATE" },
	{ value: GongdocCommandType.DIAGRAM_POINTER_TO_GONGSTRUCT_DELETE, viewValue: "DIAGRAM_POINTER_TO_GONGSTRUCT_DELETE" },
	{ value: GongdocCommandType.DIAGRAM_SLICE_OF_POINTERS_TO_GONGSTRUCT_CREATE, viewValue: "DIAGRAM_SLICE_OF_POINTERS_TO_GONGSTRUCT_CREATE" },
	{ value: GongdocCommandType.DIAGRAM_SLICE_OF_POINTERS_TO_GONGSTRUCT_DELETE, viewValue: "DIAGRAM_SLICE_OF_POINTERS_TO_GONGSTRUCT_DELETE" },
	{ value: GongdocCommandType.MARSHALL_DIAGRAM, viewValue: "MARSHALL_ALL_DIAGRAMS" },
	{ value: GongdocCommandType.PRINT_ALL_DOCUMENTS, viewValue: "PRINT_ALL_DOCUMENTS" },
];
