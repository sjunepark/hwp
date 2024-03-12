package constant

type TagID uint16

func (t TagID) String() string {
	switch t {
	case HWPTAG_DOCUMENT_PROPERTIES:
		return "HWPTAG_DOCUMENT_PROPERTIES"
	case HWPTAG_ID_MAPPINGS:
		return "HWPTAG_ID_MAPPINGS"
	case HWPTAG_BIN_DATA:
		return "HWPTAG_BIN_DATA"
	case HWPTAG_FACE_NAME:
		return "HWPTAG_FACE_NAME"
	case HWPTAG_BORDER_FILL:
		return "HWPTAG_BORDER_FILL"
	case HWPTAG_CHAR_SHAPE:
		return "HWPTAG_CHAR_SHAPE"
	case HWPTAG_TAB_DEF:
		return "HWPTAG_TAB_DEF"
	case HWPTAG_NUMBERING:
		return "HWPTAG_NUMBERING"
	case HWPTAG_BULLET:
		return "HWPTAG_BULLET"
	case HWPTAG_PARA_SHAPE:
		return "HWPTAG_PARA_SHAPE"
	case HWPTAG_STYLE:
		return "HWPTAG_STYLE"
	case HWPTAG_DOC_DATA:
		return "HWPTAG_DOC_DATA"
	case HWPTAG_DISTRIBUTE_DOC_DATA:
		return "HWPTAG_DISTRIBUTE_DOC_DATA"
	case HWPTAG_COMPATIBLE_DOCUMENT:
		return "HWPTAG_COMPATIBLE_DOCUMENT"
	case HWPTAG_LAYOUT_COMPATIBILITY:
		return "HWPTAG_LAYOUT_COMPATIBILITY"
	case HWPTAG_TRACKCHANGE:
		return "HWPTAG_TRACKCHANGE"
	case HWPTAG_FORBIDDEN_CHAR:
		return "HWPTAG_FORBIDDEN_CHAR"
	case HWPTAG_TRACK_CHANGE:
		return "HWPTAG_TRACK_CHANGE"
	case HWPTAG_TRACK_CHANGE_AUTHOR:
		return "HWPTAG_TRACK_CHANGE_AUTHOR"
	case HWPTAG_PARA_HEADER:
		return "HWPTAG_PARA_HEADER"
	case HWPTAG_PARA_TEXT:
		return "HWPTAG_PARA_TEXT"
	case HWPTAG_PARA_CHAR_SHAPE:
		return "HWPTAG_PARA_CHAR_SHAPE"
	case HWPTAG_PARA_LINE_SEG:
		return "HWPTAG_PARA_LINE_SEG"
	case HWPTAG_PARA_RANGE_TAG:
		return "HWPTAG_PARA_RANGE_TAG"
	case HWPTAG_CTRL_HEADER:
		return "HWPTAG_CTRL_HEADER"
	case HWPTAG_LIST_HEADER:
		return "HWPTAG_LIST_HEADER"
	case HWPTAG_PAGE_DEF:
		return "HWPTAG_PAGE_DEF"
	case HWPTAG_FOOTNOTE_SHAPE:
		return "HWPTAG_FOOTNOTE_SHAPE"
	case HWPTAG_PAGE_BORDER_FILL:
		return "HWPTAG_PAGE_BORDER_FILL"
	case HWPTAG_SHAPE_COMPONENT:
		return "HWPTAG_SHAPE_COMPONENT"
	case HWPTAG_TABLE:
		return "HWPTAG_TABLE"
	case HWPTAG_SHAPE_COMPONENT_LINE:
		return "HWPTAG_SHAPE_COMPONENT_LINE"
	case HWPTAG_SHAPE_COMPONENT_RECTANGLE:
		return "HWPTAG_SHAPE_COMPONENT_RECTANGLE"
	case HWPTAG_SHAPE_COMPONENT_ELLIPSE:
		return "HWPTAG_SHAPE_COMPONENT_ELLIPSE"
	case HWPTAG_SHAPE_COMPONENT_ARC:
		return "HWPTAG_SHAPE_COMPONENT_ARC"
	case HWPTAG_SHAPE_COMPONENT_POLYGON:
		return "HWPTAG_SHAPE_COMPONENT_POLYGON"
	case HWPTAG_SHAPE_COMPONENT_CURVE:
		return "HWPTAG_SHAPE_COMPONENT_CURVE"
	case HWPTAG_SHAPE_COMPONENT_OLE:
		return "HWPTAG_SHAPE_COMPONENT_OLE"
	case HWPTAG_SHAPE_COMPONENT_PICTURE:
		return "HWPTAG_SHAPE_COMPONENT_PICTURE"
	case HWPTAG_SHAPE_COMPONENT_CONTAINER:
		return "HWPTAG_SHAPE_COMPONENT_CONTAINER"
	case HWPTAG_CTRL_DATA:
		return "HWPTAG_CTRL_DATA"
	case HWPTAG_EQEDIT:
		return "HWPTAG_EQEDIT"
	case HWPTAG_SHAPE_COMPONENT_TEXTART:
		return "HWPTAG_SHAPE_COMPONENT_TEXTART"
	case HWPTAG_FORM_OBJECT:
		return "HWPTAG_FORM_OBJECT"
	case HWPTAG_MEMO_LIST:
		return "HWPTAG_MEMO_LIST"
	case HWPTAG_CHART_DATA:
		return "HWPTAG_CHART_DATA"
	case HWPTAG_VIDEO_DATA:
		return "HWPTAG_VIDEO_DATA"
	case HWPTAG_SHAPE_COMPONENT_UNKNOWN:
		return "HWPTAG_SHAPE_COMPONENT_UNKNOWN"
	case HWPTAG_MEMO_SHAPE:
		return "HWPTAG_MEMO_SHAPE"
	default:
		return "Unknown"
	}
}

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage,GoCommentStart
const HWPTAG_BEGIN TagID = 0x10

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage,GoCommentStart
const (
	// Tags for DocInfo
	HWPTAG_DOCUMENT_PROPERTIES  TagID = HWPTAG_BEGIN
	HWPTAG_ID_MAPPINGS          TagID = HWPTAG_BEGIN + 1
	HWPTAG_BIN_DATA             TagID = HWPTAG_BEGIN + 2
	HWPTAG_FACE_NAME            TagID = HWPTAG_BEGIN + 3
	HWPTAG_BORDER_FILL          TagID = HWPTAG_BEGIN + 4
	HWPTAG_CHAR_SHAPE           TagID = HWPTAG_BEGIN + 5
	HWPTAG_TAB_DEF              TagID = HWPTAG_BEGIN + 6
	HWPTAG_NUMBERING            TagID = HWPTAG_BEGIN + 7
	HWPTAG_BULLET               TagID = HWPTAG_BEGIN + 8
	HWPTAG_PARA_SHAPE           TagID = HWPTAG_BEGIN + 9
	HWPTAG_STYLE                TagID = HWPTAG_BEGIN + 10
	HWPTAG_DOC_DATA             TagID = HWPTAG_BEGIN + 11
	HWPTAG_DISTRIBUTE_DOC_DATA  TagID = HWPTAG_BEGIN + 12
	HWPTAG_COMPATIBLE_DOCUMENT  TagID = HWPTAG_BEGIN + 14
	HWPTAG_LAYOUT_COMPATIBILITY TagID = HWPTAG_BEGIN + 15
	HWPTAG_TRACKCHANGE          TagID = HWPTAG_BEGIN + 16
	HWPTAG_FORBIDDEN_CHAR       TagID = HWPTAG_BEGIN + 78
	HWPTAG_TRACK_CHANGE         TagID = HWPTAG_BEGIN + 80
	HWPTAG_TRACK_CHANGE_AUTHOR  TagID = HWPTAG_BEGIN + 81

	// Tags for Section
	HWPTAG_PARA_HEADER               TagID = HWPTAG_BEGIN + 50
	HWPTAG_PARA_TEXT                 TagID = HWPTAG_BEGIN + 51
	HWPTAG_PARA_CHAR_SHAPE           TagID = HWPTAG_BEGIN + 52
	HWPTAG_PARA_LINE_SEG             TagID = HWPTAG_BEGIN + 53
	HWPTAG_PARA_RANGE_TAG            TagID = HWPTAG_BEGIN + 54
	HWPTAG_CTRL_HEADER               TagID = HWPTAG_BEGIN + 55
	HWPTAG_LIST_HEADER               TagID = HWPTAG_BEGIN + 56
	HWPTAG_PAGE_DEF                  TagID = HWPTAG_BEGIN + 57
	HWPTAG_FOOTNOTE_SHAPE            TagID = HWPTAG_BEGIN + 58
	HWPTAG_PAGE_BORDER_FILL          TagID = HWPTAG_BEGIN + 59
	HWPTAG_SHAPE_COMPONENT           TagID = HWPTAG_BEGIN + 60
	HWPTAG_TABLE                     TagID = HWPTAG_BEGIN + 61
	HWPTAG_SHAPE_COMPONENT_LINE      TagID = HWPTAG_BEGIN + 62
	HWPTAG_SHAPE_COMPONENT_RECTANGLE TagID = HWPTAG_BEGIN + 63
	HWPTAG_SHAPE_COMPONENT_ELLIPSE   TagID = HWPTAG_BEGIN + 64
	HWPTAG_SHAPE_COMPONENT_ARC       TagID = HWPTAG_BEGIN + 65
	HWPTAG_SHAPE_COMPONENT_POLYGON   TagID = HWPTAG_BEGIN + 66
	HWPTAG_SHAPE_COMPONENT_CURVE     TagID = HWPTAG_BEGIN + 67
	HWPTAG_SHAPE_COMPONENT_OLE       TagID = HWPTAG_BEGIN + 68
	HWPTAG_SHAPE_COMPONENT_PICTURE   TagID = HWPTAG_BEGIN + 69
	HWPTAG_SHAPE_COMPONENT_CONTAINER TagID = HWPTAG_BEGIN + 70
	HWPTAG_CTRL_DATA                 TagID = HWPTAG_BEGIN + 71
	HWPTAG_EQEDIT                    TagID = HWPTAG_BEGIN + 72
	HWPTAG_SHAPE_COMPONENT_TEXTART   TagID = HWPTAG_BEGIN + 74
	HWPTAG_FORM_OBJECT               TagID = HWPTAG_BEGIN + 75
	HWPTAG_MEMO_LIST                 TagID = HWPTAG_BEGIN + 77
	HWPTAG_CHART_DATA                TagID = HWPTAG_BEGIN + 79
	HWPTAG_VIDEO_DATA                TagID = HWPTAG_BEGIN + 82
	HWPTAG_SHAPE_COMPONENT_UNKNOWN   TagID = HWPTAG_BEGIN + 99

	//	Common tags
	HWPTAG_MEMO_SHAPE TagID = HWPTAG_BEGIN + 76
)