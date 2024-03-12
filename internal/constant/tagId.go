package constant

const HWPTAG_BEGIN = 0x10

const (
	DOCINFO_HWPTAG_DOCUMENT_PROPERTIES = HWPTAG_BEGIN
	DOCINFO_HWPTAG_ID_MAPPINGS         = HWPTAG_BEGIN + 1
	DOCINFO_HWPTAG_BIN_DATA            = HWPTAG_BEGIN + 2
	DOCINFO_HWPTAG_FACE_NAME           = HWPTAG_BEGIN + 3
	DOCINFO_HWPTAG_BORDER_FILL         = HWPTAG_BEGIN + 4
	DOCINFO_HWPTAG_CHAR_SHAPE          = HWPTAG_BEGIN + 5
	DOCINFO_HWPTAG_TAB_DEF             = HWPTAG_BEGIN + 6
	DOCINFO_HWPTAG_NUMBERING           = HWPTAG_BEGIN + 7
	DOCINFO_HWPTAG_BULLET              = HWPTAG_BEGIN + 8
	DOCINFO_HWPTAG_PARA_SHAPE          = HWPTAG_BEGIN + 9
	DOCINFO_HWPTAG_STYLE               = HWPTAG_BEGIN + 10
	DOCINFO_HWPTAG_DOC_DATA            = HWPTAG_BEGIN + 11
	DOCINFO_HWPTAG_DISTRIBUTE_DOC_DATA = HWPTAG_BEGIN + 12
	//RESERVED                            = HWPTAG_BEGIN + 13
	DOCINFO_HWPTAG_COMPATIBLE_DOCUMENT  = HWPTAG_BEGIN + 14
	DOCINFO_HWPTAG_LAYOUT_COMPATIBILITY = HWPTAG_BEGIN + 15
	DOCINFO_HWPTAG_TRACKCHANGE          = HWPTAG_BEGIN + 16
	DOCINFO_HWPTAG_MEMO_SHAPE           = HWPTAG_BEGIN + 76
	DOCINFO_HWPTAG_FORBIDDEN_CHAR       = HWPTAG_BEGIN + 78
	DOCINFO_HWPTAG_TRACK_CHANGE         = HWPTAG_BEGIN + 80
	DOCINFO_HWPTAG_TRACK_CHANGE_AUTHOR  = HWPTAG_BEGIN + 81
)

const (
	SECTION_HWPTAG_PARA_HEADER               = HWPTAG_BEGIN + 50
	SECTION_HWPTAG_PARA_TEXT                 = HWPTAG_BEGIN + 51
	SECTION_HWPTAG_PARA_CHAR_SHAPE           = HWPTAG_BEGIN + 52
	SECTION_HWPTAG_PARA_LINE_SEG             = HWPTAG_BEGIN + 53
	SECTION_HWPTAG_PARA_RANGE_TAG            = HWPTAG_BEGIN + 54
	SECTION_HWPTAG_CTRL_HEADER               = HWPTAG_BEGIN + 55
	SECTION_HWPTAG_LIST_HEADER               = HWPTAG_BEGIN + 56
	SECTION_HWPTAG_PAGE_DEF                  = HWPTAG_BEGIN + 57
	SECTION_HWPTAG_FOOTNOTE_SHAPE            = HWPTAG_BEGIN + 58
	SECTION_HWPTAG_PAGE_BORDER_FILL          = HWPTAG_BEGIN + 59
	SECTION_HWPTAG_SHAPE_COMPONENT           = HWPTAG_BEGIN + 60
	SECTION_HWPTAG_TABLE                     = HWPTAG_BEGIN + 61
	SECTION_HWPTAG_SHAPE_COMPONENT_LINE      = HWPTAG_BEGIN + 62
	SECTION_HWPTAG_SHAPE_COMPONENT_RECTANGLE = HWPTAG_BEGIN + 63
	SECTION_HWPTAG_SHAPE_COMPONENT_ELLIPSE   = HWPTAG_BEGIN + 64
	SECTION_HWPTAG_SHAPE_COMPONENT_ARC       = HWPTAG_BEGIN + 65
	SECTION_HWPTAG_SHAPE_COMPONENT_POLYGON   = HWPTAG_BEGIN + 66
	SECTION_HWPTAG_SHAPE_COMPONENT_CURVE     = HWPTAG_BEGIN + 67
	SECTION_HWPTAG_SHAPE_COMPONENT_OLE       = HWPTAG_BEGIN + 68
	SECTION_HWPTAG_SHAPE_COMPONENT_PICTURE   = HWPTAG_BEGIN + 69
	SECTION_HWPTAG_SHAPE_COMPONENT_CONTAINER = HWPTAG_BEGIN + 70
	SECTION_HWPTAG_CTRL_DATA                 = HWPTAG_BEGIN + 71
	SECTION_HWPTAG_EQEDIT                    = HWPTAG_BEGIN + 72
	//RESERVED                                 = HWPTAG_BEGIN + 73
	SECTION_HWPTAG_SHAPE_COMPONENT_TEXTART = HWPTAG_BEGIN + 74
	SECTION_HWPTAG_FORM_OBJECT             = HWPTAG_BEGIN + 75
	SECTION_HWPTAG_MEMO_SHAPE              = HWPTAG_BEGIN + 76
	SECTION_HWPTAG_MEMO_LIST               = HWPTAG_BEGIN + 77
	SECTION_HWPTAG_CHART_DATA              = HWPTAG_BEGIN + 79
	SECTION_HWPTAG_VIDEO_DATA              = HWPTAG_BEGIN + 82
	SECTION_HWPTAG_SHAPE_COMPONENT_UNKNOWN = HWPTAG_BEGIN + 99
)